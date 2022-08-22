package controller

import (
	"CAKE-STORE/entity"
	"CAKE-STORE/service"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type (
	cakeController struct {
		CakeService service.CakeService
	}

	CakeController interface {
		ListCake(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		DetailCake(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		AddNewCake(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		UpdateCake(w http.ResponseWriter, r *http.Request, p httprouter.Params)
		DeleteCake(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	}
)

func NewCakeController(cakeService *service.CakeService) CakeController {
	return &cakeController{
		CakeService: *cakeService,
	}
}

func (cc *cakeController) ListCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	titleParam := r.URL.Query().Get("title")
	cakes, err := cc.CakeService.ListCake(ctx, titleParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cakesByte, err := json.Marshal(cakes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(cakesByte))
}

func (cc *cakeController) DetailCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idParamURL := strings.TrimPrefix(r.URL.Path, "/cakes/:")
	idParam, err := strconv.Atoi(idParamURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cakes, err := cc.CakeService.DetailCake(ctx, idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cakesByte, err := json.Marshal(cakes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(cakesByte))
}

func (cc *cakeController) AddNewCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cake entity.Cake
	err := json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = cc.CakeService.AddNewCake(ctx, cake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success Insert DB"))
}

func (cc *cakeController) UpdateCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idParamURL := strings.TrimPrefix(r.URL.Path, "/cakes/:")
	idParam, err := strconv.Atoi(idParamURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cake entity.Cake
	err = json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = cc.CakeService.UpdateCake(ctx, idParam, cake)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Success Update cake"))
}

func (cc *cakeController) DeleteCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	idParamURL := strings.TrimPrefix(r.URL.Path, "/cakes/:")
	idParam, err := strconv.Atoi(idParamURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = cc.CakeService.DeleteCake(ctx, idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("Success Delete cake"))
}
