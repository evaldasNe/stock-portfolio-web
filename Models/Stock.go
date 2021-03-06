package Models

import (
	"fmt"
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
)

// Stock model struct
type Stock struct {
	ID           uint      `json:"id"`
	Ticker       string    `gorm:"unique;not null" json:"ticker" binding:"required"`
	CompanyName  string    `gorm:"not null" json:"company_name" binding:"required"`
	SellingPrice float64   `gorm:"not null" json:"selling_price" binding:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

//GetAllStocks Fetch all stocks data
func GetAllStocks(stocks *[]Stock) (err error) {
	if err = Config.DB.Find(stocks).Error; err != nil {
		return err
	}
	return nil
}

//CreateStock ... Insert New data
func CreateStock(stock *Stock) (err error) {
	if err = Config.DB.Create(stock).Error; err != nil {
		return err
	}
	return nil
}

//GetStockByID ... Fetch only one stock by ID
func GetStockByID(stock *Stock, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(stock).Error; err != nil {
		return err
	}
	return nil
}

//UpdateStock ... Update stock
func UpdateStock(stock *Stock, id string) (err error) {
	fmt.Println(stock)
	Config.DB.Save(stock)
	return nil
}

//DeleteStock ... Delete stock
func DeleteStock(stock *Stock, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(stock)
	return nil
}
