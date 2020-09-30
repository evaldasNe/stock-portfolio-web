package Controllers

import (
	"fmt"
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

//CreateStock ... Create Stock
func CreateStock(c *gin.Context) {
	var stock Models.Stock
	c.BindJSON(&stock)
	err := Models.CreateStock(&stock)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, stock)
	}
}

//UpdateStock ... Update the stock information
func UpdateStock(c *gin.Context) {
	var stock Models.Stock
	id := c.Params.ByName("id")
	err := Models.GetStockByID(&stock, id)
	if err != nil {
		c.JSON(http.StatusNotFound, stock)
	}
	c.BindJSON(&stock)
	err = Models.UpdateStock(&stock, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stock)
	}
}

//DeleteStock ... Delete the stock
func DeleteStock(c *gin.Context) {
	var stock Models.Stock
	id := c.Params.ByName("id")
	err := Models.DeleteStock(&stock, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"message": "Stock has been deleted",
		})
	}
}
