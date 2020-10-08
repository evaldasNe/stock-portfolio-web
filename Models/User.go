package Models

import (
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
	"gorm.io/gorm/clause"
)

// User model struct
type User struct {
	ID               uint         `json:"id"`
	Email            string       `gorm:"unique;not null" json:"email"`
	FirstName        string       `gorm:"not null;size:50" json:"first_name"`
	LastName         string       `gorm:"not null;size:50" json:"last_name"`
	Blocked          bool         `gorm:"not null;default:false" json:"blocked"`
	Role             string       `gorm:"not null;default:USER" json:"role"`
	OwnedStocks      []OwnedStock `json:"owned_stocks"`
	AuthorOfComments []Comment    `gorm:"foreignKey:AuthorID" json:"author_of_comments"`
	ReceivedComments []Comment    `gorm:"foreignKey:ReceiverID" json:"received_comments"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

//GetAllUsers Fetch all users data
func GetAllUsers(users *[]User) (err error) {
	if err = Config.DB.Preload(clause.Associations).Find(users).Error; err != nil {
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
	if err = Config.DB.Preload(clause.Associations).Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

//UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}

//DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

//GetUserByEmail ... Get User id
func GetUserByEmail(user *User, email string) (err error) {
	if err = Config.DB.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

//GetOrCreateUser ... Get or create new user
func GetOrCreateUser(user *User, userToLookFor User) (err error) {
	if err = Config.DB.FirstOrCreate(&user, userToLookFor).Error; err != nil {
		return err
	}
	return nil
}

// GetUserRoleByID ...
func GetUserRoleByID(role *string, id uint) (err error) {
	row := Config.DB.Table("users").Where("id = ?", id).Select("role").Row()
	if err = row.Scan(role); err != nil {
		return err
	}
	return nil
}
