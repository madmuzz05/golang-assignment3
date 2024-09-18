package router

import (
	"update-data/controllers"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/log")

	userRouter.POST("/update_status", controllers.UpdateStatus)

	return r
}
