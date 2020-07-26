package usecases

import (
	"fmt"
	"testing"

	"github.com/erichenry/dgo/domain/entities"
	"github.com/erichenry/dgo/domain/services"
	"github.com/google/uuid"
)

// Mocks
type mockRepo struct{}

func (m *mockRepo) Save(i *entities.Item) error {
	return nil
}

func (m *mockRepo) FindById(id uuid.UUID) (item *entities.Item, err error) {
	return &entities.Item{
		ID: id,
	}, nil
}

func (m *mockRepo) FindByName(name string) (item *entities.Item, err error) {
	return &entities.Item{
		Name: name,
	}, nil
}

func (m *mockRepo) FindAll() (items []*entities.Item, err error) {
	return []*entities.Item{}, nil
}

// Tests
func TestItemUsecase(t *testing.T) {
	repo := &mockRepo{}
	service := services.NewItemService(repo)

	usecase := NewItemUsecase(repo, service)
	item, err := usecase.SaveItem("Chex mix", "Chex mix description", 10.00)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

}
