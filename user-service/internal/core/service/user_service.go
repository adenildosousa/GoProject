package service

import (
	"user-service/internal/core/domain"
	"user-service/internal/ports"
	"user-service/pkg/jwt"
)

type UserService interface {
	ValidateToken(token string) (*domain.User, error)
}

type userService struct {
	userRepo   ports.UserRepository
	jwtManager *jwt.JWTManager
}

func NewUserService(repo ports.UserRepository, jwtManager *jwt.JWTManager) UserService {
	return &userService{userRepo: repo, jwtManager: jwtManager}
}

func (s *userService) ValidateToken(token string) (*domain.User, error) {
	username, err := s.jwtManager.ValidateToken(token)
	if err != nil {
		return nil, err
	}

	return s.userRepo.GetUserByUsername(username)
}
