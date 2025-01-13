package mapper

import (
	"wascherei-go/api/users/dto"
	"wascherei-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapUserInputToModel(input dto.UserInput) models.Users {
	var userModel models.Users

	mapstructure.Decode(input, &userModel)
	return userModel
}
