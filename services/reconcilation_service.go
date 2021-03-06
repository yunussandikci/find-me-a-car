package services

import (
	"find-me-a-car/common"
	"find-me-a-car/models"
	"find-me-a-car/services/sahibinden"
)

type ReconciliationService struct {
	telegramService  *TelegramService
	dataService      *DataService
	sahibindenClient *sahibinden.Client
	SellerDomains    []string
	Brand            string
	Model            string
	MinimumYear      int
	MaximumMileage   int
	MaximumPrice     int
}

func NewReconciliationService(telegramService *TelegramService, dataService *DataService,
	sahibindenClient *sahibinden.Client) *ReconciliationService {
	return &ReconciliationService{
		telegramService:  telegramService,
		dataService:      dataService,
		sahibindenClient: sahibindenClient,
	}
}

func (s *ReconciliationService) Run() {
	for {
		s.reconcileSellers()
	}
}

func (s *ReconciliationService) reconcileSellers() {
	for _, seller := range s.SellerDomains {
		s.reconcileCars(s.sahibindenClient.GetSellerCars(seller))
	}
}

func (s *ReconciliationService) reconcileCars(cars []models.Car) {
	for _, car := range cars {
		s.reconcileCar(car)
	}
}

func (s *ReconciliationService) reconcileCar(car models.Car) {
	if car.IsMatches(s.Brand, s.Model, s.MinimumYear, s.MaximumMileage, s.MaximumPrice) {
		if !s.dataService.IfExist(car.Url) {
			common.Logger.Infof("Found new car: %s", car.Url)
			s.dataService.Append(car.Url)
			s.telegramService.SendMessage(car.GetMarkdown())
		}
	}
}
