package pizza

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newFixtureIngredients() []entity.Ingredient {
	return []entity.Ingredient{"cheese", "ham"}
}

func TestCreatePizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()
	o, err := m.CreatePizza("queijo", ingredients)

	assert.Nil(t, err)
	assert.NotNil(t, o)
}

func TestCreatePizzaWithEmptyIngredientsShouldFail(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	_, err := m.CreatePizza("queijo", nil)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Empty ingredients list")
}

func TestListPizzas(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()

	u1, _ := m.CreatePizza("queijo", ingredients)
	u2, _ := m.CreatePizza("queijo", ingredients)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPizzas()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
		assert.Equal(t, u1.ID, all[0].ID)
		assert.Equal(t, u2.ID, all[1].ID)
	})
}

func TestGetPizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()
	o, _ := m.CreatePizza("queijo", ingredients)

	saved, _ := m.GetPizza(o.ID)

	assert.Equal(t, saved.ID, o.ID)
	assert.NotNil(t, saved.CreatedAt)
	assert.NotNil(t, saved.Ingredients)
}

func TestDeletePizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()
	o, _ := m.CreatePizza("queijo", ingredients)

	err := m.DeletePizza(o.ID)

	assert.Nil(t, err)
	all, err := m.ListPizzas()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(all))
}

func TestDeletepizzaWithNonExistingpizzaShouldFail(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	nonExistentID := entity.NewID()
	err := m.DeletePizza(nonExistentID)

	assert.NotNil(t, err)
}
