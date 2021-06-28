package entity

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestNewIngrediente(t *testing.T) {
    q, err := NewIngrediente("queijo")
    assert.Nil(t, err)
    assert.Equal(t, q.Nome, "queijo")
    assert.NotNil(t, q.ID)
    assert.NotNil(t, q.CreatedAt)
}

func TestNewIngredienteWithEmptyParameterShouldFail(t *testing.T) {
    _, err := NewIngrediente("")
    assert.NotNil(t, err)
}
