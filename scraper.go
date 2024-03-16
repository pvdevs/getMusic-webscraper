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
	// Get title
	c.OnHTML("div.review", func(h *colly.HTMLElement) {
		urls := []string {}
		//teste := h.ChildAttrs("a", "href")
		urls = append(urls, h.ChildAttr("a","href"))
		// Before each url insert -> https://pitchfork.com/

		for _,v := range urls{

			h.Request.Visit("https://pitchfork.com"+v)
			//fmt.Println(v)

		}
	})

	c.Wait()

	c.OnHTML("div.ScoreCircle-jAxRuP p", func(h *colly.HTMLElement) {
		contentText := h.Text
		if contentText != ""{
			fmt.Println(h.Text)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL.String())
	})

	c.Visit("https://pitchfork.com/best/high-scoring-albums/")
}