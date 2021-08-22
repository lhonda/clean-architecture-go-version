package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPizzaShouldSucceed(t *testing.T) {

	ingredients := []Ingredient{"queijo", "calabreza"}
	p, err := NewPizza("queijo", ingredients, NewID())
	assert.Nil(t, err)
	assert.Equal(t, p.Ingredients[0], Ingredient("queijo"))
	assert.NotNil(t, p.ID)
}
