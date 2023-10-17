package todocrud

import (
	"api/db"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        int        `json:"id"`
	Task      string     `json:"task"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

func GetTodos(c *gin.Context) {
	result := []Todo{}
	tx := db.DB.Model(db.Todos{}).Find(&result)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, &result)
}

func GetSingleTodos(c *gin.Context) {
	id := c.Param("id")
	result := []Todo{}
	tx := db.DB.Model(db.Todos{}).Where("id=?", id).Find(&result)
	if tx.Error != nil || tx.RowsAffected == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, &result)
}
func Insert(c *gin.Context) {
	input := db.Todos{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := fmt.Errorf("parsing payload error %w", err)
		c.AbortWithError(http.StatusBadRequest, errors.New(response.Error()))
	}
	tx := db.DB.Model(db.Todos{}).Create(&input)
	if tx.Error != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.AbortWithStatus(200)
}
func Update(c *gin.Context) {
	var todo_data Todo
	err := c.ShouldBindJSON(&todo_data)
	if err != nil {
		response := fmt.Errorf("parsing payload error %w", err)
		c.AbortWithError(http.StatusBadRequest, errors.New(response.Error()))
	}
	tx := db.DB.Model(db.Todos{}).Where("id=?", todo_data.ID).Update("task", todo_data.Task)
	if tx.Error != nil {
		response_error := fmt.Errorf("update db error %w", tx.Error)
		c.AbortWithError(http.StatusBadRequest, errors.New(response_error.Error()))
	}
	c.AbortWithStatus(200)
}
func Delete(c *gin.Context) {
	id := c.Param("id")
	var delete_entry *db.Todos
	interger_id, err := strconv.Atoi(id)
	if err != nil {
		response := fmt.Errorf("parsing payload error %w", err)
		c.AbortWithError(http.StatusBadRequest, errors.New(response.Error()))
	}
	tx := db.DB.Table("todos").Where("id=?", interger_id).Delete(&delete_entry)
	if tx.Error != nil {
		response_error := fmt.Errorf("delete db error %w", tx.Error)
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.New(response_error.Error()))
	}
	c.AbortWithStatus(200)
}
