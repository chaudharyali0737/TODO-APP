package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	time_now := time.Now().String()
	c.IndentedJSON(http.StatusOK, "albums    "+time_now)
}
func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
