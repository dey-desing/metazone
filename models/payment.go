package models

type Payment struct {
	ID int
	OrderID int
	Amount float64
	Method string
	Status string
	
}

// Procesar el pago
func (p *Payment) Process() {
	p.Status = "Pago procesado"
}