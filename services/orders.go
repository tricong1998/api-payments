package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrdersService struct{}

func (ordersService OrdersService) ConfirmOrder(orderId string) {

}

func (ordersService OrdersService) CancelOrder(orderId string) {

}

func (ordersService OrdersService) callOrderApi(websiteURL string, c *gin.Context) {

	client := &http.Client{
		// Set timeout to abort if the request takes too long
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", websiteURL, nil)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err})
	}

	// Make website request call
	resp, err := client.Do(request)

	// If we have a successful request
	if resp.StatusCode == 200 {

	}
}
