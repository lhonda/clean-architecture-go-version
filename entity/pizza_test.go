package entity

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewPizzaShouldSucceed(t *testing.T) {
    queijo, _ := NewIngrediente("queijo")
    calabreza, _ := NewIngrediente("calabreza")

    ingredientes := []Ingrediente{*queijo, *calabreza}
    p, err := NewPizza(ingredientes)
    assert.Nil(t, err)
    assert.Equal(t, p.Ingredientes[0].Nome, "queijo")
    assert.NotNil(t, p.ID)
}
