package user

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(req *CreateUserRequest) (*UserResponse, error)
	GetUserByID(id string) (*UserResponse, error)
	DeleteUser(id string) error
	GetUserByEmail(email string) (*User, error)
	UpdateUser(id uint, req *UpdateUserRequest) (*UserResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req *CreateUserRequest) (*UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

func (s *service) GetUserByID(id string) (*UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

func (s *service) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *service) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *service) UpdateUser(id uint, req *UpdateUserRequest) (*UserResponse, error) {
	// Convert uint to string for repository
	idStr := fmt.Sprintf("%d", id)

	user, err := s.repo.GetUserByID(idStr)
	if err != nil {
		return nil, err
	}

	// Update fields if provided
	if req.Age != nil {
		user.Age = sql.NullInt64{Int64: int64(*req.Age), Valid: true}
	}
	if req.Gender != nil {
		user.Gender = sql.NullString{String: *req.Gender, Valid: true}
	}
	if req.Weight != nil {
		user.Weight = sql.NullFloat64{Float64: *req.Weight, Valid: true}
	}
	if req.Height != nil {
		user.Height = sql.NullFloat64{Float64: *req.Height, Valid: true}
	}
	if req.RestingHeartRate != nil {
		user.RestingHeartRate = sql.NullInt64{Int64: int64(*req.RestingHeartRate), Valid: true}
	}
	if req.Units != nil {
		user.Units = sql.NullString{String: *req.Units, Valid: true}
	}

	if err := s.repo.UpdateUser(idStr, user); err != nil {
		return nil, err
	}

	return s.userToResponse(user), nil
}

func (s *service) userToResponse(user *User) *UserResponse {
	resp := &UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	if user.Age.Valid {
		age := int(user.Age.Int64)
		resp.Age = &age
	}
	if user.Gender.Valid {
		resp.Gender = &user.Gender.String
	}
	if user.Weight.Valid {
		resp.Weight = &user.Weight.Float64
	}
	if user.Height.Valid {
		resp.Height = &user.Height.Float64
	}
	if user.RestingHeartRate.Valid {
		rhr := int(user.RestingHeartRate.Int64)
		resp.RestingHeartRate = &rhr
	}
	if user.Units.Valid {
		resp.Units = &user.Units.String
	}

	return resp
}
