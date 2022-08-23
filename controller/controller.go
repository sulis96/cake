package controller

import (
	"CAKE-STORE/config"
	"CAKE-STORE/entity"
	"CAKE-STORE/service"
	"CAKE-STORE/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
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
		HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	}
)

func NewCakeController(cakeService *service.CakeService) CakeController {
	return &cakeController{
		CakeService: *cakeService,
	}
}

func (cc *cakeController) HealthCheck(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	database, err := config.MySqlDatabase()
	if err != nil {
		log.Panic(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		log.Panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status OK"))
}

func (cc *cakeController) ListCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	titleParam := r.URL.Query().Get("title")
	cakes, err := cc.CakeService.ListCake(ctx, titleParam)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	cakesByte, err := json.Marshal(cakes)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cakesByte)
}

func (cc *cakeController) DetailCake(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	idParamURL := strings.TrimPrefix(r.URL.Path, "/cakes/:")
	idParam, err := strconv.Atoi(idParamURL)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	cakes, err := cc.CakeService.DetailCake(ctx, idParam)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	cakesByte, err := json.Marshal(cakes)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
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
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(cake)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err = cc.CakeService.AddNewCake(ctx, cake)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
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
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	var cake entity.Cake
	err = json.NewDecoder(r.Body).Decode(&cake)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	validate := validator.New()
	err = validate.Struct(cake)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	err = cc.CakeService.UpdateCake(ctx, idParam, cake)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
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
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusBadRequest, "Bad request", err.Error())
		return
	}

	err = cc.CakeService.DeleteCake(ctx, idParam)
	if err != nil {
		log.Println(err)
		utils.HandleErrorResponse(w, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Delete cake"))
}
