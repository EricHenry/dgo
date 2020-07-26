package services

import "github.com/erichenry/dgo/domain/repositories"

type ItemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}
