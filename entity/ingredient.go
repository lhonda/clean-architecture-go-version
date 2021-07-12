package entity

import (
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
        return nil, EmptyNomeError
    }

    i := &Ingredient{
        ID:        NewID(),
        Name:      name,
        CreatedAt: time.Now(),
    }
    return i, nil
}
