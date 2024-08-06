package main

// TODO host on `deb01`
// - [ ] nginx reverse proxy
// - [ ] deploy pipeline

import (
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/skip2/go-qrcode"
)

const PORT = "4242"

// main
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("*.html")

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.GET("/qr", handleQr)

	r.StaticFile("/hash-favicon.png", "./hash-favicon.png")

	fmt.Printf("\033[32mhttp://localhost:%s\033[0m\n", PORT)
	r.Run(":" + PORT)
}

// Handle the /qr route
func handleQr(c *gin.Context) {
	str := c.Query("str")

	qr, err := qrcode.New(str, qrcode.Medium)
	if err != nil {
		c.String(500, "Error generating QR code: %v", err)
		return
	}

	png, err := qr.PNG(256)
	if err != nil {
		c.String(500, "Error generating PNG: %v", err)
		return
	}

	// Encode the PNG data to base64
	qrImage := base64.StdEncoding.EncodeToString(png)

	data := gin.H{
		"qrImage": qrImage,
		"text":    str,
	}

	c.HTML(200, "qr.html", data)
}
