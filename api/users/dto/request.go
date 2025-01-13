package dto

type UserInput struct {
	Username        string `json:"username" validate:"required,min=4"`
	Password        string `json:"password" validate:"required,min=8"`
	BusinessName    string `json:"business_name" validate:"required"`
	BusinessOwner   string `json:"business_owner" validate:"required"`
	BusinessPhone   string `json:"business_phone" validate:"required,e164"`
	BusinessAddress string `json:"business_address" validate:"required"`
	BusinessLogo    string `json:"business_logo" validate:"required,url"`
	Slogan          string `json:"slogan" validate:"required"`
}

type Login struct {
	Username string `validate:"required,min=4"`
	Password string `validate:"required,min=8"`
}
