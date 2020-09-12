package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

// Person struct
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})

	router.GET("/test", func(c *gin.Context) {
		users := make([]Person, 3)
		users[0] = Person{Name: "Bill", Age: 44}
		users[1] = Person{Name: "Jon", Age: 16}
		users[2] = Person{Name: "Joe", Age: 66}

		c.JSON(200, gin.H{
			"users": users,
		})
	})
	autotls.Run(router, "http://ec2-3-22-241-112.us-east-2.compute.amazonaws.com")
	//router.Run(":80")
}
