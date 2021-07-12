package entity

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestNewPizzaShouldSucceed(t *testing.T) {
    queijo, _ := NewIngredient("queijo")
    calabreza, _ := NewIngredient("calabreza")

    ingredients := []Ingredient{*queijo, *calabreza}
    p, err := NewPizza(ingredients)
    assert.Nil(t, err)
    assert.Equal(t, p.Ingredients[0].Name, "queijo")
    assert.NotNil(t, p.ID)
}
