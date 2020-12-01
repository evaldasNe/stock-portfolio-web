package Models

import (
	"fmt"
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
)

// OwnedStock model struct
type OwnedStock struct {
	ID             uint      `json:"id"`
	AmountOfShares float64   `gorm:"not null;<-:create" json:"amount_of_shares" binding:"required"`
	Price          float64   `gorm:"not null;<-:create" json:"price" binding:"required"`
	StockID        uint      `gorm:"not null;<-:create" json:"stock_id" binding:"required"`
	UserID         uint      `gorm:"not null;<-:create" json:"owner_id" binding:"required"`
	Sold           bool      `grom:"not null;default:false" json:"is_sold"`
	Profit         float64   `json:"profit"`
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
func UpdateOwnedStock(ownedStock *OwnedStock) (err error) {
	fmt.Println(ownedStock)
	Config.DB.Save(ownedStock)
	return nil
}

//DeleteOwnedStock ... Delete Owned Stock
func DeleteOwnedStock(ownedStock *OwnedStock) (err error) {
	Config.DB.Delete(ownedStock)
	return nil
}

//GetAllStocksUserSold ... Fetch all stocks user sold
func GetAllStocksUserSold(soldStocks *[]OwnedStock, userID string) (err error) {
	if err = Config.DB.Where("user_id = ? AND sold = ?", userID, true).Find(soldStocks).Error; err != nil {
		return err
	}
	return nil
}

//GetUserProfit ... Get User profit
func GetUserProfit(profit *float64, userID string) (err error) {
	query := "SELECT SUM(profit) FROM owned_stocks WHERE user_id = ? AND sold = ?"
	if err = Config.DB.Raw(query, userID, true).Scan(profit).Error; err != nil {
		return err
	}
	return nil
}
