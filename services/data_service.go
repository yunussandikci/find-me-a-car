package services

import (
	"encoding/csv"
	"find-me-a-car/models"
	"os"
)

const dataFile = "./cars.csv"

type DataService struct {
}

func NewDataService() *DataService {
	return &DataService{}
}

func (d*DataService) Read() []*models.Car {
	var cars []*models.Car
	data := d.readCSV()
	for _, line := range data {
		cars = append(cars, &models.Car{
			Id:      line[0],
			Brand:   line[1],
			Model:   line[2],
			Year:    line[3],
			Mileage: line[4],
			Price:   line[5],
		})
	}
	return cars
}

func (d*DataService) Write(cars []*models.Car) {
	var data [][]string
	for _, car := range cars {
		data = append(data, []string{car.Id, car.Brand, car.Model, car.Year, car.Mileage, car.Price})
	}
	d.writeCSV(data)
}

func (d*DataService) readCSV() [][]string {
	file, err := os.Open(dataFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	if _, err := reader.Read(); err != nil {
		panic(err)
	}
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	return records
}

func (d*DataService) writeCSV(data [][]string) {
	file, err := os.OpenFile(dataFile, os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		panic(err)
	}
}

func init() {
	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	file.Close()
}
