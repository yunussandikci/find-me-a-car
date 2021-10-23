package main

import (
	"find-me-a-car/common"
	"find-me-a-car/services"
	"find-me-a-car/services/sahibinden"
	"time"
)

var DefaultDelayRandom = 10 * time.Second
var DefaultDelayConstant = 50 * time.Second

func main() {
	telegramToken := GetTelegramToken()
	telegramChatID := GetTelegramChatId()
	delayRandom := GetDelayRandom()
	delayConstant := GetDelayConstant()
	brand := GetBrand()
	model := GetModel()
	maximumMileage := GetMaximumMileage()
	domains := GetSellerDomains()

	common.Logger.Infof("Telegram Token: %s\n", telegramToken)
	common.Logger.Infof("Telegram ChatId: %d\n", telegramChatID)
	common.Logger.Infof("Delay Random: %s\n", delayRandom)
	common.Logger.Infof("Delay Constant: %s\n", delayConstant)
	common.Logger.Infof("Brand: %s\n", brand)
	common.Logger.Infof("Model: %s\n", model)
	common.Logger.Infof("Maximum Mileage: %d\n", maximumMileage)
	common.Logger.Infof("Domains: %s\n", domains)

	return
	dataService := services.NewDataService()
	telegramService := services.NewTelegramService(telegramToken, telegramChatID)
	sahibindenClient := sahibinden.NewClient(delayRandom, delayConstant)
	reconciliationService := services.NewReconciliationService(telegramService, dataService, sahibindenClient)
	reconciliationService.SellerDomains = domains
	reconciliationService.Brand = brand
	reconciliationService.Model = model
	reconciliationService.MaximumMileage = maximumMileage
	reconciliationService.Run()
}

