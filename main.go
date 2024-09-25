package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi gin router
	router := gin.Default()

	// Mengatur proxy yang dipercayai (opsional, hilangkan peringatan)
	router.SetTrustedProxies(nil) // atau atur IP proxy yang spesifik

	// Mengatur direktori 'templates' untuk HTML templates
	router.LoadHTMLGlob("templates/*.html")

	// Mengatur direktori 'static' untuk file statis (CSS, JS, images, dll.)
	router.Static("/assets", "./static/assets")

	// Tambahkan route handler untuk homepage
	router.GET("/", func(c *gin.Context) {
		// Render template index.html dengan data tambahan (jika diperlukan)
		c.HTML(200, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	// Menjalankan server pada port 8877
	router.Run(":8877")
}
