package main

import (
	"github.com/gin-gonic/gin"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	router := gin.Default()

	router.POST("/name", func(c *gin.Context) {
		var p person
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{
			"first_name": p.FirstName,
			"last_name":  p.LastName,
		})
	})

	router.Run(":8081")
}
