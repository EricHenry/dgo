package domain

type UserRepository interface {
	FindAll() ([]*User, error)
	FindByID(id string) (*User, error)
	FindByName(name string) (*User, error)
	Save(*User) error
}
