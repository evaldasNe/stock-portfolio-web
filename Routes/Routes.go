package Routes

import (
	"github.com/evaldasNe/stock-portfolio-web/Controllers"
	"github.com/evaldasNe/stock-portfolio-web/Middlewares"
	"github.com/evaldasNe/stock-portfolio-web/Services"
	ginsession "github.com/go-session/gin-session"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(ginsession.New())

	auth := r.Group("/auth")
	Services.InitGoogleAuth()
	{
		auth.GET("/googleLogin", Services.HandleGoogleLogin)
		auth.GET("/callback", Services.HandleGoogleCallback)
	}

	grp1 := r.Group("/api")
	grp1.Use(Middlewares.AuthMiddleware())
	{
		grp1.GET("users", Controllers.GetUsers)
		grp1.POST("users", Controllers.CreateUser)
		grp1.GET("users/:id", Controllers.GetUserByID)
		grp1.PATCH("users/:id", Controllers.UpdateUser)
		grp1.DELETE("users/:id", Controllers.DeleteUser)
		grp1.GET("users/:id/soldStocks", Controllers.GetAllStocksUserSold)
		grp1.GET("users/:id/profit", Controllers.GetUserProfit)

		grp1.GET("stocks", Controllers.GetStocks)
		grp1.POST("stocks", Controllers.CreateStock)
		grp1.GET("stocks/:id", Controllers.GetStockByID)
		grp1.PATCH("stocks/:id", Controllers.UpdateStock)
		grp1.DELETE("stocks/:id", Controllers.DeleteStock)

		grp1.GET("ownedStocks", Controllers.GetOwnedStocks)
		grp1.POST("ownedStocks", Controllers.CreateOwnedStock)
		grp1.GET("ownedStocks/:id", Controllers.GetOwnedStockByID)
		grp1.PATCH("ownedStocks/:id", Controllers.UpdateOwnedStock)
		grp1.DELETE("ownedStocks/:id", Controllers.DeleteOwnedStock)

		grp1.GET("comments", Controllers.GetComments)
		grp1.POST("comments", Controllers.CreateComment)
		grp1.GET("comments/:id", Controllers.GetCommentByID)
		grp1.PATCH("comments/:id", Controllers.UpdateComment)
		grp1.DELETE("comments/:id", Controllers.DeleteComment)
	}
	return r
}
