package ingredient

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

//CreateIngredient Create an ingredient
func (s *Service) CreateIngredient(ingredient string ) (*entity.Ingredient, error) {
	e, err := entity.NewIngredient(ingredient)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(e)
}

//GetIngredient retrieves an Ingredient given the ID
func (s *Service) GetIngredient(id entity.ID) (*entity.Ingredient, error) {
	return s.repo.Get(id)
}

//DeleteIngredient Delete an Ingredient
func (s *Service) DeleteIngredient(id entity.ID) error {
	_, err := s.GetIngredient(id)
	if err != nil {
		return entity.NotFoundError
	}

	return s.repo.Delete(id)
}

//ListIngredients List Ingredients
func (s *Service) ListIngredients() ([]*entity.Ingredient, error) {
	return s.repo.List()
}
