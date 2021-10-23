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
	minimumYear := GetMinimumYear()
	maximumMileage := GetMaximumMileage()
	maximumPrice := GetMaximumPrice()
	domains := GetSellerDomains()

	common.Logger.Infof("Telegram Token: %s", telegramToken)
	common.Logger.Infof("Telegram ChatId: %d", telegramChatID)
	common.Logger.Infof("Delay Random: %s", delayRandom)
	common.Logger.Infof("Delay Constant: %s", delayConstant)
	common.Logger.Infof("Brand: %s", brand)
	common.Logger.Infof("Model: %s", model)
	common.Logger.Infof("Minimum Year: %d", minimumYear)
	common.Logger.Infof("Maximum Mileage: %d", maximumMileage)
	common.Logger.Infof("Maximum Price: %d", maximumPrice)
	common.Logger.Infof("Domains: %s", domains)

	dataService := services.NewDataService()
	telegramService := services.NewTelegramService(telegramToken, telegramChatID)
	sahibindenClient := sahibinden.NewClient(delayRandom, delayConstant)
	reconciliationService := services.NewReconciliationService(telegramService, dataService, sahibindenClient)
	reconciliationService.SellerDomains = domains
	reconciliationService.Brand = brand
	reconciliationService.Model = model
	reconciliationService.MaximumMileage = maximumMileage
	reconciliationService.MaximumPrice = maximumPrice
	reconciliationService.MinimumYear = minimumYear
	reconciliationService.Run()
}

