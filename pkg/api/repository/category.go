package repository

import (
	"context"
	"time"

	"github.com/Austal1a/BudgetingApi/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CategoriesCollectionName = "categories"

type CategoriesRepository interface {
	Create(category *models.Category) error
	GetAll() (categories []*models.Category, err error)
	IsCategoryExists(categoryName string) (isExists bool, err error)
}

type categoriesRepository struct {
	c *mongo.Collection
}

func NewCategoriesRepository(db *mongo.Database) CategoriesRepository {
	return &categoriesRepository{db.Collection(CategoriesCollectionName)}
}

func (r *categoriesRepository) IsCategoryExists(categoryName string) (isExists bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categoryFindRes := r.c.FindOne(ctx, bson.M{"name": categoryName})
	if categoryFindRes.Err() != nil {
		return false, err
	}

	return true, nil
}

func (r *categoriesRepository) Create(category *models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.c.InsertOne(ctx, &category)

	return err
}

func (r *categoriesRepository) GetAll() (categories []*models.Category, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	categoriesCursor, err := r.c.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if categoriesCursor.Err() != nil {
		return nil, categoriesCursor.Err()
	}

	err = categoriesCursor.All(ctx, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
