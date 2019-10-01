package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	apps := getAllApps()
	presences := getAllPresences()
	render(c, gin.H{
		"title":     "TheInvader360 - Home",
		"payload":   apps,
		"presences": presences,
		"year":      time.Now().UTC().Year()},
		"index.html")
}

func showAppPage(c *gin.Context) {
	if appID, err := strconv.Atoi(c.Param("app_id")); err == nil {
		if app, err := getAppByID(appID); err == nil {
			presences := getAllPresences()
			render(c, gin.H{
				"title":     "TheInvader360 - " + app.Name,
				"payload":   app,
				"presences": presences,
				"year":      time.Now().UTC().Year()},
				"app.html")
		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func showAboutPage(c *gin.Context) {
	presences := getAllPresences()
	render(c, gin.H{
		"title":     "TheInvader360 - About",
		"presences": presences,
		"year":      time.Now().UTC().Year()},
		"about.html")
}

func showPrivacyPage(c *gin.Context) {
	presences := getAllPresences()
	render(c, gin.H{
		"title":     "TheInvader360 - Privacy",
		"presences": presences,
		"year":      time.Now().UTC().Year()},
		"privacy.html")
}
