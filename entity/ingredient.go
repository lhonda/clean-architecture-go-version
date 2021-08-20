package entity

import "time"

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
		Name: name,
	}
	return i, nil
}
