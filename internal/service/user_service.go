package service

import (
	"errors"
	"user_crud/internal/model"
	"user_crud/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(req model.User) error {
	if req.Role != "admin" && req.Role != "owner" && req.Role != "customer" {
		return errors.New("role tidak valid, gunakan: admin, owner, atau customer")
	}
	return s.Repo.Create(&req)
}