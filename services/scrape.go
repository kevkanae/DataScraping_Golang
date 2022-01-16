package services

import (
	"fmt"

	"github.com/gocolly/colly"
)

type ChessEco struct {
	Code  string `json:"Code"`
	Title string `json:"Title"`
	Move  string `json:"Move"`
}

func ScrapeURL() []ChessEco {

	var moves []string
	var code []string
	var title []string

	url := "https://www.chessgames.com/chessecohelp.html"

	//Init Data Scraper
	res := colly.NewCollector()

	//Scrape
	res.OnHTML("table", func(e *colly.HTMLElement) {

		// Save Move's Code
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			code = append(code, el.ChildText("td:nth-child(1)"))
		})

		//Save Moves
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			moves = append(moves, el.ChildText("td:nth-child(2)>font>font"))
		})

		//Save Title
		e.ForEach("tr>td>font>b", func(_ int, el *colly.HTMLElement) {
			title = append(title, el.Text)
		})
		fmt.Println("Scrapping Complete")
	})
	res.Visit(url)

	// Array of Objects
	var arr []ChessEco
	for i := 0; i < len(moves); i++ {
		arr = append(arr, ChessEco{code[i], title[i], moves[i]})
	}

	return arr
}
