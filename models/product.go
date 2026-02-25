package models

import "fmt"

type Product struct {
	ID int
	Name string
	Description string
	Price float64
	stock int  //privado encapsulado
}

//Setter de stock con validación
func (p *Product) SetStock(s int) error {
	if s < 0 {
		return fmt.Errorf("el stock no puede ser negativo")
	}
	p.stock = s
	return nil
}

//Getter del stock
func (p *Product) GetStock() int {
	return p.stock
}	

//Reducir stock al realizar una compra
func (p *Product) ReduceStock(s int) error {
	if s > p.stock {
		return fmt.Errorf("stock insuficiente")
	}			
	p.stock -= s
	return nil
}	

//Constructor para crear un nuevo producto con validaciones
func NewProduct(id int, name, description string, price float64, stock int) (*Product, error) {
	if price <= 0 {
		return nil, fmt.Errorf("el precio debe ser mayor a 0")
	}

	p := &Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}

	if err := p.SetStock(stock); err != nil {
		return nil, err
	}

	return p, nil
}