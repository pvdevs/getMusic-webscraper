/*
func main() {
	c := colly.NewCollector()
	// Get title
	c.OnHTML("div.review", func(h *colly.HTMLElement) {
		urls := []string {}
		urls = append(urls, h.ChildAttr("a","href"))

		for _,v := range urls{

			h.Request.Visit("https://pitchfork.com"+v)

		}
	})

	c.Wait()

	c.OnHTML("div.ScoreCircle-jAxRuP  p", func(h *colly.HTMLElement) {
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
*/
