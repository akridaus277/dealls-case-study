package repositories

import (
	"employee-service/config"
	"employee-service/models"
)

func CreateEmployee(e *models.Employee) error {
	return config.DB.Create(e).Error
}

func GetAllEmployees() ([]models.Employee, error) {
	var emps []models.Employee
	err := config.DB.Find(&emps).Error
	return emps, err
}

func GetEmployeeByID(id int) (models.Employee, error) {
	var emp models.Employee
	err := config.DB.First(&emp, id).Error
	return emp, err
}

func GetEmployeeByNip(nip string) (models.Employee, error) {
	var emp models.Employee
	err := config.DB.Where("nip = ?", nip).First(&emp).Error
	return emp, err
}

func UpdateEmployee(id int, data map[string]interface{}) error {
	return config.DB.Model(&models.Employee{}).Where("id = ?", id).Updates(data).Error
}

func DeleteEmployee(id int) error {
	return config.DB.Delete(&models.Employee{}, id).Error
}
