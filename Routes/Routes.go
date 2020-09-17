package Routes

import (
	"github.com/evaldasNe/stock-portfolio-web/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)

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
	}
	return r
}
