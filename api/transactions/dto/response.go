package dto

import userDTO "wascherei-go/api/users/dto"
import productDTO "wascherei-go/api/products/dto"

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}

type TransactionOutput struct {
	UUID          string                   `json:"uuid"`
	UserUUID      string                   `json:"user_uuid"`
	ProductUUID   string                   `json:"product_uuid"`
	UnitSize      float32                  `json:"unit_size"`
	TotalPrice    uint                     `json:"total_price"`
	PaymentStatus string                   `json:"payment_status"`
	Customer      string                   `json:"customer"`
	CustomerPhone *string                  `json:"customer_phone"`
	User          userDTO.UserOutput       `json:"user,omitempty"`
	Product       productDTO.ProductOutput `json:"product,omitempty"`
}
