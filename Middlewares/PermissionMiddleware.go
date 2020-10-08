package Middlewares

import (
	"fmt"
	"net/http"

	"github.com/evaldasNe/stock-portfolio-web/Models"

	"github.com/gin-gonic/gin"
)

// IsTheSameUserOrIsAdmin ...
func IsTheSameUserOrIsAdmin(userID uint, authUserID uint) bool {
	if authUserID != userID {
		var userRole string
		err := Models.GetUserRoleByID(&userRole, authUserID)
		if err != nil {
			return false
		} else if userRole == "ADMIN" {
			return true
		}

		return false
	}
	return true
}

// UserRestriction ...
func UserRestriction() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("authUserID").(uint)
		id := c.Param("id")
		var userRole string

		err := Models.GetUserRoleByID(&userRole, userID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else if userRole != "ADMIN" {
			if c.Request.Method == "POST" {
				c.AbortWithStatus(http.StatusForbidden)
			} else if id != "" && c.Request.Method != "GET" {
				isTheSameUser := (fmt.Sprint(userID) == id)
				if !isTheSameUser {
					c.AbortWithStatus(http.StatusForbidden)
				}
			}
		}
		c.Set("authUserRole", userRole)
	}
}

// OnlyForAdmin ...
func OnlyForAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("authUserID").(uint)
		var userRole string

		err := Models.GetUserRoleByID(&userRole, userID)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else if userRole != "ADMIN" {
			c.AbortWithStatus(http.StatusForbidden)
		}
	}
}
