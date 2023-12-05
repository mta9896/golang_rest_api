package ginrest

import (
	_ "fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"mta9896/restapi/internal/entity"
	"mta9896/restapi/internal/database"
)

var validate *validator.Validate

func Initialize() {
	initDB()
	initValidator()
	initRouter()
}

func initRouter() {
	router := gin.Default()
	router.GET("/items", GetItems)
	router.POST("/items", CreateItem)
	router.Run("localhost:8080")
}

func initValidator() {
	validate = validator.New()
}

func initDB() {
	err := database.Initialize()

	if (err != nil) {
		log.Println("Error: " , err.Error())
		
		return
	}
}

func GetItems(c *gin.Context) {
	items, err := database.FetchAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item entity.Item

	// validation
	// if err := validate.Struct(&item); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    //     return
    // }

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.InsertItem(item)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, item)
}