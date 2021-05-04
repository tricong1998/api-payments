package services

import (
	"api-payments/models"
	"math/rand"
	"time"
)

type PaymentsService struct{}

func (paymentsService PaymentsService) ProcessPayment(transaction models.Transaction) {
	order := new(OrdersService)

	if result := paymentsService.handlePaymentLogic(transaction); result {
		go order.ConfirmOrder(transaction.OrderId)
		return
	}

	go order.CancelOrder(transaction.OrderId)
	return
}

func (paymentsService PaymentsService) handlePaymentLogic(transaction models.Transaction) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(3)
	return number%2 == 0
}
