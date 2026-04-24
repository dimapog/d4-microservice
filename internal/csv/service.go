package csv

import (
	"encoding/csv"
	"fmt"
	"mime/multipart"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Service interface {
	UploadCSV(fileHeader *multipart.FileHeader) (*CSVUploadResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) UploadCSV(fileHeader *multipart.FileHeader) (*CSVUploadResponse, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}

	jobID := fmt.Sprintf("csv-job-%d", time.Now().UnixNano())

	go s.processCSV(file, jobID)

	return &CSVUploadResponse{
		JobID:  jobID,
		Status: "processing",
	}, nil
}

func (s *service) processCSV(file multipart.File, jobID string) {
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	header, err := reader.Read()
	if err != nil {
		fmt.Printf("[csv:%s] failed to read header: %v\n", jobID, err)
		return
	}

	columnIndex := mapHeaderIndexes(header)
	rows := make(chan *Client, 2000)
	wg := sync.WaitGroup{}
	workerCount := runtime.NumCPU()

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.consumeBatch(rows, jobID)
		}()
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Printf("[csv:%s] read error: %v\n", jobID, err)
			break
		}

		client := mapRecordToClient(record, columnIndex)
		if client == nil {
			continue
		}

		rows <- client
	}

	close(rows)
	wg.Wait()
	fmt.Printf("[csv:%s] finished processing\n", jobID)
}

func (s *service) consumeBatch(rows chan *Client, jobID string) {
	batchSize := 500
	clients := make([]*Client, 0, batchSize)

	for row := range rows {
		clients = append(clients, row)
		if len(clients) >= batchSize {
			if err := s.repo.CreateClientsBatch(clients); err != nil {
				fmt.Printf("[csv:%s] insert batch error: %v\n", jobID, err)
			}
			clients = make([]*Client, 0, batchSize)
		}
	}

	if len(clients) > 0 {
		if err := s.repo.CreateClientsBatch(clients); err != nil {
			fmt.Printf("[csv:%s] insert batch error: %v\n", jobID, err)
		}
	}
}

func mapHeaderIndexes(header []string) map[string]int {
	indexes := make(map[string]int)
	for i, value := range header {
		indexes[strings.TrimSpace(strings.ToLower(value))] = i
	}
	return indexes
}

func mapRecordToClient(record []string, header map[string]int) *Client {
	if len(record) == 0 {
		return nil
	}

	client := &Client{
		Name:        getField(record, header, "name"),
		LastName:    getField(record, header, "last_name"),
		Email:       getField(record, header, "email"),
		Phone:       getField(record, header, "phone"),
		City:        getField(record, header, "city"),
		Street:      getField(record, header, "street"),
		HouseNumber: getField(record, header, "house_number"),
		State:       getField(record, header, "state"),
	}

	if idStr := getField(record, header, "id"); idStr != "" {
		if idVal, err := strconv.ParseUint(idStr, 10, 64); err == nil && idVal > 0 {
			client.ID = uint(idVal)
		}
	}

	return client
}

func getField(record []string, header map[string]int, key string) string {
	idx, ok := header[key]
	if !ok || idx < 0 || idx >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[idx])
}
