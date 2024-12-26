package main

import (
	"net/http"
	"os"

	"github.com/MinhTran2k3/Imgur-clone/handlers"
	"github.com/gin-gonic/gin"
)

var images []string

func main() {
	r := gin.Default()
	r.Static("/images", "./images")

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	r.POST("/upload", handlers.UploadImage)
	r.GET("/image/:id", handlers.GetImage)

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"images": images})
	})

	r.POST("/upload-image", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Fehler beim Hochladen der Datei.")
			return
		}

		os.Mkdir("images", os.ModePerm)
		filePath := "images/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.String(http.StatusInternalServerError, "Fehler beim Speichern der Datei")
			return
		}

		images = append(images, file.Filename)
		c.Redirect(http.StatusFound, "/")
	})

	r.POST("/delete-image", func(c *gin.Context) {
		filename := c.PostForm("filename")
		filePath := "images/" + filename
		os.Remove(filePath)

		for i, img := range images {
			if img == filename {
				images = append(images[:i], images[i+1:]...)
				break
			}
		}
		c.Redirect(http.StatusFound, "/")
	})

	if err := r.Run(":8080"); err != nil {
		panic("Fehler beim Starten des Servers: " + err.Error())
	}
}
