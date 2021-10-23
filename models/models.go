package models

import (
	"fmt"
	"strings"
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

func (c *Car) IsBrand(brand string) bool {
	return strings.Contains(c.Brand, brand)
}

func (c *Car) IsModel(model string) bool {
	return strings.Contains(c.Model, model)
}

func (c *Car) ExceedMaximumMileage(maximumMileage int) bool {
	return c.Mileage <= maximumMileage
}

func (c *Car) LowerThanMinimumYear(minimumYear int) bool {
	return c.Year >= minimumYear
}

func (c *Car) GetMarkdown() string {
	return fmt.Sprintf("*Model*: %s\n*Marka*: %s\n*Yıl:* %d\n*KM:* %d\n*Renk:* %s\n*Fiyat:* %d TL\n*Link:* %s\n\n[Görsel:](%s)", c.Model, c.Brand , c.Year, c.Mileage, c.Color, c.Price, c.Url, c.Image)
}
