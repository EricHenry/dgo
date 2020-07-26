package usecases

import (
	"github.com/erichenry/dgo/domain"

	"github.com/google/uuid"
)

type ItemUsecase interface {
	ListItems() ([]*domain.Item, error)
	SaveItem(name, description string, price float32) (domain.Item, error)
}

type itemUsecase struct {
	repo domain.ItemRepository
}

func NewItemUsecase(repo domain.ItemRepository) *itemUsecase {
	return &itemUsecase{
		repo: repo,
	}
}

func (i *itemUsecase) ListItems() (items []*domain.Item, err error) {
	items, err = i.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemUsecase) SaveItem(name, description string, price float32) (item *domain.Item, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	item = domain.NewItem(id, name, description, price)
	err = i.repo.Save(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}
