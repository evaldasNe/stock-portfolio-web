package Models

import (
	"fmt"
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
)

// OwnedStock model struct
type OwnedStock struct {
	ID             uint      `json:"id"`
	AmountOfShares float64   `gorm:"not null" json:"amount_of_shares"`
	Price          float64   `gorm:"not null" json:"price"`
	StockID        uint      `gorm:"not null" json:"stock_id"`
	UserID         uint      `gorm:"not null" json:"owner_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

//GetAllOwnedStocks Fetch all owend stocks data
func GetAllOwnedStocks(ownedStocks *[]OwnedStock) (err error) {
	if err = Config.DB.Find(ownedStocks).Error; err != nil {
		return err
	}
	return nil
}

//CreateOwnedStock ... Insert New data
func CreateOwnedStock(ownedStock *OwnedStock) (err error) {
	if err = Config.DB.Create(ownedStock).Error; err != nil {
		return err
	}
	return nil
}

//GetOwnedStockByID ... Fetch only one owned stock by ID
func GetOwnedStockByID(ownedStock *OwnedStock, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(ownedStock).Error; err != nil {
		return err
	}
	return nil
}

//UpdateOwnedStock ... Update Owned Stock
func UpdateOwnedStock(ownedStock *OwnedStock, id string) (err error) {
	fmt.Println(ownedStock)
	Config.DB.Save(ownedStock)
	return nil
}

//DeleteOwnedStock ... Delete Owned Stock
func DeleteOwnedStock(ownedStock *OwnedStock, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(ownedStock)
	return nil
}
