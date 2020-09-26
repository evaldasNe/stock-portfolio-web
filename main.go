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

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error

func main() {
	dsn := Config.DbURL(Config.BuildDBConfig())
	Config.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	Config.DB.AutoMigrate(&Models.User{})
	Config.DB.AutoMigrate(&Models.Stock{})
	Config.DB.AutoMigrate(&Models.OwnedStock{})
	Config.DB.AutoMigrate(&Models.Comment{})
	Config.DB.AutoMigrate(&Models.Token{})

	//addStocksToDB()

	r := Routes.SetupRouter()

	//go refreshTokens(5 * time.Second)

	//running
	r.Run(":" + os.Getenv("APP_PORT"))
}

/*
func refreshTokens(period time.Duration) {
	for {
		time.Sleep(period)
		go func() {
			var tokens []Models.Token
			if err := Models.GetAllTokens(&tokens); err != nil {
				fmt.Printf("[%s] Error getting all tokens: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
				return
			}

			for _, token := range tokens {
				source := Services.GetTokenSource(token.Token)
				newToken, err := source.Token()
				if err != nil {
					fmt.Printf("[%s] Error getting new token: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
					return
				}
				token.Token = newToken
				if err = Models.UpdateToken(&token); err != nil {
					fmt.Printf("[%s] Error updating token: %s\n", time.Now().Format("2006-01-02 15:04:05"), err.Error())
					return
				}
			}
			fmt.Printf("[%s] All Tokens Successfully Updated\n", time.Now().Format("2006-01-02 15:04:05"))
		}()
	}
}*/

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
