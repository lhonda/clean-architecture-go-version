package entity

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
	"time"
)

//Ingredient data
type Ingredient struct {
	ID        ID
	Name      string
	CreatedAt time.Time
}

//NewIngredient create a new ingredient
func NewIngredient(name string) (*Ingredient, error) {
	if name == "" {
		return nil, EmptyNameError
	}

	i := &Ingredient{
		ID: entity.NewID(),
		Name: name,
		CreatedAt: time.Now(),
	}
	return i, nil
}
