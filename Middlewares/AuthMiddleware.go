package Middlewares

import (
	"net/http"
	"regexp"

	"github.com/evaldasNe/stock-portfolio-web/Services"

	"github.com/evaldasNe/stock-portfolio-web/Models"
	"github.com/gin-gonic/gin"
)

type header struct {
	Auth string `header:"Authorization"`
}

//AuthMiddleware ...
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := header{}

		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		if h.Auth != "" {
			reg := regexp.MustCompile("Bearer ")
			headerToken := reg.ReplaceAllString(h.Auth, "${1}")
			var userToken Models.Token

			if err := Models.GetTokenByAccessToken(&userToken, headerToken); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "Token " + err.Error(),
				})
				return
			}

			if !userToken.Token.Valid() {
				if newToken, err := Services.GetNewToken(userToken.Token); err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"message": "Error while refreshing token. " + err.Error(),
					})
				} else {
					userToken.Token = newToken
					if err := Models.UpdateToken(&userToken); err != nil {
						c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
							"message": "Error while updating token. " + err.Error(),
						})
						return
					}
					c.AbortWithStatusJSON(http.StatusResetContent, userToken.Token)
				}

			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
