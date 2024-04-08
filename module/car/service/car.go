package service

import (
	"context"

	"github.com/ichetiva/effective-mobile-test-task/module/car/models"
	"github.com/ichetiva/effective-mobile-test-task/module/car/storage"
)

type CarServicer interface {
	Create(ctx context.Context, car *models.Car) error
	Update(ctx context.Context, car *models.Car) error
	Delete(ctx context.Context, carID uint) error
	Get(ctx context.Context, id uint, regNum, mark, model string) ([]models.Car, error)
}

type CarService struct {
	storage storage.CarStorager
}

func New(storage storage.CarStorager) CarServicer {
	return &CarService{
		storage: storage,
	}
}

func (c *CarService) Create(ctx context.Context, car *models.Car) error {
	return c.storage.Create(ctx, car)
}

func (c *CarService) Update(ctx context.Context, car *models.Car) error {
	return c.storage.Update(ctx, car)
}

func (c *CarService) Delete(ctx context.Context, carID uint) error {
	return c.storage.Delete(ctx, carID)
}

func (c *CarService) Get(ctx context.Context, id uint, regNum string, mark string, model string) ([]models.Car, error) {
	return c.storage.Get(ctx, &storage.CarOptions{
		ID:     id,
		RegNum: regNum,
		Mark:   mark,
		Model:  model,
	})
}
