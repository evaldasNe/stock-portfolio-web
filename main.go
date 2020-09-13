package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/evaldasNe/stock-portfolio-web/Config"
	"github.com/evaldasNe/stock-portfolio-web/Models"
	"github.com/evaldasNe/stock-portfolio-web/Routes"
	"github.com/piquette/finance-go/quote"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Stock{})
	Config.DB.AutoMigrate(&Models.OwnedStock{})

	//addStocksToDB()

	r := Routes.SetupRouter()
	//running
	r.Run(":" + os.Getenv("APP_PORT"))
}

func addStocksToDB() {
	allStocks := getAllStocks()

	for _, stock := range allStocks {
		q, err := quote.Get(stock["symbol"].(string))
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else if q == nil {
			continue
		}

		newStock := Models.Stock{
			Ticker:       stock["symbol"].(string),
			CompanyName:  q.ShortName,
			SellingPrice: q.Bid}
		err = Models.CreateStock(&newStock)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getAllStocks() (results []map[string]interface{}) {
	client := http.Client{}
	request, err := http.NewRequest("GET", "https://finnhub.io/api/v1/stock/symbol?exchange=US&token="+os.Getenv("FINNHUB_API_TOKEN"), nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	json.NewDecoder(resp.Body).Decode(&results)

	return results
}
