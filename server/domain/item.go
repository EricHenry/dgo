package domain

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float32
}

// Constructor
func NewItem(id uuid.UUID, name, description string, price float32) *Item {
	return &Item{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

// Item Accessors
func (i *Item) GetID() uuid.UUID {
	return i.ID
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetDescription() string {
	return i.Description
}

func (i *Item) GetPrice() float32 {
	return i.Price
}
