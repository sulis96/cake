package controller

import (
	"CAKE-STORE/entity"
	"CAKE-STORE/mocks"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestListCake(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockCakeService(mockCtrl)
	cc := cakeController{CakeService: mockService}

	req, err := http.NewRequest("GET", "localhost:8080/cakes", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	p := httprouter.Params{}

	b, _ := json.Marshal([]entity.ListCake{})
	buff := bytes.NewBuffer(b)

	mockService.EXPECT().ListCake(gomock.Any(), gomock.Any()).Return([]entity.ListCake{}, nil)
	cc.ListCake(rec, req, p)
	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, buff, rec.Body)
}

func TestDetailCake(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockCakeService(mockCtrl)
	cc := cakeController{CakeService: mockService}

	req, err := http.NewRequest("GET", "/cakes/:2", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	p := httprouter.Params{}

	cake := entity.Cake{Title: "abc", Description: "desc", Rating: 2}
	b, _ := json.Marshal(cake)
	buff := bytes.NewBuffer(b)

	mockService.EXPECT().DetailCake(gomock.Any(), gomock.Any()).Return(cake, nil)

	cc.DetailCake(rec, req, p)
	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, buff, rec.Body)
}

func TestAddNewCake(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockCakeService(mockCtrl)
	cc := cakeController{CakeService: mockService}

	reqBody := `{"title":"cake","rating":9, "description":"cake"}`

	req, err := http.NewRequest("POST", "/cakes", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	p := httprouter.Params{}

	mockService.EXPECT().AddNewCake(gomock.Any(), gomock.Any()).Return(nil)
	cc.AddNewCake(rec, req, p)
	assert.Nil(t, err)
	assert.Equal(t, 201, rec.Code)
	assert.Equal(t, bytes.NewBuffer([]byte("Success Insert DB")), rec.Body)
}

func TestUpdateCake(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockCakeService(mockCtrl)
	cc := cakeController{CakeService: mockService}

	reqBody := `{"title":"cake","rating":9, "description":"cake"}`

	req, err := http.NewRequest("PATCH", "/cakes/:2", bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	p := httprouter.Params{}

	mockService.EXPECT().UpdateCake(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	cc.UpdateCake(rec, req, p)
	assert.Nil(t, err)
	assert.Equal(t, 201, rec.Code)
	assert.Equal(t, bytes.NewBuffer([]byte("Success Update cake")), rec.Body)
}

func TestDeleteCake(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockService := mocks.NewMockCakeService(mockCtrl)
	cc := cakeController{CakeService: mockService}

	req, err := http.NewRequest("DELETE", "/cakes/:2", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	rec := httptest.NewRecorder()
	p := httprouter.Params{}

	mockService.EXPECT().DeleteCake(gomock.Any(), gomock.Any()).Return(nil)
	cc.DeleteCake(rec, req, p)
	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, bytes.NewBuffer([]byte("Success Delete cake")), rec.Body)
}
