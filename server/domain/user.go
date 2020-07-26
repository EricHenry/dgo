package domain

type User struct {
	ID     string
	Name   string
	Online bool
}

// Constructor
func NewUser(id, name string, online bool) *User {
	return &User{
		ID:     id,
		Name:   name,
		Online: online,
	}
}

// Item Accessors
func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetOnline() bool {
	return u.Online
}
