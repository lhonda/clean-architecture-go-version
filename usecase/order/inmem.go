package order

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//InMem in memory repo
type InMem struct {
	m map[entity.ID]*entity.Order
}

//inMem create new repository
func inMem() *InMem {
	var m = map[entity.ID]*entity.Order{}
	return &InMem{
		m: m,
	}
}

//Create an Order
func (r *InMem) Create(e *entity.Order) (*entity.Order, error) {
	r.m[e.ID] = e
	return e, nil
}

//Get an order
func (r *InMem) Get(id entity.ID) (*entity.Order, error) {
	if r.m[id] == nil {
		return nil, entity.NotFoundError
	}
	return r.m[id], nil
}

//List orders
func (r *InMem) List() ([]*entity.Order, error) {
	var d []*entity.Order
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

// Delete order
func (r *InMem) Delete(id entity.ID) error {
	_, found := r.m[id]
	if found {
		delete(r.m, id)
		return nil
	}
	return entity.NotFoundError
}