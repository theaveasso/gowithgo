package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	var quotes []Quote
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		div := e.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		// fmt.Printf("Quote: %s\nby ~ %s\n", quote, author)
		q := Quote{
			Quote:  quote,
			Author: author,
		}
		quotes = append(quotes, q)
	})

	c.OnRequest(func(req *colly.Request) {
		req.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36")
		fmt.Println("Visting ~", req.URL.String())
	})

	c.OnResponse(func(res *colly.Response) {
		fmt.Println("Status Code ~", res.StatusCode)
	})

	c.OnError(func(res *colly.Response, err error) {
		fmt.Println("Error", err.Error())
	})

	c.Visit("http://quotes.toscrape.com")
	fmt.Println(quotes)
}
