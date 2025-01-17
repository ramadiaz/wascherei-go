package dto

type TransactionInput struct {
	UserUUID      string  `json:"user_uuid" validate:"required,uuid4"`
	ProductUUID   string  `json:"product_uuid" validate:"required,uuid4"`
	UnitSize      float32 `json:"unit_size" validate:"required,number"`
	PaymentStatus string  `json:"payment_status" validate:"required"`
	Customer      string  `json:"customer" validate:"required"`
	CustomerPhone *string `json:"customer_phone"`
}
