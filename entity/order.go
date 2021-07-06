package entity

import "time"

// Order -
type Order struct {
    ID        ID
    Pizzas    []Pizza
    Owner     string
    CreatedAt time.Time
}

// NewOrder creates a new order
func NewOrder(owner string, pizzas []Pizza) (*Order, error) {

    if owner == "" {
        return nil, EmptyOwnerError
    }

    return &Order{
        ID:        NewID(),
        Owner:     owner,
        Pizzas:    pizzas,
        CreatedAt: time.Now(),
    }, nil
}
