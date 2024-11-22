package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Image struct {
	Filename string
}

var images []Image

func main() {

	if _, err := os.Stat("./images"); os.IsNotExist(err) {
		os.Mkdir("./images", os.ModePerm)
	}

	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images"))))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.HandleFunc("/delete", handleDelete)

	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, images)
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		file, header, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed to upload file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		filename := filepath.Join("images", header.Filename)
		destFile, err := os.Create(filename)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer destFile.Close()

		_, err = destFile.ReadFrom(file)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		images = append(images, Image{Filename: header.Filename})

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		filename := r.FormValue("filename")

		os.Remove(filepath.Join("images", filename))

		for i, img := range images {
			if img.Filename == filename {
				images = append(images[:i], images[i+1:]...)
				break
			}
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
