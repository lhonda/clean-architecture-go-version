package pizza

import (
	"sort"

	"github.com/lhonda/clean-architecture-go-version/entity"
)

//InMem in memory repo
type InMem struct {
	m map[entity.ID]*entity.Pizza
}
// ByCreatedAt struct used by Sort function
type ByCreatedAt []*entity.Pizza

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreatedAt.Unix() < a[j].CreatedAt.Unix() }

//inMem create new repository
func inMem() *InMem {
	var m = map[entity.ID]*entity.Pizza{}
	return &InMem{
		m: m,
	}
}

//Create a Pizza
func (r *InMem) Create(e *entity.Pizza) (*entity.Pizza, error) {
	r.m[e.ID] = e
	return e, nil
}

//Get a Pizza
func (r *InMem) Get(id entity.ID) (*entity.Pizza, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetByName a Pizza
func (r *InMem) GetByName(name string) (*entity.Pizza, error) {
	for _, p := range r.m {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, entity.ErrNotFound
}

//List Pizzas
func (r *InMem) List() ([]*entity.Pizza, error) {
	var d []*entity.Pizza
	for _, j := range r.m {
		d = append(d, j)
	}
	sort.Sort(ByCreatedAt(d))
	return d, nil
}

// Delete Pizza
func (r *InMem) Delete(id entity.ID) error {
	_, found := r.m[id]
	if found {
		delete(r.m, id)
		return nil
	}
	return entity.ErrNotFound
}
