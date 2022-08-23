package service

import (
	"CAKE-STORE/entity"
	"CAKE-STORE/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestListCake(t *testing.T) {
	var ctx context.Context
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockCakeRepository(mockCtrl)
	cs := cakeService{CakeRepository: mockRepo}

	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().GetListCake(gomock.Any(), gomock.Any()).Return([]entity.ListCake{}, nil)
		l, err := cs.ListCake(ctx, "title")
		assert.Nil(t, err)
		assert.NotNil(t, l)
	})

	t.Run("Repository Error", func(t *testing.T) {
		mockRepo.EXPECT().GetListCake(gomock.Any(), gomock.Any()).Return(nil, errors.New("ERROR"))
		l, err := cs.ListCake(ctx, "")
		assert.Nil(t, l)
		assert.Error(t, err)
	})
}

func TestDetailCake(t *testing.T) {
	var ctx context.Context
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockCakeRepository(mockCtrl)
	cs := cakeService{CakeRepository: mockRepo}

	t.Run("Succes", func(t *testing.T) {
		mockRepo.EXPECT().GetDetailCake(gomock.Any(), gomock.Any()).Return(entity.Cake{}, nil)
		d, e := cs.DetailCake(ctx, 1)
		assert.Nil(t, e)
		assert.NotNil(t, d)
	})

	t.Run("Repository error", func(t *testing.T) {
		mockRepo.EXPECT().GetDetailCake(gomock.Any(), gomock.Any()).Return(entity.Cake{}, errors.New("ERROR"))
		_, e := cs.DetailCake(ctx, 1)
		assert.Error(t, e)
	})
}

func TestUpdateCake(t *testing.T) {
	var ctx context.Context
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockCakeRepository(mockCtrl)
	cs := cakeService{CakeRepository: mockRepo}

	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().UpdateCake(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		e := cs.UpdateCake(ctx, 1, entity.Cake{})
		assert.Nil(t, e)
	})

	t.Run("Repo Error", func(t *testing.T) {
		mockRepo.EXPECT().UpdateCake(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("ERRPR"))
		e := cs.UpdateCake(ctx, 1, entity.Cake{})
		assert.Error(t, e)
	})
}

func TestAddNewCake(t *testing.T) {
	var ctx context.Context
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockCakeRepository(mockCtrl)
	cs := cakeService{CakeRepository: mockRepo}

	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().InsertCake(gomock.Any(), gomock.Any()).Return(nil)
		e := cs.AddNewCake(ctx, entity.Cake{})
		assert.Nil(t, e)
	})

	t.Run("Repo error", func(t *testing.T) {
		mockRepo.EXPECT().InsertCake(gomock.Any(), gomock.Any()).Return(errors.New("ERROR"))
		e := cs.AddNewCake(ctx, entity.Cake{})
		assert.Error(t, e)
	})
}

func TestDeleteCake(t *testing.T) {
	var ctx context.Context
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepo := mocks.NewMockCakeRepository(mockCtrl)
	cs := cakeService{CakeRepository: mockRepo}

	t.Run("Success", func(t *testing.T) {
		mockRepo.EXPECT().DeleteCake(gomock.Any(), gomock.Any()).Return(nil)
		e := cs.DeleteCake(ctx, 1)
		assert.Nil(t, e)
	})

	t.Run("Repo Error", func(t *testing.T) {
		mockRepo.EXPECT().DeleteCake(gomock.Any(), gomock.Any()).Return(errors.New("ERROR"))
		e := cs.DeleteCake(ctx, 1)
		assert.Error(t, e)
	})
}
