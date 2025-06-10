package repositories

import (
	"authentication-service/database"
	"authentication-service/models"
)

func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Roles").Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return database.DB.Delete(&models.User{}, id).Error
}
