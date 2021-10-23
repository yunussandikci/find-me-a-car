package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const dataFile = "./cars.txt"

type DataService struct {
}

func NewDataService() *DataService {
	return &DataService{}
}

func (d*DataService) Append(text string) {
	f, err := os.OpenFile(dataFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(fmt.Sprintf("%s\n",text)); err != nil {
		panic(err)
	}
}

func (d*DataService) IfExist(text string) bool {
	f, err := os.OpenFile(dataFile, os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	all, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.Contains(string(all), text)
}

func init() {
	file, err := os.OpenFile(dataFile, os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	file.Close()
}