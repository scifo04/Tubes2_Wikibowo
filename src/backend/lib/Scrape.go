package back

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
	"sync"
)

func Scrape(lenc []string, eng Engine, ch chan<- [][]string, wg *sync.WaitGroup,sem chan struct{}) { //ex: [start, first click] -> [start,first click, second click] , [start,first click, second click], ...
	defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic in Scrape:", r)
        }
        wg.Done()
    }()
	
	var links [][]string
    c := colly.NewCollector()
		// colly.AllowedDomains("en.wikipedia.org","https://en.wikipedia.org","en.wikipedia.org/","https://en.wikipedia.org/",))

    c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting: ",lenc[len(lenc) - 1])
        r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
    })

    c.OnResponse(func(r *colly.Response) {
        // fmt.Println("Response Code:", r.StatusCode)
    })

    // to get the "a" tag
    c.OnHTML("a", func(e *colly.HTMLElement) {
		// Extract the text and href attribute of the <a> tag
		linkURL := e.Attr("href")

		if (!strings.Contains(linkURL,"https://en.wikipedia.org")) {
			linkURL = "https://en.wikipedia.org" + linkURL
		}

		// Print the extracted information
		if (!eng.Namespace) {
			if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") && !(strings.Contains(linkURL,"Category:") || strings.Contains(linkURL,"Wikipedia:") || strings.Contains(linkURL,"Special:") || strings.Contains(linkURL,"File:") || strings.Contains(linkURL,"Help:") || strings.Contains(linkURL,"Talk:") || strings.Contains(linkURL,"Portal:") || strings.Contains(linkURL,"Template:") || strings.Contains(linkURL,"Template_talk:") || strings.Contains(linkURL,"Main_Page")) {
				linkToAppend := append([]string{}, lenc...) // Create a new slice with the same elements as lenc
				if (!eng.Cache[linkURL] || eng.End == linkURL) && eng.Start != linkURL {
					eng.Cache[linkURL] = true
					linkToAppend = append(linkToAppend, linkURL)
				}
				// fmt.Println(linkToAppend)
				if (len(linkToAppend) == eng.Depth+1) {
					links = append(links, linkToAppend)
				}
			}
		} else {
			if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") {
				linkToAppend := append([]string{}, lenc...) // Create a new slice with the same elements as lenc
				if (!eng.Cache[linkURL] || eng.End == linkURL) && eng.Start != linkURL {
					eng.Cache[linkURL] = true
					linkToAppend = append(linkToAppend, linkURL)
				}
				// fmt.Println(linkToAppend)
				if (len(linkToAppend) == eng.Depth+1) {
					links = append(links, linkToAppend)
				}
			}
		}
	})

    c.Visit(lenc[len(lenc) - 1])

	// if (lenc[len(lenc)-1] == "https://en.wikipedia.org/wiki/Ichthyotitan") {
	// 	fmt.Println(links)
	// }
	sem <- struct{}{}
    defer func() { <-sem }()
	wg.Add(1)

	ch <- links
}