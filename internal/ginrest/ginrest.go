package ginrest

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"mta9896/restapi/internal/crud"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func HandleRequests() {
	validate = validator.New()

	router := gin.Default()
	router.GET("/items", GetItems)
	router.POST("/items", CreateItem)
	router.Run("localhost:8080")
}

func GetItems(c *gin.Context) {
	items := crud.List()
	c.IndentedJSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item crud.Item

	if err := validate.Struct(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	crud.Create(item)

	c.IndentedJSON(http.StatusCreated, item)
}