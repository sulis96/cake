package main

import (
	"CAKE-STORE/config"
	"CAKE-STORE/controller"
	"CAKE-STORE/repository"
	"CAKE-STORE/router"
	"CAKE-STORE/service"
	"log"
	"net/http"
)

func main() {
	// setup config
	database, err := config.MySqlDatabase()
	if err != nil {
		log.Panic(err)
	}
	defer database.Close()

	// Health Check
	err = database.Ping()
	if err != nil {
		log.Panic(err)
		return
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Status OK"))
	})

	// setup repository
	cakeRepository := repository.NewCakeRepository(database)

	// setup service
	cakeService := service.NewCakeService(&cakeRepository)

	// setup controller
	cakeController := controller.NewCakeController(&cakeService)

	// setup router
	CakeRouter := router.NewCakeRouter(&cakeController)
	CakeRouter.Route()
}
