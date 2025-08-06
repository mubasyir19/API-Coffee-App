package main

import (
	"api-coffee-app/db"
	"api-coffee-app/handlers"
	"api-coffee-app/middleware"
	"api-coffee-app/repositories"
	"api-coffee-app/services"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Izinkan frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "welcome",
		})
	})

	db.ConnectDB()
	database := db.DB

	// seeders.CategorySeeds()
	// seeders.ProductSeed()

	customerRepository := repositories.NewCustomerRepository(database)
	customerService := services.NewCustomerService(customerRepository)
	customerHandler := handlers.NewCustomerHandler(&customerService)

	categoryRepository := repositories.NewCategoryRepository(database)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(&categoryService)

	productRepository := repositories.NewProductRepository(database)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(&productService)

	cartRepository := repositories.NewCartRepository(database)
	cartService := services.NewCartService(cartRepository, productRepository, customerRepository)
	cartHandler := handlers.NewCartHandler(&cartService)

	api := app.Group("/apiv1")

	api.GET("/verify", middleware.VerifyAuth, customerHandler.VerifyAuth)

	api.POST("/customer/register", customerHandler.Register)
	api.POST("/customer/login", customerHandler.Login)

	api.GET("/categories", categoryHandler.GetAllCategories)
	api.GET("/category/:id", categoryHandler.GetCategoryByID)
	api.POST("/category/add", categoryHandler.AddCategory)

	api.GET("/products", productHandler.GetAllProducts)
	api.GET("/product/id/:id", productHandler.GetProdutcByID)
	api.GET("/product/slug/:slug", productHandler.GetProductBySlug)
	api.GET("/product/name/:name", productHandler.GetProductByName)
	api.GET("/product/category/:category_id", productHandler.GetProductsByCategory)
	api.POST("/product/add", productHandler.AddProduct)

	api.GET("/cart", cartHandler.GetCartItems)
	api.POST("/cart/add", cartHandler.AddItemToCart)
	api.PUT("/cart/update", cartHandler.UpdateItemCart)
	api.DELETE("/cart/delete/:id", cartHandler.RemoveItemCart)

	app.Run(":8080")
}
