package sahibinden

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"time"
)

type Client struct {
	constantDelay time.Duration
	collector     *colly.Collector
}

const UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36"

func NewClient(randomDelay, constantDelay time.Duration) *Client {
	collector := colly.NewCollector(colly.UserAgent(UserAgent))
	limitErr := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*sahibinden.com*",
		RandomDelay: randomDelay,
	})
	if limitErr != nil {
		panic(limitErr)
	}
	return &Client{
		collector:     collector,
		constantDelay: constantDelay,
	}
}

func (c *Client) GetSellerCars(seller string) []Car {
	time.Sleep(c.constantDelay)
	collector := c.collector.Clone()
	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL.String())
	})
	var cars []Car
	collector.OnHTML(".store-page .classified-list tbody tr", func(e *colly.HTMLElement) {
		var car Car
		e.ForEach("td", func(i int, e *colly.HTMLElement) {
			switch i {
			case 0:
				car.Url = e.ChildAttr("a", "href")
				car.Image = e.ChildAttr("img", "src")
				break
			case 1:
				car.Brand = ClearText(e.Text)
				break
			case 2:
				car.Model = ClearText(e.Text)
				break
			case 3:
				car.Year = ParseInteger(ClearText(e.Text))
				break
			case 4:
				car.Mileage = ParseInteger(ClearText(e.Text))
				break
			case 5:
				car.Color = ClearText(e.Text)
				break
			case 6:
				car.Price = ParseInteger(ClearText(e.Text))
				break
			}
		})
		cars = append(cars, car)
	})
	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.Contains(e.Text, "Sonraki") {
			_ = e.Request.Visit(link)
		}
	})
	visitErr := collector.Visit(GetSellerPageUrl(seller, 50))
	if visitErr != nil {
		fmt.Println(visitErr)
	}
	return cars
}
