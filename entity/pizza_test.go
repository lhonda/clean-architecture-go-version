package entity

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPizzaShouldSucceed(t *testing.T) {

	ingredients := []Ingredient{"queijo", "calabreza"}
	p, err := NewPizza("queijo", ingredients)
	assert.Nil(t, err)
	assert.Equal(t, p.Ingredients[0], Ingredient("queijo"))
	assert.NotNil(t, p.ID)
}

func TestStringToID(t *testing.T) {
	expected, _ := uuid.NewUUID()
	id, _ := StringToID(expected.String())

	assert.Equal(t, id, expected)
}

func TestEmptyIngredients(t *testing.T) {
	p, err := NewPizza("queijo", []Ingredient{})

	assert.Nil(t, p)
	assert.ErrorIs(t, err, EmptyIngredientsListError)
}

func TestGetSetIngredients(t *testing.T) {
	ingredients := []Ingredient{"queijo", "calabreza"}
	p, _ := NewPizza("queijo", ingredients)

	p.SetIngredientsAsList("queijo;presunto;tomate")

	assert.Equal(t, p.Ingredients, []Ingredient{"queijo", "presunto", "tomate"})

	assert.Equal(t, "queijo;presunto;tomate", p.GetIngredientsAsString())
}
