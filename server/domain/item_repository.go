package domain

import (
	"github.com/google/uuid"
)

type ItemRepository interface {
	FindAll() ([]*Item, error)
	FindById(id uuid.UUID) (*Item, error)
	FindByName(name string) (*Item, error)
	Save(*Item) error
}
