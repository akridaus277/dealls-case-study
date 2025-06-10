package seed

import (
	"authentication-service/database"
	"authentication-service/models"
	"authentication-service/utils"
	"strconv"

	"github.com/bxcodec/faker/v3"
)

func SeedUsers() {
	// Seed roles terlebih dahulu
	adminRole := models.Role{}
	employeeRole := models.Role{}
	database.DB.Debug()
	var countRoleAdmin int64
	database.DB.Model(&models.Role{}).
		Where("name = ?", "ADMIN_HR_PAYROLL").
		Count(&countRoleAdmin)
	if countRoleAdmin == 0 {
		database.DB.FirstOrCreate(&adminRole, models.Role{Name: "ADMIN_HR_PAYROLL"})
	}
	var countRoleEmployee int64
	database.DB.Model(&models.Role{}).
		Where("name = ?", "EMPLOYEE").
		Count(&countRoleEmployee)
	if countRoleEmployee == 0 {
		database.DB.FirstOrCreate(&employeeRole, models.Role{Name: "EMPLOYEE"})
	}
	hashedPassword, _ := utils.HashPassword("Dealls30!")

	// Cek apakah sudah ada admin dengan role ADMIN_HR_PAYROLL
	var countAdmin int64
	database.DB.Model(&models.User{}).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("roles.name = ?", "ADMIN_HR_PAYROLL").
		Count(&countAdmin)

	if countAdmin == 0 {
		adminUser := models.User{
			Email:    faker.LastName() + "@dealls.com",
			Username: "3030300",
			Password: hashedPassword,
			Roles:    []models.Role{employeeRole, adminRole}, // dua role
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		database.DB.Create(&adminUser)
	}

	// Cek jumlah user dengan role EMPLOYEE
	var countEmployees int64
	database.DB.Model(&models.User{}).
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("roles.name = ?", "EMPLOYEE").
		Count(&countEmployees)

	// Jika kurang dari 3 employee, buat tambahan sampai 3
	if countEmployees < 4 {
		for i := int(countEmployees); i < 4; i++ {
			employeeUser := models.User{
				Email:    faker.LastName() + "@dealls.com",
				Username: strconv.Itoa(3030301 + i),
				Password: hashedPassword,
				Roles:    []models.Role{employeeRole},
				Auditable: models.Auditable{
					CreatedBy: "SYSTEM",
					UpdatedBy: "SYSTEM",
				},
			}
			database.DB.Create(&employeeUser)
		}
	}
}
