package pizza

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//Service  interface
type Service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreatePizza Create a Pizza
func (s *Service) CreatePizza(name string, ingredients []entity.Ingredient) (*entity.Pizza, error) {
	e, err := entity.NewPizza(name, ingredients)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(e)
}

//GetPizza retrieves an Pizza given the ID
func (s *Service) GetPizza(id entity.ID) (*entity.Pizza, error) {
	return s.repo.Get(id)
}

//GetPizzaByName retrieves an Pizza given the ID
func (s *Service) GetPizzaByName(name string) (*entity.Pizza, error) {
	return s.repo.GetByName(name)
}

//ListPizzas List Pizzas
func (s *Service) ListPizzas() ([]*entity.Pizza, error) {
	return s.repo.List()
}

//DeletePizza Delete an Pizza
func (s *Service) DeletePizza(id entity.ID) error {
	return s.repo.Delete(id)
}
