package models

import (
	"api-payments/forms"
	"api-payments/services"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (model Transaction) Create(input forms.CreateTransaction) (primitive.ObjectID, error) {
	var amount float32
	for _, v := range input.Products {
		amount += v.Price * float32(v.Amount)
	}
	transaction := Transaction{
		Id:      primitive.NewObjectID(),
		Status:  TRANSACTION_STATUS_CREATED,
		OrderId: input.Id,
		UserId:  input.UserId,
		Amount:  amount,
	}

	collection := getCollection()
	result, err := collection.InsertOne(context.TODO(), transaction)

	go new(services.PaymentsService).ProcessPayment(result.InsertedID)

	if err != nil {
		log.Printf("Could not create Transaction: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), err
}

func (model Transaction) FindOneWithUserId(id primitive.ObjectID, userId string) (*Transaction, error) {
	var transaction Transaction

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}, bson.E{"userId", userId}})

	if result == nil {
		return nil, errors.New("Could not find a Transaction")
	}
	err := result.Decode(&transaction)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Transaction")
		}
		return nil, err
	}

	return &transaction, nil
}

func (model Transaction) Cancel(id primitive.ObjectID, userId string) (*Transaction, error) {
	var transaction Transaction

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}, bson.E{"userId", userId}})
	if result == nil {
		return nil, errors.New("Could not find an Transaction")
	}

	err := result.Decode(&transaction)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Transactions: %v", transaction)
	if transaction.Status != TRANSACTION_STATUS_CREATED {
		return nil, errors.New("Could not cancel Transaction if Transaction status is not CREATED")
	}
	var updatedDocument Transaction
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", TRANSACTION_STATUS_CANCELLED}}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	updateError := getCollection().FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedDocument)

	if updateError != nil {
		if updateError == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Transaction")
		}
		log.Fatal(err)
	}
	return &updatedDocument, nil
}

func getCollection() *mongo.Collection {
	return dbConnect.Db.Collection(TRANSACTION_COLLECTION_NAME)
}
