package repository

import (
	"context"
	"time"

	"github.com/Austal1a/BudgetingApi/pkg/api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const CategoriesCollectionName = "categories"

const cancelQueryTimeoutSc = 4 * time.Second

type CategoriesRepository interface {
	Create(category *models.Category) error
	GetAll() (categories []*models.Category, err error)
	GetById(categoryId primitive.ObjectID) (category *models.Category, err error)
	IsCategoryExists(categoryId primitive.ObjectID) (isExists bool, err error)
	AddTransaction(transaction *models.Transaction) error
}

type categoriesRepository struct {
	c *mongo.Collection
}

func NewCategoriesRepository(db *mongo.Database) CategoriesRepository {
	return &categoriesRepository{db.Collection(CategoriesCollectionName)}
}

func (r *categoriesRepository) AddTransaction(transaction *models.Transaction) error {
	ctx, cancel := context.WithTimeout(context.Background(), cancelQueryTimeoutSc)
	defer cancel()

	addTransactionRes := r.c.FindOneAndUpdate(ctx, bson.D{{"_id", &transaction.CategoryId}}, bson.M{
		"$push": bson.M{
			"transactions": &transaction,
		},
	})

	if addTransactionRes.Err() != nil {
		return addTransactionRes.Err()
	}

	return nil
}

func (r *categoriesRepository) GetById(categoryId primitive.ObjectID) (category *models.Category, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelQueryTimeoutSc)
	defer cancel()

	categoryFindRes := r.c.FindOne(ctx, bson.M{"_id": categoryId})
	if categoryFindRes.Err() != nil {
		return nil, err
	}

	err = categoryFindRes.Decode(&category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoriesRepository) IsCategoryExists(categoryId primitive.ObjectID) (isExists bool, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelQueryTimeoutSc)
	defer cancel()

	categoryDocsAmount, err := r.c.CountDocuments(ctx, bson.M{"_id": categoryId})
	if err != nil {
		return false, nil
	}

	if categoryDocsAmount >= 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func (r *categoriesRepository) Create(category *models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), cancelQueryTimeoutSc)
	defer cancel()

	_, err := r.c.InsertOne(ctx, &category)

	return err
}

func (r *categoriesRepository) GetAll() (categories []*models.Category, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelQueryTimeoutSc)
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
