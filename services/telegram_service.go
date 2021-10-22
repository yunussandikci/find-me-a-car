package services

import "find-me-a-car/models"

type TelegramService struct {
	providers []models.Provider
}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}
