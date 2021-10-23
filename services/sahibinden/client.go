package sahibinden

import (
	"find-me-a-car/common"
	"find-me-a-car/models"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

type Client struct {
	randomDelay   time.Duration
	constantDelay time.Duration
}

func NewClient(randomDelay, constantDelay time.Duration) *Client {
	return &Client{
		randomDelay: randomDelay,
		constantDelay: constantDelay,
	}
}

func (c *Client) GetSellerCars(seller string) []models.Car {
	collector := c.newCollector()
	var cars []models.Car
	collector.OnHTML(".store-page .classified-list tbody tr", func(e *colly.HTMLElement) {
		var car models.Car
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
	visitErr := collector.Visit(fmt.Sprintf(SellerUrl, seller, Domain))
	if visitErr != nil {
		common.Logger.Errorln(visitErr)
	}
	return cars
}

func (c *Client) newCollector() *colly.Collector {
	collector := colly.NewCollector()
	collector.OnRequest(func(r *colly.Request) {
		time.Sleep(c.constantDelay)
		common.Logger.Infof("Visiting: %s", r.URL.String())
		r.Headers.Set("cookie", Cookies)
		r.Headers.Set("user-agent", UserAgent)
	})
	_ = collector.Limit(&colly.LimitRule{
		DomainGlob:  fmt.Sprintf("*%s*", Domain),
		RandomDelay: c.randomDelay,
	})
	return collector
}
