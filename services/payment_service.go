package services

import (
	"metazone/models"
	"errors")

var payments []models.Payment

func CreatePayment(orderID int, amount float64, method string) (*models.Payment, error) {
	
	if amount <= 0 {
		return nil, errors.New("El monto del pago debe ser mayor a cero")
	}

	payment := models.Payment{
		ID:      len(payments) + 1,
		OrderID: orderID,
		Amount:  amount,
		Method:  method,
		Status:  "PENDIENTE",
	}

	// Procesar el pago
	payment.Process()
	payments = append(payments, payment)
	return &payment, nil
}

//Listar todos los pagos
func GetPayments() []models.Payment {
	return payments
}
