package routes

import (
	"github.com/gin-gonic/gin"
	"rest-api/controllers"
)

func RegisterTaskRoutes(router *gin.Engine) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("/", controllers.GetAllTasks)
		tasks.GET("/:id", controllers.GetTaskDetail)
		tasks.POST("/", controllers.CreateTask)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
