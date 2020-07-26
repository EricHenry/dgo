package repositories

import (
	"github.com/erichenry/dgo/domain/entities"
	"github.com/google/uuid"
)

type ItemRepository interface {
	FindAll() ([]*entities.Item, error)
	FindById(id uuid.UUID) (*entities.Item, error)
	FindByName(name string) (*entities.Item, error)
	Save(*entities.Item) error
}
