package controllers

import (
	"fmt"
	"lead/services"

	"github.com/gin-gonic/gin"
)

func Code(c *gin.Context) {
	paramCode := c.Param("code")

	// Scrape Data
	data := services.ScrapeURL()
	var title, moves string
	for _, v := range data {
		if v.Code == paramCode {
			title = v.Title
			moves = v.Move
		}
	}
	fmt.Println(paramCode)
	c.HTML(200, "Code.html", gin.H{
		"title": title,
		"moves": moves,
	})
}
