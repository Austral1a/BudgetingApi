package routes

import (
	"github.com/Austal1a/BudgetingApi/pkg/api/controllers"
	"github.com/gofiber/fiber/v2"
)

type categoriesRoutes struct {
	categoriesController controllers.CategoriesController
}

func NewCategoriesRoutes(categoriesController controllers.CategoriesController) Routes {
	return &categoriesRoutes{categoriesController}
}

func (r *categoriesRoutes) Install(app *fiber.App) {
	app.Post("/categories/create", r.categoriesController.Create)
	app.Get("/categories/getAll", r.categoriesController.GetAll)
}
