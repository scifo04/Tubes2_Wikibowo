package back

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
	"time"
)

type LinkInfo struct {
	LinkValue string `json:"linkValue"`
	FinValue  string `json:"finValue"`
	IsOn      bool   `json:"isOn"`
}

func Scrape(lenc string) []string {
	var links []string
    c := colly.NewCollector()
		// colly.AllowedDomains("en.wikipedia.org","https://en.wikipedia.org","en.wikipedia.org/","https://en.wikipedia.org/",))

    c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ",lenc)
        r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Response Code:", r.StatusCode)
    })

    // to get the "a" tag
    c.OnHTML("a", func(e *colly.HTMLElement) {
		// Extract the text and href attribute of the <a> tag
		linkURL := e.Attr("href")
		if (!strings.Contains(linkURL,"https://en.wikipedia.org")) {
			linkURL = "https://en.wikipedia.org" + linkURL
		}

		// Print the extracted information
		if (strings.Contains(linkURL,"https://en.wikipedia.org/wiki/")) {
			links = append(links, linkURL)
		}
	})

    c.Visit(lenc)

	return links
}

func Back_Main(inpute LinkInfo) []string {
	t1 := time.Now()
	stuff := Scrape(inpute.LinkValue)
	t2 := time.Since(t1)
	fmt.Println(len(stuff))
	fmt.Println(t2.Milliseconds())
	return stuff
}