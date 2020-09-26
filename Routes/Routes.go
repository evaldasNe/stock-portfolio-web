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
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
		grp1.GET("user/:id/soldStocks", Controllers.GetAllStocksUserSold)
		grp1.GET("user/:id/profit", Controllers.GetUserProfit)

		grp1.GET("stock", Controllers.GetStocks)
		grp1.POST("stock", Controllers.CreateStock)
		grp1.GET("stock/:id", Controllers.GetStockByID)
		grp1.PUT("stock/:id", Controllers.UpdateStock)
		grp1.DELETE("stock/:id", Controllers.DeleteStock)

		grp1.GET("ownedStock", Controllers.GetOwnedStocks)
		grp1.POST("ownedStock", Controllers.CreateOwnedStock)
		grp1.GET("ownedStock/:id", Controllers.GetOwnedStockByID)
		grp1.PUT("ownedStock/:id", Controllers.UpdateOwnedStock)
		grp1.DELETE("ownedStock/:id", Controllers.DeleteOwnedStock)

		grp1.GET("comment", Controllers.GetComments)
		grp1.POST("comment", Controllers.CreateComment)
		grp1.GET("comment/:id", Controllers.GetCommentByID)
		grp1.PUT("comment/:id", Controllers.UpdateComment)
		grp1.DELETE("comment/:id", Controllers.DeleteComment)
	}
	return r
}
