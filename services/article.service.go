package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/syed-ashraf123/go-crud/database"
	"github.com/syed-ashraf123/go-crud/models"
)

func CreatePost(c *gin.Context) {

	var article models.Article
	c.ShouldBindJSON(&article)

	validate := validator.New()
	// validation
	err := validate.Struct(article)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// saving in db
	result := database.DB.Create(&article)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": article})

}

func GetPost(c *gin.Context) {

	var article []models.Article
	result := database.DB.Find(&article)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"article": article})
}
