package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	Id      primitive.ObjectID `json:"id"`
	Status  string             `json:"status"`
	OrderId string             `json:"orderId"`
	UserId  string             `json:"userId"`
	Amount  float32            `json:"amount"`
}
