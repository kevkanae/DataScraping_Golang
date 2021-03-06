package controllers

import (
	"lead/services"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// Scrape Data
	data := services.ScrapeURL()

	//Display Scraped Data
	c.HTML(200, "Display.html", gin.H{
		"Array": data,
	})
}
