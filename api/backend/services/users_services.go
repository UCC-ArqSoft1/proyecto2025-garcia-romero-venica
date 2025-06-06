package services

import (
	"Backend/dao"
	"Backend/utils"
	"fmt"
)

type UsersClient interface {
	GetUserByUsername(username string) (dao.User, error)
}

type UsersService struct {
	usersClient UsersClient
}

func NewUsersService(usersClient UsersClient) *UsersService {
	return &UsersService{
		usersClient: usersClient,
	}
}

func (s *UsersService) Login(username string, password string) (int, string, error) {
	userDAO, err := s.usersClient.GetUserByUsername(username)
	if err != nil {
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}

	if utils.HashSHA256(password) != userDAO.PasswordHash {
		return 0, "", fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWT(userDAO.ID)
	if err != nil {
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}

	return userDAO.ID, token, nil
}
