package main

import (
	"find-me-a-car/providers/mercedes_benz"
	"find-me-a-car/services"
)

func main() {
	collectorService := services.NewCollectorService()
	collectorService.AddProvider(&mercedes_benz.Provider{
		MinPrice: 0,
		MaxPrice: 500000,
		Classes: []mercedes_benz.Class{mercedes_benz.CClass, mercedes_benz.EClass},
		FuelTypes: []mercedes_benz.FuelType{mercedes_benz.Bensin},
	})
	collectedCars := collectorService.Collect()

	dataService := services.NewDataService()
	dataService.Write(collectedCars)

	read := dataService.Read()
	print(read)
}
