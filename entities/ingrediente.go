package entity

import (
	"time"
)

//Ingrediente data
type Ingrediente struct {
	ID        ID
	Nome      string
	CreatedAt time.Time
}

//NewIngrediente create a new ingrediente
func NewIngrediente(nome string) (*Ingrediente, error) {
	if nome =="" {
		return nil, EmptyNomeError
	}

	i := &Ingrediente{
		ID:        NewID(),
		Nome:      nome,
		CreatedAt: time.Now(),
	}
	return i, nil
}
