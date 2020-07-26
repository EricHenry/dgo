package repositories

import "github.com/erichenry/dgo/domain/entities"

type UserRepository interface {
	FindAll() ([]*entities.User, error)
	FindByID(id string) (*entities.User, error)
	FindByName(name string) (*entities.User, error)
	Save(*entities.User) error
}
