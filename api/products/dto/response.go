package dto

type ProductOutput struct {
	UUID     string `json:"uuid"`
	UserUUID string `json:"user_uuid"`
	Type     string `json:"type"`
	Duration uint   `json:"duration"`
	Price    uint   `json:"price"`
	Unit     string `json:"unit"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
}