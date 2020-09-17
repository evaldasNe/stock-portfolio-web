package Models

import (
	"fmt"
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
)

// User model struct
type User struct {
	ID          uint         `json:"id"`
	Email       string       `gorm:"unique;not null" json:"email"`
	FirstName   string       `gorm:"not null;size:50" json:"first_name"`
	LastName    string       `gorm:"not null;size:50" json:"last_name"`
	Phone       string       `json:"phone"`
	Address     string       `json:"address"`
	OwnedStocks []OwnedStock `json:"owned_stocks"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

//GetAllUsers Fetch all users data
func GetAllUsers(users *[]User) (err error) {
	if err = Config.DB.Preload("OwnedStocks").Find(users).Error; err != nil {
		return err
	}
	return nil
}

//CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Preload("OwnedStocks").Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}
