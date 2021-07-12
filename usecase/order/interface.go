package order

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Order, error)
	List() ([]*entity.Order, error)
}

//Writer Order writer
type Writer interface {
	Create(e *entity.Order) (*entity.Order, error)
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetOrder(id entity.ID) (*entity.Order, error)
	ListOrders() ([]*entity.Order, error)
	CreateOrder(customer string, pizza []entity.Pizza) (*entity.Order, error)
	DeleteOrder(entity.ID) error
}
