package entity

import "time"

// Order -
type Order struct {
	ID        ID
	Pizza     []Pizza
	Owner     string
	CreatedAt time.Time
}

// New NewOrder create a new order
func NewOrder(owner string, pizza []Pizza) (*Order, error) {

	if owner == "" {
		return nil, EmptyOwnerError
	}

	return &Order{
		ID:        NewID(),
		Pizza:     pizza,
		CreatedAt: time.Now(),
	}, nil
}
