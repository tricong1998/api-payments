package controllers

import (
	"api-payments/forms"
	"api-payments/models"
	"api-payments/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BackendTransactionController struct{}

var transactionModel = new(models.Transaction)

//
// @ID api-payments-create-transaction
// @Summary Create an transaction
// @Description Create an transaction with status Created
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param user body forms.CreateTransaction true "Add Product"
// @Success 200 {string} string	"id"
// @Router /transactions [post]
// @Security ApiKeyAuth
func (controller BackendTransactionController) Create(c *gin.Context) {
	var input forms.CreateTransaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := transactionModel.Create(input, user.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
	return
}

// @Summary Find an transaction
// @ID api-payments-read-transaction
// @Description Get an created transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Transaction
// @Router /transactions/{id} [get]
// @Security ApiKeyAuth
func (controller BackendTransactionController) FindOne(c *gin.Context) {
	user, _ := c.MustGet("User").(services.User)

	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	transactionId, _ := primitive.ObjectIDFromHex(id)
	transaction, err := transactionModel.FindOneWithUserId(transactionId, user.UserId)

	if err != nil {
		if transaction == nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transaction})
	return
}

// @Summary Cancel an transaction
// @ID api-payments-cancel-transaction
// @Description Get an created transaction
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Transaction
// @Router /transactions/{id}/cancel [post]
// @Security ApiKeyAuth
func (controller BackendTransactionController) Cancel(c *gin.Context) {
	user, _ := c.MustGet("User").(services.User)

	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	transactionId, _ := primitive.ObjectIDFromHex(id)

	transaction, err := transactionModel.Cancel(transactionId, user.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": transaction})
	return
}
