package seeders

import (
	"api-coffee-app/db"
	"api-coffee-app/models"
	"log"
)

func CategorySeeds() {
	categories := []models.Category{
		{Name: "Coffee"},
		{Name: "Non-Coffee"},
		{Name: "Tea"},
	}

	for _, category := range categories {
		result := db.DB.Create(&category)
		if result.Error != nil {
			log.Println("Failed insert category", result.Error)
		}
	}
	log.Println("Seeder Category has been successfully running")
}
