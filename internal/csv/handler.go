package csv

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	csvRoutes := router.Group("/csv")
	{
		csvRoutes.POST("/upload", h.UploadCSV)
	}
}

// UploadCSV godoc
// @Summary Upload CSV for async import
// @Description Upload a CSV file containing client rows for async batch import into the database
// @Tags csv
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file"
// @Success 202 {object} csv.CSVUploadResponse
// @Failure 400 {object} csv.ErrorResponse
// @Failure 500 {object} csv.ErrorResponse
// @Router /csv/upload [post]
func (h *Handler) UploadCSV(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file upload failed"})
		return
	}

	resp, err := h.service.UploadCSV(fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, resp)
}
