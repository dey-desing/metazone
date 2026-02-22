package models

type Payment struct {
	ID      int
	OrderID int
	Amount  float64
	Method  string
	Status  string
	prueba  string
	internalID string
	runerID int	
}
