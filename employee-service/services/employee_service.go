package services

import (
	"context"
	"employee-service/config"
	responseDto "employee-service/dto/response"
	"employee-service/models"
	"employee-service/repositories"
	"employee-service/utils"
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func RegisterEmployee(emp models.Employee) utils.Response {

	if err := repositories.CreateEmployee(&emp); err != nil {
		return utils.NewInternalServerErrorResponse("Failed to create employee")
	}
	return utils.NewSuccessResponse("Employee registered successfully", nil)
}

func ListEmployees() utils.Response {
	cacheKey := "employees:list"
	var summaries []responseDto.EmployeeSummary
	val, err := config.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &summaries)
		return utils.NewSuccessResponse("Employee list retrieved successfully", summaries)
	}

	emps, err := repositories.GetAllEmployees()
	if err != nil {
		return utils.NewInternalServerErrorResponse("Failed to get employee list")
	}
	// Mapping ke summary
	for _, emp := range emps {
		summaries = append(summaries, responseDto.EmployeeSummary{
			Nip:    emp.Nip,
			Name:   emp.Name,
			Salary: emp.Salary,
		})
	}

	cacheData, _ := json.Marshal(summaries)
	config.RedisClient.Set(context.Background(), cacheKey, cacheData, 5*time.Minute)
	return utils.NewSuccessResponse("Employee list retrieved successfully", summaries)
}

func GetEmployeeByID(id int) utils.Response {
	cacheKey := fmt.Sprintf("employees:%d", id)
	val, err := config.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var emp models.Employee
		json.Unmarshal([]byte(val), &emp)
		return utils.NewSuccessResponse("Employee retrieved successfully", emp)
	}

	emp, err := repositories.GetEmployeeByID(id)
	if err != nil {
		return utils.NewNotFoundResponse("Employee not found")
	}

	cacheData, _ := json.Marshal(emp)
	config.RedisClient.Set(context.Background(), cacheKey, cacheData, 5*time.Minute)
	return utils.NewSuccessResponse("Employee retrieved successfully", emp)
}

func UpdateEmployee(id int, data map[string]interface{}) utils.Response {

	if err := repositories.UpdateEmployee(id, data); err != nil {
		return utils.NewInternalServerErrorResponse("Failed to update employee")
	}
	config.RedisClient.Del(context.Background(), fmt.Sprintf("employees:%d", id))
	return utils.NewSuccessResponse("Employee updated successfully", nil)
}

func DeleteEmployee(id int) utils.Response {
	if err := repositories.DeleteEmployee(id); err != nil {
		return utils.NewInternalServerErrorResponse("Failed to delete employee")
	}
	config.RedisClient.Del(context.Background(), fmt.Sprintf("employees:%d", id))
	return utils.NewSuccessResponse("Employee deleted successfully", nil)
}

func GetEmployeeByNIP(nip string) utils.Response {

	cacheKey := "employees:nip:" + nip
	val, err := config.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var emp models.Employee
		if err := json.Unmarshal([]byte(val), &emp); err == nil {
			return utils.NewSuccessResponse("Employee retrieved successfully", emp)
		}
	}

	emp, err := repositories.GetEmployeeByNip(nip)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewInternalServerErrorResponse("Employee not found")
		}
		return utils.NewInternalServerErrorResponse("Failed to get employee")
	}

	cacheData, _ := json.Marshal(emp)
	config.RedisClient.Set(context.Background(), cacheKey, cacheData, 5*time.Minute)

	return utils.NewSuccessResponse("Employee retrieved successfully", emp)
}
