package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewPizza(t *testing.T) {

	ingredients := []Ingredient{"queijo", "calabreza"}
	pizza, _ := NewPizza("queijo", ingredients)

	pizzas := []Pizza{*pizza}
	order, err := NewOrder("Jukinha", pizzas)

	assert.Nil(t, err)
	assert.Equal(
		t, order.Pizzas[0].ID, pizza.ID)
	assert.NotNil(t, order.ID)
}

func TestCreateNewPizzaWithEmptyCustomerShouldFail(t *testing.T) {
	ingredients := []Ingredient{"queijo", "calabreza"}
	pizza, _ := NewPizza("queijo", ingredients)

	pizzas := []Pizza{*pizza}
	o, err := NewOrder("", pizzas)

	assert.NotNil(t, err)
	assert.Nil(t, o)
}
