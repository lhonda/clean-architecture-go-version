package order

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

func TestCreateOrder(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    pizzas := newFixturePizza()
    o, err := m.CreateOrder("Dennis", pizzas)
    assert.Nil(t, err)
    assert.NotNil(t, o)
}

func TestCreateOrderWithEmptyCustomerShouldFail(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    pizzas := newFixturePizza()
    _, err := m.CreateOrder("", pizzas)
    assert.NotNil(t, err)
}

func TestListOrders(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    pizzas := newFixturePizza()

    u1, _ := m.CreateOrder("Dennis", pizzas)
    u2, _ := m.CreateOrder("Dennis", pizzas)

    t.Run("list all", func(t *testing.T) {
        all, err := m.ListOrders()

        assert.Nil(t, err)
        assert.Equal(t, 2, len(all))
        assert.Equal(t, u1.ID, all[0].ID)
        assert.Equal(t, u2.ID, all[1].ID)
    })
}

func TestGetOrder(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    pizzas := newFixturePizza()
    o, _ := m.CreateOrder("Dennis", pizzas)

    saved, _ := m.GetOrder(o.ID)

    assert.Equal(t, saved.ID, o.ID)
    assert.NotNil(t, saved.CreatedAt)
    assert.Equal(t, saved.Owner, o.Owner)
    assert.NotNil(t, saved.Pizzas)
    assert.Equal(t, saved.Owner, o.Owner)
}

func TestDeleteOrder(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    pizzas := newFixturePizza()
    o, _ := m.CreateOrder("Dennis", pizzas)

    error := m.DeleteOrder(o.ID)

    assert.Nil(t, error)
    all, error := m.ListOrders()
    assert.Nil(t, error)
    assert.Equal(t, 0, len(all))
}

func TestDeleteOrderWithNonExistingOrderShouldFail(t *testing.T) {
    repo := inMem()
    m := NewService(repo)
    nonExistentID := entity.NewID()
    error := m.DeleteOrder(nonExistentID)

    assert.NotNil(t, error)
}
