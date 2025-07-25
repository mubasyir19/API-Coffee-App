package seeders

import (
	"api-coffee-app/db"
	"api-coffee-app/models"
	"log"
)

func ProductSeed() {
	var coffee models.Category
	var nonCoffee models.Category
	var tea models.Category

	db.DB.First(&coffee, "name = ? ", "Coffee")
	db.DB.First(&nonCoffee, "name = ? ", "Non-Coffee")
	db.DB.First(&tea, "name = ? ", "Tea")

	products := []models.Product{
		{
			Name:        "Kopi Susu Gula Aren",
			CategoryID:  coffee.ID,
			Description: "Perpaduan Kopi, susu, dan manisnya gula aren",
			Price:       12000,
			Image:       "/images/coffee-milk.png",
		},
		{
			Name:        "Kopi Susu Gula Karamel",
			CategoryID:  coffee.ID,
			Description: "Perpaduan Kopi, susu, dan gurihnya karamel",
			Price:       12000,
			Image:       "/images/coffee-milk.png",
		},
		{
			Name:        "Kopi Susu Gula Hazelnut",
			CategoryID:  coffee.ID,
			Description: "Perpaduan Kopi, susu, dan gurihnya hazelnut",
			Price:       12000,
			Image:       "/images/coffee-milk.png",
		},
		{
			Name:        "Es Coklat Susu",
			CategoryID:  nonCoffee.ID,
			Description: "Perpaduan bubuk coklat asli dengan susu murni",
			Price:       12000,
			Image:       "/images/chocolate-milk.png",
		},
		{
			Name:        "Es Taro Susu",
			CategoryID:  nonCoffee.ID,
			Description: "Perpaduan bubuk taro dengan susu murni",
			Price:       12000,
			Image:       "/images/taro-milk2.png",
		},
		{
			Name:        "Es Teh Lemon",
			CategoryID:  nonCoffee.ID,
			Description: "Perpaduan teh dengan perasan lemon segar",
			Price:       12000,
			Image:       "/images/ice-lemon-tea.png",
		},
	}

	for _, product := range products {
		result := db.DB.Create(&product)
		if result.Error != nil {
			log.Println("Failed insert product", result.Error)
		}
	}
	log.Println("Seeder Product has been successfully running")
}
