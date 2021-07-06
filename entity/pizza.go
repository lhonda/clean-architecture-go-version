package entity

import (
	"time"
)

//Pizza data
type Pizza struct {
	ID          ID
	Ingredients []Ingredient
	CreatedAt   time.Time
}

//NewPizza create a new pizza
func NewPizza(ingredients []Ingredient) (*Pizza, error) {
	p := &Pizza{
		ID:          NewID(),
		Ingredients: ingredients,
		CreatedAt:   time.Now(),
	}
	return p, nil
}
