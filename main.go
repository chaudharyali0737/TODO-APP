package main

import (
	"api/db"
	todocrud "api/todoCrud"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
func main() {
	db.Init()
	err := db.TodoMigrations()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(CORSMiddleware())
	router.GET("/todos", todocrud.GetTodos)
	router.GET("/todo/:id", todocrud.GetSingleTodos)
	router.POST("/todo/insert", todocrud.Insert)
	router.POST("/todo/update", todocrud.Update)
	router.DELETE("/todo/deleteall", todocrud.DeleteAll)
	router.DELETE("/todo/delete/:id", todocrud.Delete)
	router.Run(":8080")
}
