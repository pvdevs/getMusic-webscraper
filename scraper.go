package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type album struct {
	Title 	string `json: title`
	Rating 	string `json: rating`
	Genre	string `json: genre`
	ImgUrl 	string `json: imgurl`
}

func main() {
	c := colly.NewCollector()
	c.OnHTML("div.review a.review__link div.review__title ul li", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("https://pitchfork.com/best/high-scoring-albums/")
}