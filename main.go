package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"home.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})

	router.GET("/login", func(c *gin.Context) {

		c.HTML(
			http.StatusOK,
			"login.html",
			gin.H{
				"title": "Login Page",
			},
		)
	})

	router.Run()
}
