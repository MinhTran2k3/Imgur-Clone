package routes

import (
	"github.com/MinhTran2k3/Imgur-clone/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}
