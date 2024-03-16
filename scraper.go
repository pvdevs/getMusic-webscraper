package main

import (
	"github.com/gocolly/colly"
)

type album struct {
	Title 	string `json: title`
	Rating 	string `json: rating`
	ImgUrl 	string `json: imgurl`
}

func main() {
	c := colly.NewCollector()
}