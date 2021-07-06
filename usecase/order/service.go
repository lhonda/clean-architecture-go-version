package order

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

//CreateOrder Create an Order
func (s *Service) CreateOrder(customer string, pizzas []entity.Pizza) (*entity.Order, error) {
    e, err := entity.NewOrder(customer, pizzas)
    if err != nil {
        return nil, err
    }
    return s.repo.Create(e)
}

//GetOrder retrieves an Order given the ID
func (s *Service) GetOrder(id entity.ID) (*entity.Order, error) {
    return s.repo.Get(id)
}

//ListOrders List Orders
func (s *Service) ListOrders() ([]*entity.Order, error) {
    return s.repo.List()
}

//DeleteOrder Delete an Order
func (s *Service) DeleteOrder(id entity.ID) error {
    o, err := s.GetOrder(id)
    if o == nil {
        return entity.NotFoundError
    }
    if err != nil {
        return err
    }
    return s.repo.Delete(id)
}
