package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewPizza(t *testing.T) {
	queijo, _ := NewIngrediente("queijo")
	calabreza, _ := NewIngrediente("calabreza")

	ingredientes := []Ingrediente{*queijo, *calabreza}
	pizza, _ := NewPizza(ingredientes)

	pizzas := []Pizza{*pizza}
	order, err := NewOrder("Jukinha", pizzas)

	assert.Nil(t, err)
	assert.Equal(
		t, order.Pizza[0], pizza)
	assert.NotNil(t, order.ID)
}
