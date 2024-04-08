package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ichetiva/effective-mobile-test-task/module/car/models"
	"github.com/ichetiva/effective-mobile-test-task/module/car/service"
	"github.com/ichetiva/effective-mobile-test-task/pkg/responder"
	"github.com/sirupsen/logrus"
)

type CarController struct {
	log     *logrus.Logger
	service service.CarServicer
	responder.Responder
}

func New(log *logrus.Logger, service service.CarServicer, responder responder.Responder) *CarController {
	return &CarController{
		log:       log,
		service:   service,
		Responder: responder,
	}
}

// swagger:route post /api/car/ car createCar
// creates car
// responses:
//
//	200: CarResponse
//	400: ErrorResponse
//	500: ErrorResponse
func (c *CarController) Create(w http.ResponseWriter, r *http.Request) {
	var car CarRequest
	var err error

	err = json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		c.log.Errorln("failed to decode request data:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbCar := &models.Car{
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Owner: models.Owner{
			Name:    car.Owner.Name,
			Surname: car.Owner.Surname,
		},
	}

	ctx := context.Background()
	err = c.service.Create(ctx, dbCar)
	if err != nil {
		c.log.Errorln("failed to create car in database:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.OutputJSON(w, CarResponse{
		ID:     dbCar.ID,
		RegNum: dbCar.RegNum,
		Mark:   dbCar.Mark,
		Model:  dbCar.Model,
		Owner: OwnerResponse{
			Name:    dbCar.Owner.Name,
			Surname: dbCar.Owner.Surname,
		},
	})
}

// swagger:route put /api/car/{carId}/ car car
// updates car
// responses:
//
//	200: CarResponse
//	400: ErrorResponse
//	500: ErrorResponse
func (c *CarController) Update(w http.ResponseWriter, r *http.Request) {
	var err error

	carIDstr := chi.URLParam(r, "carId")
	if carIDstr == "" {
		c.ErrorJSON(w, "car id isn't set", http.StatusBadRequest)
		return
	}

	carID, err := strconv.Atoi(carIDstr)
	if err != nil {
		c.log.Errorln("failed to convert car id to int:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var car CarRequest

	err = json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		c.log.Errorln("failed to decode request data:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbCar := &models.Car{
		ID:     uint(carID),
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Owner: models.Owner{
			Name:    car.Owner.Name,
			Surname: car.Owner.Surname,
		},
	}

	ctx := context.Background()
	err = c.service.Update(ctx, dbCar)
	if err != nil {
		c.log.Errorln("failed to update car in database:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.OutputJSON(w, CarResponse{
		ID:     dbCar.ID,
		RegNum: dbCar.RegNum,
		Mark:   dbCar.Mark,
		Model:  dbCar.Model,
		Owner: OwnerResponse{
			Name:    dbCar.Owner.Name,
			Surname: dbCar.Owner.Surname,
		},
	})
}

// swagger:route delete /api/car/{carId}/ car carId
// deletes car
// responses:
//
//	200: EmptyResponse
//	400: ErrorResponse
//	500: ErrorResponse
func (c *CarController) Delete(w http.ResponseWriter, r *http.Request) {
	var err error

	carIDstr := chi.URLParam(r, "carId")
	if carIDstr == "" {
		c.ErrorJSON(w, "car id isn't set", http.StatusBadRequest)
		return
	}

	carID, err := strconv.Atoi(carIDstr)
	if err != nil {
		c.log.Errorln("failed to convert car id to int:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	err = c.service.Delete(ctx, uint(carID))
	if err != nil {
		c.log.Errorln("failed to delete car from databse:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.OutputJSON(w, EmptyResponse{})
}

// swagger:route get /api/car/ car carFilters
// gets car by filters
// responses:
//
//	200: CarsResponse
//	500: ErrorResponse
func (c *CarController) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	regNum := r.URL.Query().Get("regNum")
	mark := r.URL.Query().Get("mark")
	model := r.URL.Query().Get("model")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		id = 0
	}

	ctx := context.Background()
	cars, err := c.service.Get(ctx, uint(id), regNum, mark, model)
	if err != nil {
		c.log.Errorln("failed to get cars from database:", err.Error())
		c.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	carResponse := make(CarsResponse, 0, len(cars))
	for _, car := range cars {
		carResponse = append(carResponse, CarResponse{
			ID:     car.ID,
			RegNum: car.RegNum,
			Mark:   car.Mark,
			Model:  car.Model,
			Owner: OwnerResponse{
				Name:    car.Owner.Name,
				Surname: car.Owner.Surname,
			},
		})
	}
	c.OutputJSON(w, carResponse)
}
