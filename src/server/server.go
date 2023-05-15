package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	apiURI := "/api/v1"

	router := gin.Default()

	router.GET(apiURI, func(c *gin.Context) {
		go func() {
			processRequest(c)
		}()
	})

	router.Run(":80")

	fmt.Println("Server running on port 80")
}

func processRequest(c *gin.Context) {
	// Perform the processing logic for the request here
	response := "Initial backend for betify!"
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}
