package ingredient

import (
	"github.com/lhonda/clean-architecture-go-version/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Ingredient, error)
	List() ([]*entity.Ingredient, error)
}

//Writer Ingredient writer
type Writer interface {
	Create(e *entity.Ingredient) (*entity.Ingredient, error)
	Delete(id entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetIngredient(id entity.ID) (*entity.Ingredient, error)
	ListIngredients() ([]*entity.Ingredient, error)
	CreateIngredient(ingredient string) (*entity.Ingredient, error)
	DeleteIngredient(entity.ID) error
}
