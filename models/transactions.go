package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TRANSACTION_STATUS_CREATED = "CREATED"
const TRANSACTION_STATUS_CONFIRMED = "CONFIRMED"
const TRANSACTION_STATUS_CANCELLED = "CANCELLED"
const TRANSACTION_STATUS_DELIVERED = "DELIVERED"

const TRANSACTION_COLLECTION_NAME = "transactions"

type Transaction struct {
	Id      primitive.ObjectID `json:"id"`
	Status  string             `json:"status"`
	OrderId string             `json:"orderId"`
	UserId  string             `json:"userId"`
	Amount  float32            `json:"amount"`
}
