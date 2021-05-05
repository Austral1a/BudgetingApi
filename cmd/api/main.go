package main

import (
	"log"

	"github.com/Austal1a/BudgetingApi/pkg/api/controllers"
	"github.com/Austal1a/BudgetingApi/pkg/api/repository"
	"github.com/Austal1a/BudgetingApi/pkg/api/routes"
	"github.com/Austal1a/BudgetingApi/pkg/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

const DBName = "mydb"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	dbInstance := db.DB{}

	dbInstance.Connect()
	dbInstance.Ping()
	defer dbInstance.Disconnect()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	categoriesRepo := repository.NewCategoriesRepository(dbInstance.Client.Database(DBName))
	categoriesController := controllers.NewCategoriesController(categoriesRepo)
	categoriesRoutes := routes.NewCategoriesRoutes(categoriesController)

	categoriesRoutes.Install(app)

	log.Fatal(app.Listen(":8080"))
}
