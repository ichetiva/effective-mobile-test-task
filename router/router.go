package router

import (
	logger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	ccontroller "github.com/ichetiva/effective-mobile-test-task/module/car/controller"
	crouter "github.com/ichetiva/effective-mobile-test-task/module/car/router"
	"github.com/sirupsen/logrus"
)

func New(log *logrus.Logger, cc *ccontroller.CarController) *chi.Mux {
	r := chi.NewRouter()

	r.Use(logger.Logger("router", log))
	r.Use(middleware.Recoverer)

	r.Mount("/api/car", crouter.RouteCar(cc))

	RouteSwagger(log, r)

	return r
}
