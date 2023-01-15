package router

import (
	"github.com/gin-gonic/gin"
	"github.com/syed-ashraf123/go-crud/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(gin.Recovery())
	controller.ArticleController(r)

	return r
}
