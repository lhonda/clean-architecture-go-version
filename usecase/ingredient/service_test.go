package ingredient

import (
    "github.com/lhonda/clean-architecture-go-version/entity"
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestCreateIngredient(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    o, err := m.CreateIngredient("ham")

    assert.Nil(t, err)
    assert.NotNil(t, o)
}

func TestCreateIngredientWithEmptyIngredientsShouldFail(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    _, err := m.CreateIngredient("")

    assert.NotNil(t, err)
    assert.EqualError(t, err, "Invalid empty name")
}

func TestListIngredients(t *testing.T) {
    repo := inMem()
    m := NewService(repo)

    u1, _ := m.CreateIngredient("cheese")
    u2, _ := m.CreateIngredient("ham")

    t.Run("list all", func(t *testing.T) {
        all, err := m.ListIngredients()

        assert.Nil(t, err)
        assert.Equal(t, 2, len(all))
        assert.Equal(t, u1.ID, all[0].ID)
        assert.Equal(t, u2.ID, all[1].ID)
    })
}

func TestGetIngredient(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    o, _ := m.CreateIngredient("ham")

    saved, _ := m.GetIngredient(o.ID)

    assert.Equal(t, saved.ID, o.ID)
    assert.NotNil(t, saved.CreatedAt)
    assert.NotNil(t, saved.Name)
}

func TestDeleteIngredient(t *testing.T) {
    repo := inMem()
    m := NewService(repo)

    o, _ := m.CreateIngredient("ham")

    error := m.DeleteIngredient(o.ID)

    assert.Nil(t, error)
    all, error := m.ListIngredients()
    assert.Nil(t, error)
    assert.Equal(t, 0, len(all))
}

func TestDeleteIngredientWithNonExistingIngredientShouldFail(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    nonExistentID := entity.NewID()
    error := m.DeleteIngredient(nonExistentID)

    assert.NotNil(t, error)
}
