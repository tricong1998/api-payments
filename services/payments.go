package services

import (
	"math/rand"
	"time"
)

type PaymentsService struct{}

type PaymentParams struct {
	Amount  float32
	UserId  string
	OrderId string
}

func (paymentsService PaymentsService) ProcessPayment(transaction PaymentParams) bool {
	return paymentsService.handlePaymentLogic(transaction)
}

func (paymentsService PaymentsService) handlePaymentLogic(transaction PaymentParams) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(3)
	return number%2 == 0
}
