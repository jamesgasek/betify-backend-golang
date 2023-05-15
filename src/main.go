package main

import (
	"fmt"
	"net/http"

	//gin
	"github.com/gin-gonic/gin"
)

func main() {

	api_uri := "/api/v1"

	router := gin.Default()

	router.GET(api_uri, func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})

	})

	router.Run(":8080")

	fmt.Println("Server running on port 8080")

	//router.Run()

}
