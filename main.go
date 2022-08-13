package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/hello", hello)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Hello world!")
}
