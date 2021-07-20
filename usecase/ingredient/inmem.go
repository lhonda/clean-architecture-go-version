package ingredient

import (
	"sort"

	"github.com/lhonda/clean-architecture-go-version/entity"
)

//InMem in memory repo
type InMem struct {
	m map[entity.ID]*entity.Ingredient
}
type ByCreatedAt []*entity.Ingredient

func (a ByCreatedAt) Len() int           { return len(a) }
func (a ByCreatedAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCreatedAt) Less(i, j int) bool { return a[i].CreatedAt.Unix() < a[j].CreatedAt.Unix() }

//inMem create new repository
func inMem() *InMem {
	var m = map[entity.ID]*entity.Ingredient{}
	return &InMem{
		m: m,
	}
}

//Create an Ingredient
func (r *InMem) Create(e *entity.Ingredient) (*entity.Ingredient, error) {
	r.m[e.ID] = e
	return e, nil
}

//Get an Ingredient
func (r *InMem) Get(id entity.ID) (*entity.Ingredient, error) {
	if r.m[id] == nil {
		return nil, entity.NotFoundError
	}
	return r.m[id], nil
}

//List Ingredients
func (r *InMem) List() ([]*entity.Ingredient, error) {
	var d []*entity.Ingredient
	for _, j := range r.m {
		d = append(d, j)
	}
	sort.Sort(ByCreatedAt(d))
	return d, nil
}

// Delete Ingredient
func (r *InMem) Delete(id entity.ID) error {
	_, found := r.m[id]
	if found {
		delete(r.m, id)
		return nil
	}
	return entity.NotFoundError
}
