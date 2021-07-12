package pizza

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Pizza, error)
	List() ([]*entity.Pizza, error)
}

//Writer Pizza writer
type Writer interface {
	Create(e *entity.Pizza) (*entity.Pizza, error)
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPizza(id entity.ID) (*entity.Pizza, error)
	ListPizzas() ([]*entity.Pizza, error)
	CreatePizza(pizza []entity.Ingredient) (*entity.Pizza, error)
	DeletePizza(entity.ID) error
}
