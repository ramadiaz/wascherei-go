package dto

type TransactionInput struct {
	UserUUID      string  `json:"user_uuid" validate:"required,uuid4"`
	ProductUUID   string  `json:"product_uuid" validate:"required,uuid4"`
	UnitSize      uint    `json:"unit_size" validate:"required,number"`
	TotalPrice    uint    `json:"total_price" validate:"required,number"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	Customer      string  `json:"customer" validate:"required"`
	CustomerPhone *string `json:"customer_phone"`
}
