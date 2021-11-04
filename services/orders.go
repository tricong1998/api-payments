package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OrdersService struct{}

func (ordersService OrdersService) CancelOrder(orderId string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	orderUrl := os.Getenv("API_ORDERS_HOST") + os.Getenv("API_ORDERS_PORT")
	cancelOrderUrl := orderUrl + "/orders/" + orderId + "/cancel"
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(cancelOrderUrl, "application/json", nil)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Printf(sb)

	return nil
}

func (ordersService OrdersService) ConfirmOrder(orderId string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	orderUrl := os.Getenv("API_ORDERS_HOST") + os.Getenv("API_ORDERS_PORT")
	cancelOrderUrl := orderUrl + "/orders/" + orderId + "/confirm"
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(cancelOrderUrl, "application/json", nil)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Printf(sb)

	return nil
}
