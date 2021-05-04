package forms

type Product struct {
	Name   string  `json:"name" binding:"required"`
	Amount int     `json:"amount" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
}

type CreateTransaction struct {
	Products []Product `json:"product" binding:"required"`
	Id       string    `json:"id"`
	Status   string    `json:"status"`
	UserId   string    `json:"userId"`
}
