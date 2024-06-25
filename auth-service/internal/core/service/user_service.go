package service

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/ports"
	"auth-service/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *domain.User) error
	Authenticate(username, password string) (*domain.User, error)
	GenerateToken(username string) (string, error)
}

type userService struct {
	userRepo   ports.UserRepository
	jwtManager *jwt.JWTManager
}

func NewUserService(repo ports.UserRepository, jwtManager *jwt.JWTManager) UserService {
	return &userService{userRepo: repo, jwtManager: jwtManager}
}

func (s *userService) RegisterUser(user *domain.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *userService) Authenticate(username, password string) (*domain.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GenerateToken(username string) (string, error) {
	return s.jwtManager.GenerateToken(username)
}
