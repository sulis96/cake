package main

import (
	"CAKE-STORE/config"
	"CAKE-STORE/controller"
	"CAKE-STORE/repository"
	"CAKE-STORE/router"
	"CAKE-STORE/service"
	"log"
)

func main() {
	// setup config
	database, err := config.MySqlDatabase()
	if err != nil {
		log.Panic(err)
	}
	defer database.Close()

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
