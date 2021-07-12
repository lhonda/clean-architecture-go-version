package pizza

import (
	"testing"
	"time"

	"github.com/lhonda/clean-architecture-go-version/entity"

	"github.com/stretchr/testify/assert"
)

func newFixturePizza() []entity.Pizza {
	cheese := entity.Ingredient{
		ID:        entity.NewID(),
		Name:      "cheese",
		CreatedAt: time.Now(),
	}

	ham := entity.Ingredient{
		ID:        entity.NewID(),
		Name:      "ham",
		CreatedAt: time.Now(),
	}

	p, _ := entity.NewPizza([]entity.Ingredient{ham, cheese})
	p2, _ := entity.NewPizza([]entity.Ingredient{cheese})

	return []entity.Pizza{*p, *p2}
}

func newFixtureIngredients() []entity.Ingredient {
	cheese := entity.Ingredient{
		ID:        entity.NewID(),
		Name:      "cheese",
		CreatedAt: time.Now(),
	}

	ham := entity.Ingredient{
		ID:        entity.NewID(),
		Name:      "ham",
		CreatedAt: time.Now(),
	}

	return []entity.Ingredient{cheese, ham}
}

func TestCreatePizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()
	o, err := m.CreatePizza(ingredients)

	assert.Nil(t, err)
	assert.NotNil(t, o)
}

func TestCreatePizzaWithEmptyIngredientsShouldFail(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	_, err := m.CreatePizza(nil)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Empty Ingredients list")
}

func TestListPizzas(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()

	u1, _ := m.CreatePizza(ingredients)
	u2, _ := m.CreatePizza(ingredients)

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListPizzas()

		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
		assert.Equal(t, u1.ID, all[0].ID)
		assert.Equal(t, u2.ID, all[1].ID)
	})
}

func TestGetpizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()
	o, _ := m.CreatePizza(ingredients)

	saved, _ := m.GetPizza(o.ID)

	assert.Equal(t, saved.ID, o.ID)
	assert.NotNil(t, saved.CreatedAt)
	assert.NotNil(t, saved.Ingredients)
}

func TestDeletepizza(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	ingredients := newFixtureIngredients()

	o, _ := m.CreatePizza(ingredients)

	error := m.DeletePizza(o.ID)

	assert.Nil(t, error)
	all, error := m.ListPizzas()
	assert.Nil(t, error)
	assert.Equal(t, 0, len(all))
}

func TestDeletepizzaWithNonExistingpizzaShouldFail(t *testing.T) {
	repo := inMem()
	m := NewService(repo)
	nonExistentID := entity.NewID()
	error := m.DeletePizza(nonExistentID)

	assert.NotNil(t, error)
}
