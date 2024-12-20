package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Bilder wurde erfolgreich hochgeladen",
	})
}

func GetImage(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id":      id,
		"message": "Bilder wurde erfolgreich abgerufen",
	})
}
