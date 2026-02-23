package services

import "metazone/models"

// Lista de productos en memoria
var products []*models.Product

// Crear un nuevo producto
func CreateProduct(name, desc string, price float64, stock int) (*models.Product, error) {
	product, err := models.NewProduct(
		len(products)+1,
		name,
		desc,
		price,
		stock,
	)
	if err != nil {
		return nil, err
	}

	products = append(products, product)
	return product, nil
}

// Obtener todos los productos
func GetProducts() []*models.Product {
	return products
}