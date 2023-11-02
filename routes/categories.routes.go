package routes

import (
	"github.com/gin-gonic/gin"
	"luizafer.com/pvd/controllers"
)

func CategoriesRoutes(router *gin.Engine) *gin.RouterGroup {
	// tweetController := controllers.NewTweetController()
	categories := router.Group("/categories")
	{
		categories.GET("/list", controllers.ListCategories)
		categories.POST("/create", controllers.CreateCategory)
		categories.PUT("/:id", controllers.UpdateCategory)
		categories.DELETE("/:id", controllers.DeleteCategory)
	}
	return categories
}
