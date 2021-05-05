package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id         primitive.ObjectID `json:"id" bson:"_id"`
	CategoryId primitive.ObjectID `json:"categoryId" bson:"category_id"`
	Msg        string             `json:"msg" bson:"msg"`
	IsIncome   bool               `json:"isIncome" bson:"isIncome"`
	IsOutcome  bool               `json:"isOutcome" bson:"isOutcome"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
}

type Category struct {
	Id           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Transactions []*Transaction     `json:"transactions" bson:"transactions"`
}
