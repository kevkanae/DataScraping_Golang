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
	res := colly.NewCollector()
	res.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			code = append(code, el.ChildText("td:nth-child(1)"))
		})
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			moves = append(moves, el.ChildText("td:nth-child(2)>font>font"))
		})
		e.ForEach("tr>td>font>b", func(_ int, el *colly.HTMLElement) {
			title = append(title, el.Text)
		})
		fmt.Println("Scrapping Complete")
	})
	res.Visit(url)

	var arr []ChessEco
	for i := 0; i < len(moves); i++ {
		arr = append(arr, ChessEco{code[i], title[i], moves[i]})
	}

	return arr
}
