package Controllers

import (
	"fmt"
	"net/http"

	"github.com/evaldasNe/stock-portfolio-web/Models"

	"github.com/gin-gonic/gin"
)

//GetComments ... Get all comments
func GetComments(c *gin.Context) {
	var comments []Models.Comment
	err := Models.GetAllComments(&comments)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, comments)
	}
}

//CreateComment ... Create Comment
func CreateComment(c *gin.Context) {
	var comment Models.Comment
	c.BindJSON(&comment)
	err := Models.CreateComment(&comment)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, comment)
	}
}

//GetCommentByID ... Get the comment by id
func GetCommentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var comment Models.Comment
	err := Models.GetCommentByID(&comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

//UpdateComment ... Update comment information
func UpdateComment(c *gin.Context) {
	var comment Models.Comment
	id := c.Params.ByName("id")
	err := Models.GetCommentByID(&comment, id)
	if err != nil {
		c.JSON(http.StatusNotFound, comment)
	}
	c.BindJSON(&comment)
	err = Models.UpdateComment(&comment)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, comment)
	}
}

//DeleteComment ... Delete comment
func DeleteComment(c *gin.Context) {
	var comment Models.Comment
	id := c.Params.ByName("id")
	err := Models.DeleteComment(&comment, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"comment_id": id,
			"message":    "Comment has been deleted",
		})
	}
}
