package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/syed-ashraf123/go-crud/services"
)

func ArticleController(r *gin.Engine) {

	r.POST("/create-post", services.CreatePost)
	r.GET("/get-posts", services.GetPost)
}
