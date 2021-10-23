package main

import (
	"math/bits"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetTelegramToken() string {
	token, found := os.LookupEnv("TELEGRAM_TOKEN")
	if !found {
		panic("Telegram Token Not Found.")
	}
	return token
}

func GetTelegramChatId() int64 {
	chatId, found := os.LookupEnv("TELEGRAM_CHAT_ID")
	if !found {
		panic("Telegram Chat ID Found.")
	}
	chatIdInt, err := strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		panic(err)
	}
	return chatIdInt
}

func GetDelayRandom() time.Duration {
	delayRandom, found := os.LookupEnv("DELAY_RANDOM")
	if !found {
		return DefaultDelayRandom
	}
	intVar, err := strconv.Atoi(delayRandom)
	if err != nil {
		panic(err)
	}
	return time.Duration(rand.Int31n(int32(intVar))) * time.Second
}

func GetDelayConstant() time.Duration {
	delayConstant, found := os.LookupEnv("DELAY_CONSTANT")
	if !found {
		return DefaultDelayConstant
	}
	intVar, err := strconv.Atoi(delayConstant)
	if err != nil {
		panic(err)
	}
	return time.Duration(rand.Int31n(int32(intVar))) * time.Second
}

func GetBrand() string {
	return os.Getenv("BRAND")
}

func GetModel() string {
	return os.Getenv("MODEL")
}

func GetMinimumYear() int {
	minimumYear, found := os.LookupEnv("MINIMUM_YEAR")
	if !found {
		return 0
	}
	intVar, err := strconv.Atoi(minimumYear)
	if err != nil {
		panic(err)
	}
	return intVar
}

func GetMaximumMileage() int {
	minimumMileage, found := os.LookupEnv("MAXIMUM_MILEAGE")
	if !found {
		return (1 << bits.UintSize) / 2 - 1
	}
	intVar, err := strconv.Atoi(minimumMileage)
	if err != nil {
		panic(err)
	}
	return intVar
}

func GetSellerDomains() []string {
	sellerDomains, found := os.LookupEnv("SELLER_DOMAINS")
	if !found {
		panic("Seller domains not found.")
	}
	return strings.Split(strings.ReplaceAll(sellerDomains," ",""),",")
}

