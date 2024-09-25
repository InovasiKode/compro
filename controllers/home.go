package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomePageHandler adalah fungsi untuk menangani request ke halaman home ("/")
func HomePageHandler(c *gin.Context) {
	// Mengirimkan respon HTML dengan menggunakan template 'index.html'
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome to Home",
	})
}
