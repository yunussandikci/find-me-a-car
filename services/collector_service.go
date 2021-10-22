package services

import (
	"find-me-a-car/models"
)

type CollectorService struct {
	providers []models.Provider
}

func NewCollectorService() *CollectorService {
	return &CollectorService{}
}

func (c *CollectorService) AddProvider(provider models.Provider) {
	c.providers = append(c.providers, provider)
}

func (c *CollectorService) Collect() []*models.Car {
	var cars []*models.Car
	for _,provider := range c.providers {
		cars = append(cars, provider.Provide()...)
	}
	return cars
}