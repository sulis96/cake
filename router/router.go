package router

import (
	"CAKE-STORE/controller"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	cakeRouter struct {
		CakeController controller.CakeController
	}

	CakeRouter interface {
		Route()
	}
)

func NewCakeRouter(cakeController *controller.CakeController) CakeRouter {
	return &cakeRouter{
		CakeController: *cakeController,
	}
}

func (cr *cakeRouter) Route() {
	router := httprouter.New()
	router.GET("/cakes", cr.CakeController.ListCake)
	router.POST("/cakes", cr.CakeController.AddNewCake)
	router.GET("/cakes/:{id}", cr.CakeController.DetailCake)
	router.PATCH("/cakes/:{id}", cr.CakeController.UpdateCake)
	router.DELETE("/cakes/:{id}", cr.CakeController.DeleteCake)

	log.Fatal(http.ListenAndServe(":8080", router))
}
