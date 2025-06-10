package services

import (
	"authentication-service/internal/feignclient"
	"authentication-service/models"
	"authentication-service/repositories"
	"authentication-service/utils"
)

type AuthService struct {
	EmployeeClient feignclient.IEmployeeClient
}

func NewAuthService(client feignclient.IEmployeeClient) *AuthService {
	return &AuthService{
		EmployeeClient: client,
	}
}

func (s *AuthService) Login(username, password string) utils.Response {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return utils.NewBadRequestResponse("Invalid credentials")
	}

	_, err = s.EmployeeClient.GetEmployeeByNIP(username)
	if err != nil {
		return utils.NewBadRequestResponse("Failed to get employee")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return utils.NewBadRequestResponse("Invalid credentials")
	}

	token, err := utils.GenerateJWT(user.Username, user.Roles)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to generate token")
	}

	return utils.NewSuccessResponse("Successfully logged in", token)
}

func Register(user *models.User) utils.Response {
	_, err := repositories.GetUserByEmail(user.Email)
	if err == nil {
		return utils.NewBadRequestResponse("Email already exists")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to encrypt password")
	}
	user.Password = hashedPassword

	if err := repositories.CreateUser(user); err != nil {
		return utils.NewInternalServerErrorResponse("Failed to create user")
	}
	return utils.NewSuccessResponse("User registered successfully", nil)
}
