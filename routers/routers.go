package routers

import (
	"gin-gionic-demo/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/posts", controllers.CreatePost)
	router.GET("/post/:id", controllers.GetPostById)
	router.GET("/posts", controllers.GetPosts)

	return router
}
