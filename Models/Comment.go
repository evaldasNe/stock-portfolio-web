package Models

import (
	"time"

	"github.com/evaldasNe/stock-portfolio-web/Config"
)

// Comment model struct
type Comment struct {
	ID         uint      `json:"id"`
	AuthorID   uint      `gorm:"not null" json:"author_id"`
	ReceiverID uint      `gorm:"not null" json:"receiver_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

//GetAllComments Fetch all comments
func GetAllComments(comments *[]Comment) (err error) {
	if err = Config.DB.Find(comments).Error; err != nil {
		return err
	}
	return nil
}

//CreateComment ... Insert New data
func CreateComment(comment *Comment) (err error) {
	if err = Config.DB.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

//GetCommentByID ... Fetch only one comment by ID
func GetCommentByID(comment *Comment, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(comment).Error; err != nil {
		return err
	}
	return nil
}

//UpdateComment ... Update Comment
func UpdateComment(comment *Comment) (err error) {
	Config.DB.Save(comment)
	return nil
}

//DeleteComment ... Delete Comment
func DeleteComment(comment *Comment, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(comment)
	return nil
}
