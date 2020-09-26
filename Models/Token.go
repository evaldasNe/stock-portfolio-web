package Models

import (
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
	"golang.org/x/oauth2"
)

// Token model struct
type Token struct {
	ID uint `json:"id"`
	*oauth2.Token
	UserID    uint      `gorm:"not null;<-:create" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//GetAllTokens Fetch all tokens
func GetAllTokens(tokens *[]Token) (err error) {
	if err = Config.DB.Find(tokens).Error; err != nil {
		return err
	}
	return nil
}

//CreateOrUpdateToken ... Insert New data
func CreateOrUpdateToken(token *Token) (err error) {
	var existingToken Token
	if err = GetTokenByUserID(&existingToken, token.UserID); err != nil {
		if err = Config.DB.Create(token).Error; err != nil {
			return err
		}
	} else {
		token.ID = existingToken.ID
		token.CreatedAt = existingToken.CreatedAt
		if err = UpdateToken(token); err != nil {
			return err
		}
	}

	return nil
}

//GetTokenByUserID ... Fetch only one token by user ID
func GetTokenByUserID(token *Token, userID uint) (err error) {
	if err = Config.DB.Where("user_id = ?", userID).First(token).Error; err != nil {
		return err
	}
	return nil
}

//GetTokenByAccessToken ... Fetch only one token by access token
func GetTokenByAccessToken(token *Token, accessToken string) (err error) {
	if err = Config.DB.Where("access_token = ?", accessToken).First(token).Error; err != nil {
		return err
	}
	return nil
}

//UpdateToken ... Update token
func UpdateToken(token *Token) (err error) {
	Config.DB.Save(token)
	return nil
}

//DeleteToken ... Delete token
func DeleteToken(token *Token, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(token)
	return nil
}
