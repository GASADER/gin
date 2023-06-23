package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	fmt.Println("Test")
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	r := gin.Default()
	r.GET("/", test)
	r.Run()
}
