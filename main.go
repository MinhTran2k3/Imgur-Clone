package main

import (
	"github.com/TranMinh2k3/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.UserRoute(router)
	router.Run(":8080")
}
