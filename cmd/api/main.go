//	Car info.
//
//	Documentation of cars API.
//
//	Schemes:
//	  - http
//	  - https
//	BasePath: /
//	Version: 0.0.1
//
//	Consumes:
//	  - application/json
//
//	Produces:
//	  - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ichetiva/effective-mobile-test-task/config"
	ccontroller "github.com/ichetiva/effective-mobile-test-task/module/car/controller"
	"github.com/ichetiva/effective-mobile-test-task/module/car/models"
	cservice "github.com/ichetiva/effective-mobile-test-task/module/car/service"
	cstorage "github.com/ichetiva/effective-mobile-test-task/module/car/storage"
	"github.com/ichetiva/effective-mobile-test-task/pkg/responder"
	"github.com/ichetiva/effective-mobile-test-task/router"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// get config
	conf := config.New()

	// logger
	log := logrus.New()

	// connect to database
	log.Debugln("Connecting to database..")
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		conf.Database.User, conf.Database.Password, conf.Database.Host, conf.Database.Port, conf.Database.Name,
	)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln("failed to create database connection:", err.Error())
	}

	// migrate
	db.AutoMigrate(&models.Car{}, &models.Owner{})

	// create responder
	respond := responder.New(log)

	// create storages
	carStorage := cstorage.New(db)

	// create services
	carService := cservice.New(carStorage)

	// create controllers
	carController := ccontroller.New(log, carService, respond)

	r := router.New(log, carController)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	log.Infoln("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil && err != http.ErrServerClosed {
		log.Errorln("Server error:", err.Error())
	}
}
