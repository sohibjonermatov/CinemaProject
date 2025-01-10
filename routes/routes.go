package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() *gin.Engine {

	defaultRouter := gin.Default()

	defaultRouter.POST("/movie", CreateMovie)
	defaultRouter.GET("/movie", GetMovie)
	defaultRouter.GET("/movie/:id", GetMovieById)
	defaultRouter.PUT("/movie/:id", UpdateMovie)
	defaultRouter.DELETE("/movie/:id", DeleteMovie)

	defaultRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	return defaultRouter
}
