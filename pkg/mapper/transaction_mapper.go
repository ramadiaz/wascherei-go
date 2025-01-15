package mapper

import (
	"wascherei-go/api/transactions/dto"
	"wascherei-go/models"

	"github.com/go-viper/mapstructure/v2"
)

func MapTransactionInputToModel(input dto.TransactionInput) models.Transactions {
	var transactionModel models.Transactions

	mapstructure.Decode(input, &transactionModel)
	return transactionModel
}

func MapTransactionModelToInput(transactionModel models.Transactions) dto.TransactionOutput {
	var output dto.TransactionOutput

	mapstructure.Decode(transactionModel, &output)
	return output
}
