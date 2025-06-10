package controllers

import (
	requestDto "authentication-service/dto/request"
	"authentication-service/models"
	"authentication-service/services"
	"authentication-service/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
}

// Constructor
func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request"))
		return
	}

	utils.SendResponse(c, services.Register(&user))
}

func (ac *AuthController) Login(c *gin.Context) {
	var data requestDto.LoginRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		utils.SendResponse(c, utils.NewBadRequestResponse("Invalid request body"))
		return
	}

	utils.SendResponse(c, ac.AuthService.Login(data.Username, data.Password))
}
