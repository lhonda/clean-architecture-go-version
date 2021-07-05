package pedido

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Order, error)
	List() ([]*entity.Order, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.Order) (entity.ID, error)
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPedido(id entity.ID) (*entity.Order, error)
	SearchPedidos(query string) ([]*entity.Order, error)
	ListPedidos() ([]*entity.Order, error)
	CreatePedido(dono string, pizza *entity.Pizza) (entity.ID, error)
}
