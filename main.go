package main

import (
	"github.com/MinhTran2k3/Imgur-clone/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/upload", handlers.UploadImage)
	r.GET("/image/:id", handlers.GetImage)

	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
