package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dimapog/jwt-microservice/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(req *LoginRequest) (*LoginResponse, error)
}

type service struct {
	userService user.Service
}

func NewService(userService user.Service) Service {
	return &service{userService: userService}
}

func (s *service) Login(req *LoginRequest) (*LoginResponse, error) {
	userModel, err := s.userService.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userModel.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &LoginResponse{Token: tokenString}, nil
}
