package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Msg       string             `json:"msg" bson:"msg"`
	IsIncome  bool               `json:"isIncome" bson:"isIncome"`
	IsOutcome bool               `json:"isOutcome" bson:"isOutcome"`
}
