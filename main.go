package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "albums")
}
func main() {
	time.Now()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
