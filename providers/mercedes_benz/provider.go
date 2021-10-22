package mercedes_benz

import (
	"encoding/json"
	"find-me-a-car/models"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const BaseApiUrl = "https://dcp-otos-prod.api.mercedes-benz.com/dcp-api/v2/dcp-mp-tr/products/search"

func (p *Provider) generateQuery() string {
	query := ":relevance:useProductType:UCOS:allCategories:dcp-mp-tr-vehicles"
	for _, v := range p.FuelTypes {
		query += fmt.Sprintf(":fuel_type:%s", v)
	}
	for _, v := range p.Classes {
		query += fmt.Sprintf(":model:%s", v)
	}
	query = fmt.Sprintf("%s:priceSlider:[%d TO %d]", query, p.MinPrice, p.MaxPrice)
	return url.QueryEscape(query)
}

func (p *Provider) generateApiUrl(currentPage, pageSize string) string {
	return fmt.Sprintf("%s?fields=FULL&currentPage=%s&pageSize=%s&query=%s",
		BaseApiUrl, currentPage, pageSize, p.generateQuery())
}

func (p *Provider) Provide() []*models.Car {
	resp, err := http.Get(p.generateApiUrl("0", "24"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response Response
	_ = json.Unmarshal(body, &response)
	var cars []*models.Car
	for _,v := range response.Products {
		cars = append(cars, &models.Car{
			Id: v.Code,
			Brand: "Mercedes",
			Model: v.Name,
			Year:  v.ModelYear,
			Mileage: fmt.Sprintf("%.0f %s",v.Mileage.Value, v.Mileage.Unit),
			Price: v.Price.FormattedValue,
		})
	}
	return cars
}
