package services

import (
	"api-payments/forms"
	"api-payments/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TRANSACTION_COLLECTION_NAME = "transactions"

type TransactionService struct {
}

func (model TransactionService) Create(input forms.CreateTransaction) (primitive.ObjectID, error) {
	var amount float32
	for _, v := range input.Products {
		amount += v.Price * float32(v.Amount)
	}
	transaction := models.Transaction{
		Id:      primitive.NewObjectID(),
		Status:  models.TRANSACTION_STATUS_CREATED,
		OrderId: input.Id,
		UserId:  input.UserId,
		Amount:  amount,
	}

	collection := getCollection()
	result, err := collection.InsertOne(context.TODO(), transaction)

	go model.ProcessPayment(result.InsertedID.(primitive.ObjectID))

	if err != nil {
		log.Printf("Could not create Transaction: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), err
}

func (model TransactionService) ProcessPayment(id primitive.ObjectID) {
	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}})
	if result == nil {
		errors.New("Could not find a Transaction")
		return
	}
	var transaction models.Transaction
	err := result.Decode(&transaction)
	if err != nil {
		return
	}

	paymentService := new(PaymentsService)
	params := PaymentParams{
		Amount:  transaction.Amount,
		OrderId: id.Hex(),
		UserId:  transaction.UserId,
	}

	if res := paymentService.ProcessPayment(params); res {
		getCollection().UpdateByID(context.TODO(), id, bson.D{bson.E{"status", models.TRANSACTION_STATUS_CONFIRMED}})
		return
	}

}

func (model TransactionService) FindOneWithUserId(id primitive.ObjectID, userId string) (*models.Transaction, error) {
	var transaction models.Transaction

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

func (model TransactionService) Cancel(id primitive.ObjectID) (*models.Transaction, error) {
	var transaction models.Transaction

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}})
	if result == nil {
		return nil, errors.New("Could not find an Transaction")
	}

	err := result.Decode(&transaction)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Transactions: %v", transaction)
	if transaction.Status != models.TRANSACTION_STATUS_CREATED {
		return nil, errors.New("Could not cancel Transaction if Transaction status is not CREATED")
	}
	var updatedDocument models.Transaction
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", models.TRANSACTION_STATUS_CANCELLED}}}}
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
