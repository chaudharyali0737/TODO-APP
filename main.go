package main

import (
	"api/db"
	todocrud "api/todoCrud"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	err := db.TodoMigrations()
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/todos", todocrud.GetTodos)
	router.GET("/todo/:id", todocrud.GetSingleTodos)
	router.POST("/todo/insert", todocrud.Insert)
	router.POST("/todo/update", todocrud.Update)
	router.DELETE("/todo/delete/:id", todocrud.Delete)
	router.Run("localhost:8080")
}
