package service

import (
	"CAKE-STORE/entity"
	"CAKE-STORE/repository"
	"context"
)

type (
	cakeService struct {
		CakeRepository repository.CakeRepository
	}

	CakeService interface {
		ListCake(ctx context.Context, title string) ([]entity.ListCake, error)
		DetailCake(ctx context.Context, id int) (entity.Cake, error)
		AddNewCake(ctx context.Context, cake entity.Cake) error
		UpdateCake(ctx context.Context, id int, cake entity.Cake) error
		DeleteCake(ctx context.Context, id int) error
	}
)

func NewCakeService(cakeRepository *repository.CakeRepository) CakeService {
	return &cakeService{
		CakeRepository: *cakeRepository,
	}
}

func (cs *cakeService) DetailCake(ctx context.Context, id int) (entity.Cake, error) {
	DetailCake, err := cs.CakeRepository.GetDetailCake(ctx, id)
	if err != nil {
		return DetailCake, err
	}

	return DetailCake, nil
}

func (cs *cakeService) ListCake(ctx context.Context, title string) ([]entity.ListCake, error) {
	ListCake, err := cs.CakeRepository.GetListCake(ctx, title)
	if err != nil {
		return nil, err
	}

	return ListCake, nil
}

func (cs *cakeService) AddNewCake(ctx context.Context, cake entity.Cake) error {
	err := cs.CakeRepository.InsertCake(ctx, cake)
	if err != nil {
		return err
	}
	return nil
}

func (cs *cakeService) UpdateCake(ctx context.Context, id int, cake entity.Cake) error {
	err := cs.CakeRepository.UpdateCake(ctx, id, cake)
	if err != nil {
		return err
	}

	return nil
}

func (cs *cakeService) DeleteCake(ctx context.Context, id int) error {
	err := cs.CakeRepository.DeleteCake(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
