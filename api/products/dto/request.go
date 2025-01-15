package dto

type ProductInput struct {
	UserUUID string `json:"user_uuid" validate:"required,uuid4"`
	Type     string `json:"type" validate:"required"`
	Duration uint   `json:"duration" validate:"required,number"`
	Price    uint   `json:"price" validate:"required,number"`
	Unit     string `json:"unit" validate:"required"`
}
