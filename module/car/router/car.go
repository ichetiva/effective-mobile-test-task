package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ichetiva/effective-mobile-test-task/module/car/controller"
)

func RouteCar(cc *controller.CarController) http.Handler {
	r := chi.NewRouter()

	r.Post("/", cc.Create)
	r.Put("/{carId}/", cc.Update)
	r.Delete("/{carId}/", cc.Delete)
	r.Get("/", cc.Get)

	return r
}
