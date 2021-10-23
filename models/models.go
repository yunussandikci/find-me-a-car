package models

import (
	"find-me-a-car/common"
	"fmt"
	"regexp"
)

type Car struct {
	Url     string
	Image   string
	Brand   string
	Model   string
	Year    int
	Mileage int
	Color   string
	Price   int
}

func (c *Car) IsMatches(brand, model string, minimumYear, maximumMileage, maximumPrice int) bool {
	brandMatched, regexErr := regexp.MatchString(brand, c.Brand)
	if regexErr != nil {
		common.Logger.Errorf(regexErr.Error())
	}
	modelMatched, regexErr := regexp.MatchString(model, c.Model)
	if regexErr != nil {
		common.Logger.Errorf(regexErr.Error())
	}
	mileageMatch := c.Mileage <= maximumMileage
	priceMatch := c.Price <= maximumPrice
	yearMatch := c.Year >= minimumYear
	return brandMatched && modelMatched && mileageMatch && priceMatch && yearMatch
}

func (c *Car) GetMarkdown() string {
	return fmt.Sprintf("*Model*: %s\n*Marka*: %s\n*Yıl:* %d\n*KM:* %d\n*Renk:* %s\n*Fiyat:* %d TL\n*Link:* %s\n\n[Görsel:](%s)", c.Model, c.Brand , c.Year, c.Mileage, c.Color, c.Price, c.Url, c.Image)
}
