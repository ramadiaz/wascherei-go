package mapper

import (
	"wascherei-go/api/products/dto"
	"wascherei-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapProductInputToModel(input dto.ProductInput) models.Products {
	var productModel models.Products

	mapstructure.Decode(input, &productModel)
	return productModel
}

func MapProductModelToInput(productModel models.Products) dto.ProductOutput {
	var output dto.ProductOutput

	mapstructure.Decode(productModel, &output)
	return output
}