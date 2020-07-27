package controllers

import (
	usecases "github.com/erichenry/dgo/use-cases"
)

type itemController struct {
	usecase usecases.ItemUsecase
}

func NewItemController(usecase usecases.ItemUsecase) (controller *itemController, err error) {
	controller = &itemController{
		usecase: usecase,
	}
	return controller, nil
}

type NewItemRequest struct {
	Name        string
	Description string
	Price       float32
}

type NewItemResponse struct {
	ID          string
	Name        string
	Description string
	Price       float32
}

func (i *itemController) NewItem(req NewItemRequest) (res *NewItemResponse, err error) {
	item, err := i.usecase.SaveItem(req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}
	return &NewItemResponse{
		ID:          item.ID.String(),
		Description: item.Description,
		Name:        item.Name,
		Price:       item.Price,
	}, nil
}
