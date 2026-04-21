package calculator

import (
	"fmt"
	"math"

	"github.com/dimapog/jwt-microservice/internal/user"
)

type Service interface {
	CalculateBMIByUserID(userID uint) (*BMIResponse, error)
	CalculateHeartRateZonesByUserID(userID uint) (*HRZResponse, error)
}

type service struct {
	userService user.Service
}

func NewService(userService user.Service) Service {
	return &service{userService: userService}
}

func (s *service) CalculateBMIByUserID(userID uint) (*BMIResponse, error) {
	// Fetch user by ID
	userIDStr := fmt.Sprintf("%d", userID)
	userResp, err := s.userService.GetUserByID(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if weight and height are provided
	if userResp.Weight == nil || userResp.Height == nil {
		return nil, fmt.Errorf("weight and height not set in user profile")
	}

	// Convert height from cm to meters
	heightMeters := *userResp.Height / 100

	// Calculate BMI
	bmi := *userResp.Weight / (heightMeters * heightMeters)
	bmi = math.Round(bmi*100) / 100

	category := s.getBMICategory(bmi)

	return &BMIResponse{
		Weight:   *userResp.Weight,
		Height:   *userResp.Height,
		BMI:      bmi,
		Category: category,
	}, nil
}

func (s *service) CalculateHeartRateZonesByUserID(userID uint) (*HRZResponse, error) {
	// Fetch user by ID
	userIDStr := fmt.Sprintf("%d", userID)
	userResp, err := s.userService.GetUserByID(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}

	// Check if age is provided
	if userResp.Age == nil {
		return nil, fmt.Errorf("age not set in user profile")
	}

	age := *userResp.Age
	maxHR := 220 - age

	zones := []Zone{
		{
			Name: "Z1 (Recovery)",
			Min:  int(0.5 * float64(maxHR)),
			Max:  int(0.6 * float64(maxHR)),
		},
		{
			Name: "Z2 (Fat burn)",
			Min:  int(0.6 * float64(maxHR)),
			Max:  int(0.7 * float64(maxHR)),
		},
		{
			Name: "Z3 (Endurance)",
			Min:  int(0.7 * float64(maxHR)),
			Max:  int(0.8 * float64(maxHR)),
		},
		{
			Name: "Z4 (Hard)",
			Min:  int(0.8 * float64(maxHR)),
			Max:  int(0.9 * float64(maxHR)),
		},
		{
			Name: "Z5 (Max)",
			Min:  int(0.9 * float64(maxHR)),
			Max:  maxHR,
		},
	}

	return &HRZResponse{
		Age:   age,
		MaxHR: maxHR,
		Zones: zones,
	}, nil
}

func (s *service) getBMICategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Underweight"
	case bmi < 25:
		return "Normal"
	case bmi < 30:
		return "Overweight"
	default:
		return "Obese"
	}
}
