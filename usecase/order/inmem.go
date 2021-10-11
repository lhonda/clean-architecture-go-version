package order

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
	"sort"
)

//InMem in memory repo
type InMem struct {
	m map[entity.ID]*entity.Order
}

// ByCreatedAt struct used by Sort function
type ByCreatedAt []*entity.Order

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreatedAt.Unix() < a[j].CreatedAt.Unix() }

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
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//List orders
func (r *InMem) List() ([]*entity.Order, error) {
	var d []*entity.Order
	for _, j := range r.m {
		d = append(d, j)
	}
	sort.Sort(ByCreatedAt(d))
	return d, nil
}

// Delete order
func (r *InMem) Delete(id entity.ID) error {
	_, found := r.m[id]
	if found {
		delete(r.m, id)
		return nil
	}
	return entity.ErrNotFound
}
