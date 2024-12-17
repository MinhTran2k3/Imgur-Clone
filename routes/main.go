package routes

import (
	"github.com/TranMinh2k3/gin-gorm-rest/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
