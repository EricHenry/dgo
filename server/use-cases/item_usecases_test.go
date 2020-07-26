package usecases

import (
	"testing"

	"github.com/erichenry/dgo/domain"
	"github.com/google/uuid"
)

// Mocks
type mockRepo struct{}

func (m *mockRepo) Save(i *domain.Item) error {
	return nil
}

func (m *mockRepo) FindById(id uuid.UUID) (item *domain.Item, err error) {
	return &domain.Item{
		ID: id,
	}, nil
}

func (m *mockRepo) FindByName(name string) (item *domain.Item, err error) {
	return &domain.Item{
		Name: name,
	}, nil
}

func (m *mockRepo) FindAll() (items []*domain.Item, err error) {
	return []*domain.Item{}, nil
}

// Tests
func TestItemUsecase(t *testing.T) {
	repo := &mockRepo{}

	usecase := NewItemUsecase(repo)
	_, err := usecase.SaveItem("Chex mix", "Chex mix description", 10.00)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

}
