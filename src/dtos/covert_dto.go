package dtos

import (
	"src/models"
	"strconv"
)

func ReceiptDTOToReceipt(receiptDTO ReceiptDTO) (models.Receipt, error) {
	var receipt models.Receipt
	total, err := strconv.ParseFloat(receiptDTO.Total, 64)
	if err != nil {
		return models.Receipt{}, err
	}

	var items []models.Item
	for _, itemDTO := range receiptDTO.Items {
		item, err := ItemDTOToItem(itemDTO)
		if err != nil {
			return receipt, err
		}
		items = append(items, item)
	}

	receipt.Retailer = receiptDTO.Retailer
	receipt.PurchaseDate = receiptDTO.PurchaseDate
	receipt.PurchaseTime = receiptDTO.PurchaseTime
	receipt.Total = total
	receipt.Items = items

	return receipt, nil
}

func ItemDTOToItem(itemDTO ItemDTO) (models.Item, error) {
	var item models.Item
	price, err := strconv.ParseFloat(itemDTO.Price, 64)
	if err != nil {
		return models.Item{}, err
	}

	item.Price = price
	item.ShortDescription = itemDTO.ShortDescription

	return item, nil
}
