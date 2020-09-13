package Controllers

import (
	"net/http"

	"github.com/evaldasNe/stock-portfolio-web/Models"

	"github.com/gin-gonic/gin"
)

//GetStocks ... Get all stocks
func GetStocks(c *gin.Context) {
	var stocks []Models.Stock
	err := Models.GetAllStocks(&stocks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stocks)
	}
}

//GetStockByID ... Get the stock by id
func GetStockByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var stock Models.Stock
	err := Models.GetStockByID(&stock, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stock)
	}
}
