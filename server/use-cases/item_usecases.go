package usecases

import (
	"github.com/erichenry/dgo/domain/entities"
	"github.com/erichenry/dgo/domain/repositories"
	"github.com/erichenry/dgo/domain/services"

	"github.com/google/uuid"
)

type ItemUsecase interface {
	ListItems() ([]*entities.Item, error)
	SaveItem(*entities.Item) error
}

type itemUsecase struct {
	repo    repositories.ItemRepository
	service *services.ItemService
}

func NewItemUsecase(repo repositories.ItemRepository, service *services.ItemService) *itemUsecase {
	return &itemUsecase{
		repo:    repo,
		service: service,
	}
}

func (i *itemUsecase) ListItems() (items []*entities.Item, err error) {
	items, err = i.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemUsecase) SaveItem(name, description string, price float32) (item *entities.Item, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	item = entities.NewItem(id, name, description, price)
	err = i.repo.Save(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
