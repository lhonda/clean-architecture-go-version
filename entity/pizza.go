package entity

import (
	"strings"
	"time"
)

//Pizza data
type Pizza struct {
	ID          ID
	Name        string
	Ingredients []Ingredient
	Order       ID
	CreatedAt   time.Time
}

//NewPizza create a new pizza
func NewPizza(name string, ingredients []Ingredient, orderId ID) (*Pizza, error) {
	if ingredients == nil {
		return nil, EmptyIngredientsListError
	}

	p := &Pizza{
		ID:          NewID(),
		Name:        name,
		Ingredients: ingredients,
		Order:       orderId,
		CreatedAt:   time.Now(),
	}
	return p, nil
}

// GetIngredientsAsString function
func (p *Pizza) GetIngredientsAsString() string {
	var acc []string
	for _, ingredient := range p.Ingredients {
		acc = append(acc, string(ingredient))
	}

	return strings.Join(acc, ";")
}

// SetIngredientsAsList function
func (p *Pizza) SetIngredientsAsList(ingredients string) *Pizza {
	for _, ingredient := range strings.Split(ingredients, ";") {
		p.Ingredients = append(p.Ingredients, Ingredient(ingredient))
	}
	return p
}
