package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	router.Use(headersByRequestURI())
	router.Static("/assets", "./assets")
	router.StaticFile("/ads.txt", "./assets/txt/ads.txt")
	router.StaticFile("/app-ads.txt", "./assets/txt/app-ads.txt")
	router.StaticFile("/favicon.ico", "./assets/images/favicon.png")
	router.GET("/", showIndexPage)
	router.GET("/app/:app_id", showAppPage)
	router.GET("/about", showAboutPage)
	router.GET("/privacy", showPrivacyPage)
}

func headersByRequestURI() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/assets") {
			c.Header("Cache-Control", "max-age=2592000")
		}
	}
}
