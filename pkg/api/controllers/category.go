package controllers

import (
	"net/http"

	"github.com/Austal1a/BudgetingApi/pkg/api/errors"
	"github.com/Austal1a/BudgetingApi/pkg/api/models"
	"github.com/Austal1a/BudgetingApi/pkg/api/repository"
	"github.com/Austal1a/BudgetingApi/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoriesController interface {
	Create(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
}

type categoriesController struct {
	categoriesRepo repository.CategoriesRepository
}

func NewCategoriesController(categoriesRepo repository.CategoriesRepository) CategoriesController {
	return &categoriesController{categoriesRepo}
}

func (c *categoriesController) Create(ctx *fiber.Ctx) error {
	var category models.Category

	err := ctx.BodyParser(&category)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(utils.JsonError(errors.CategoryBadPayloadCode))
	}

	isCategoryExists, err := c.categoriesRepo.IsCategoryExists(category.Name)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(utils.JsonError(errors.CategoryNotFoundCode))
	}

	if isCategoryExists {
		return ctx.Status(http.StatusBadRequest).JSON(utils.JsonError(errors.CategoryIsExistsCode))
	}

	category.Id = primitive.NewObjectID()

	err = c.categoriesRepo.Create(&category)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.JsonError(errors.CategoryWasNotCreated))
	}

	ctx.Status(http.StatusCreated)

	return nil
}

func (c *categoriesController) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.categoriesRepo.GetAll()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.JsonError(errors.CategoriesCannotBeRetrieved))
	}

	return ctx.Status(http.StatusOK).JSON(&categories)
}
