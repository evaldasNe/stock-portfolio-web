package Controllers

import (
	"fmt"
	"net/http"

	"github.com/evaldasNe/stock-portfolio-web/Middlewares"

	"github.com/evaldasNe/stock-portfolio-web/Models"

	"github.com/gin-gonic/gin"
)

//GetOwnedStocks ... Get all owned stocks
func GetOwnedStocks(c *gin.Context) {
	var ownedStocks []Models.OwnedStock
	err := Models.GetAllOwnedStocks(&ownedStocks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ownedStocks)
	}
}

//GetOwnedStockByID ... Get the owned stock by id
func GetOwnedStockByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var ownedStock Models.OwnedStock
	err := Models.GetOwnedStockByID(&ownedStock, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ownedStock)
	}
}

//CreateOwnedStock ... Create Owned Stock
func CreateOwnedStock(c *gin.Context) {
	var ownedStock Models.OwnedStock
	c.BindJSON(&ownedStock)

	if ownedStock.UserID != c.MustGet("authUserID").(uint) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	err := Models.CreateOwnedStock(&ownedStock)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, ownedStock)
	}
}

//UpdateOwnedStock ... Update owned stock information
func UpdateOwnedStock(c *gin.Context) {
	var ownedStock Models.OwnedStock
	id := c.Params.ByName("id")
	err := Models.GetOwnedStockByID(&ownedStock, id)
	if err != nil {
		c.JSON(http.StatusNotFound, ownedStock)
	}

	if !Middlewares.IsTheSameUserOrIsAdmin(ownedStock.UserID, c.MustGet("authUserID").(uint)) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	c.BindJSON(&ownedStock)
	err = Models.UpdateOwnedStock(&ownedStock)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, ownedStock)
	}
}

//DeleteOwnedStock ... Delete owned stock
func DeleteOwnedStock(c *gin.Context) {
	var ownedStock Models.OwnedStock
	id := c.Params.ByName("id")

	err := Models.GetOwnedStockByID(&ownedStock, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if !Middlewares.IsTheSameUserOrIsAdmin(ownedStock.UserID, c.MustGet("authUserID").(uint)) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	err = Models.DeleteOwnedStock(&ownedStock)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user_id": id,
			"message": "OwnedStock has been deleted",
		})
	}
}

//GetAllStocksUserSold ... Get stocks user sold
func GetAllStocksUserSold(c *gin.Context) {
	userID := c.Params.ByName("id")
	var soldStocks []Models.OwnedStock
	err := Models.GetAllStocksUserSold(&soldStocks, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, soldStocks)
	}
}

//GetUserProfit ... Get user profit
func GetUserProfit(c *gin.Context) {
	userID := c.Params.ByName("id")
	var profit float64
	err := Models.GetUserProfit(&profit, userID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"user_profit": profit,
		})
	}
}
