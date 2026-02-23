package services

import "metazone/models"

var orders []models.Order

func CreateOrder(userID int, products []models.Product) models.Order {
	var total float64
	for _, p := range products {
		total += p.Price
	}

	order := models.Order{
		ID:       len(orders) + 1,
		UserID:   userID,
		Products: products,
		Total:    total,
	}

	orders = append(orders, order)
	return order
}