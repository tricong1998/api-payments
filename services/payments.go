package services

import (
	"math/rand"
	"time"
)

type PaymentsService struct{}

type PaymentTransaction struct {
	OrderId string
	Amount  float32
}

func (paymentsService PaymentsService) ProcessPayment(transaction PaymentTransaction) {
	order := new(OrdersService)

	if result := paymentsService.handlePaymentLogic(transaction); result {
		go order.ConfirmOrder(transaction.OrderId)
		return
	}

	go order.CancelOrder(transaction.OrderId)
	return
}

func (paymentsService PaymentsService) handlePaymentLogic(transaction PaymentTransaction) bool {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(3)
	return number%2 == 0
}
