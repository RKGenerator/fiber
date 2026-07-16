package services

import (
	"test-fiber/auth"
	basefunc "test-fiber/base_func"
	"test-fiber/dto"
	"test-fiber/repositories"

	"github.com/gofiber/fiber/v3"
)

type UserService struct {
	userRepository repositories.UserInterface
}

func NewUserService(userRep repositories.UserInterface) *UserService {
	return &UserService{
		userRepository: userRep,
	}
}

func (s *UserService) UserLogin(req dto.AuthRequest) (*dto.AuthResponse, error) {

	user, err := s.userRepository.GetUserByEmail(req.Email)
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	passOk := basefunc.CheckPassword(req.Password, user.Password)
	if !passOk {
		return nil, fiber.ErrNotFound
	}

	token, err := auth.GenerateToken(*user)
	if err != nil {
		return nil, fiber.ErrInternalServerError
	}
	authResponse := dto.AuthResponse{
		AccessToken: token,
	}

	return &authResponse, nil

}
