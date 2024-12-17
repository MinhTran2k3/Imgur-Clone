package routes

import (
	"github.com/TranMinh2k3/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

// UserRoute sets up the user-related routes
func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
