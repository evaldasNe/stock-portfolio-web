package Models

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
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
