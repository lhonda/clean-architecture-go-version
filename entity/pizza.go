package entity

import (
	"time"
)

//Pizza data
type Pizza struct {
	ID           ID
	Ingredientes []Ingrediente
	CreatedAt    time.Time
}

//NewPizza create a new pizza
func NewPizza(ingredientes []Ingrediente) (*Pizza, error) {
	p := &Pizza{
		ID:           NewID(),
		Ingredientes: ingredientes,
		CreatedAt:    time.Now(),
	}
	return p, nil
}
