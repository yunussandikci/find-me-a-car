package main

import (
	"find-me-a-car/sahibinden"
	"strings"
	"time"
)

var Brand = "Mercedes"
var Model = "C 180"

var Sellers = []string{
	"otomol",
}

func main() {
	dataService := NewDataService()
	for {
		client := sahibinden.NewClient(5*time.Second, 0*time.Second)
		for _, seller := range Sellers {
			for _, car := range client.GetSellerCars(seller) {
				if strings.Contains(car.Brand, Brand) && strings.Contains(car.Model, Model) {
					if !dataService.IfExist(car.Url) {
						dataService.Append(car.Url)
					}
				}
			}
		}
	}
}