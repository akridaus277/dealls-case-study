package seed

import (
	"employee-service/config"
	"employee-service/models"
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v3"
)

func SeedEmployees() {

	targetNIPs := []string{"3030300", "3030302", "3030303", "3030304"}

	// Ambil semua NIP yang sudah ada di tabel employees dari target
	var existingNIPs []string
	config.DB.Model(&models.Employee{}).
		Where("nip IN ?", targetNIPs).
		Pluck("nip", &existingNIPs)

	// Buat map untuk lookup cepat
	existingMap := map[string]bool{}
	for _, nip := range existingNIPs {
		existingMap[nip] = true
	}

	// Loop target NIPs, hanya masukkan jika belum ada di DB
	for _, nip := range targetNIPs {
		if !existingMap[nip] {
			employee := models.Employee{
				Name:   faker.Name(),
				Nip:    nip,
				Salary: rand.Intn(5000000) + 3000000,
				Auditable: models.Auditable{
					CreatedBy: "SYSTEM",
					UpdatedBy: "SYSTEM",
				},
			}
			config.DB.Create(&employee)
			fmt.Printf("Seeded employee with NIP: %s\n", nip)
		} else {
			fmt.Printf("Skipped existing NIP: %s\n", nip)
		}
	}

	// Buat Employee dari user EMPLOYEE jika belum ada (gunakan Username sebagai NIP)
	for _, user := range existingNIPs {
		if !existingMap[user] {
			employee := models.Employee{
				Name:   faker.Name(),
				Nip:    user, // Username dipakai sebagai NIP
				Salary: rand.Intn(5000000) + 3000000,
				Auditable: models.Auditable{
					CreatedBy: "SYSTEM",
					UpdatedBy: "SYSTEM",
				},
			}
			config.DB.Create(&employee)
		}
	}

	var count int64
	config.DB.Model(&models.Employee{}).Count(&count)

	// Hanya insert jika kurang dari 100
	if count >= 100 {
		return
	}

	remaining := 100 - int(count)
	var employees []models.Employee

	for i := 0; i < remaining; i++ {
		employee := models.Employee{
			Name:   faker.Name(),
			Nip:    faker.CCNumber(),
			Salary: rand.Intn(5000000) + 3000000,
			Auditable: models.Auditable{
				CreatedBy: "SYSTEM",
				UpdatedBy: "SYSTEM",
			},
		}
		employees = append(employees, employee)
	}

	config.DB.Create(&employees)
}
