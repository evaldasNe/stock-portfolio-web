package Routes

import (
	"github.com/evaldasNe/stock-portfolio-web/Controllers"
	"github.com/evaldasNe/stock-portfolio-web/Middlewares"
	"github.com/evaldasNe/stock-portfolio-web/Services"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	/*r.Use(ginsession.New(), cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, PATCH, POST, DELETE",
		RequestHeaders:  "*",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))*/

	auth := r.Group("/auth")
	Services.InitGoogleAuth()
	{
		auth.GET("/googleLogin", Services.HandleGoogleLogin)
		auth.GET("/callback", Services.HandleGoogleCallback)
		auth.POST("/saveGoogleToken", Services.SaveGoogleToken)
	}

	grp1 := r.Group("/api")
	grp1.Use(Middlewares.AuthMiddleware())

	usersG := grp1.Group("/users")
	usersG.Use(Middlewares.UserRestriction())

	onlyForAdmin := grp1.Group("/")
	onlyForAdmin.Use(Middlewares.OnlyForAdmin())
	{
		// /api/users
		usersG.GET("/", Controllers.GetUsers)
		usersG.POST("/", Controllers.CreateUser)
		usersG.GET("/:id", Controllers.GetUserByID)
		usersG.PATCH("/:id", Controllers.UpdateUser)
		usersG.DELETE("/:id", Controllers.DeleteUser)
		usersG.GET("/:id/soldStocks", Controllers.GetAllStocksUserSold)
		usersG.GET("/:id/profit", Controllers.GetUserProfit)

		// /api
		grp1.GET("stocks", Controllers.GetStocks)
		onlyForAdmin.POST("stocks", Controllers.CreateStock)
		grp1.GET("stocks/:id", Controllers.GetStockByID)
		onlyForAdmin.PATCH("stocks/:id", Controllers.UpdateStock)
		onlyForAdmin.DELETE("stocks/:id", Controllers.DeleteStock)

		// /api
		grp1.GET("ownedStocks", Controllers.GetOwnedStocks)
		grp1.POST("ownedStocks", Controllers.CreateOwnedStock)
		grp1.GET("ownedStocks/:id", Controllers.GetOwnedStockByID)
		grp1.PATCH("ownedStocks/:id", Controllers.UpdateOwnedStock)
		grp1.DELETE("ownedStocks/:id", Controllers.DeleteOwnedStock)

		// /api
		grp1.GET("comments", Controllers.GetComments)
		grp1.POST("comments", Controllers.CreateComment)
		grp1.GET("comments/:id", Controllers.GetCommentByID)
		grp1.PATCH("comments/:id", Controllers.UpdateComment)
		grp1.DELETE("comments/:id", Controllers.DeleteComment)
	}
	return r
}
