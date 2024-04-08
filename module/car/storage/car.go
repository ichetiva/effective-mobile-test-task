package storage

import (
	"context"

	"github.com/ichetiva/effective-mobile-test-task/module/car/models"
	"gorm.io/gorm"
)

type CarOptions struct {
	ID     uint
	RegNum string
	Mark   string
	Model  string
}

type CarStorager interface {
	Create(ctx context.Context, car *models.Car) error
	Update(ctx context.Context, car *models.Car) error
	Delete(ctx context.Context, carID uint) error
	Get(ctx context.Context, options *CarOptions) ([]models.Car, error)
}

type CarStorage struct {
	storage *gorm.DB
}

func New(storage *gorm.DB) CarStorager {
	return &CarStorage{
		storage: storage,
	}
}

func (c *CarStorage) Create(ctx context.Context, car *models.Car) error {
	result := c.storage.WithContext(ctx).Create(car)
	return result.Error
}

func (c *CarStorage) Update(ctx context.Context, car *models.Car) error {
	result := c.storage.WithContext(ctx).Save(car)
	return result.Error
}

func (c *CarStorage) Delete(ctx context.Context, carID uint) error {
	result := c.storage.WithContext(ctx).Delete(&models.Car{}, "id = ?", carID)
	return result.Error
}

func (c *CarStorage) Get(ctx context.Context, options *CarOptions) ([]models.Car, error) {
	conditions := make(map[string]interface{})
	if options.ID > 0 {
		conditions["id"] = options.ID
	}
	if options.RegNum != "" {
		conditions["reg_num"] = options.RegNum
	}
	if options.Mark != "" {
		conditions["mark"] = options.Mark
	}
	if options.Model != "" {
		conditions["model"] = options.Model
	}
	var cars []models.Car
	result := c.storage.WithContext(ctx).Where(conditions).Preload("Owner").Find(&cars)
	return cars, result.Error
}
