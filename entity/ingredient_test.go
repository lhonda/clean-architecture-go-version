package entity

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewIngredient(t *testing.T) {
    q, err := NewIngredient("queijo")
    assert.Nil(t, err)
    assert.Equal(t, q.Name, "queijo")
    assert.NotNil(t, q.ID)
    assert.NotNil(t, q.CreatedAt)
}

func TestNewIngredientWithEmptyParameterShouldFail(t *testing.T) {
    _, err := NewIngredient("")
    assert.NotNil(t, err)
}
