package ids

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
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
		fmt.Println("Visiting: ", lenc)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code:", r.StatusCode)
	})

	// to get the "a" tag
	c.OnHTML("a", func(e *colly.HTMLElement) {
		// Extract the text and href attribute of the <a> tag
		linkURL := e.Attr("href")
		if !strings.Contains(linkURL, "https://en.wikipedia.org") {
			linkURL = "https://en.wikipedia.org" + linkURL
		}

		// Print the extracted information
		if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") {
			links = append(links, linkURL)
		}
	})

	c.Visit(lenc)

	return links
}

func DLS(src string, target string, limit int) ([]string, bool) {
	if limit == 0 && src == target {
		return []string{src}, true
	} else if limit == 0 && src != target {
		return nil, false
	}

	children := Scrape(src)
	for _, child := range children {
		fmt.Print("child: " + child + "\n")
		if child == target {
			return []string{src, target}, true
		}

		if child != src {
			path, found := DLS(child, target, limit-1)
			if found {
				return append([]string{src}, path...), true
			}
		}
	}
	return nil, false
}

func IDS(src string, target string) ([]string, bool) {
	depth := 0

	var path []string
	found := false

	for {
		fmt.Printf("Depth : %d\n", depth)
		path, found = DLS(src, target, depth)
		if found {
			return path, true
		}
		depth++
	}
}

func TurnToWikipedia(title string) string {
	tokens := strings.Split(title, " ")
	var temp string = ""
	for i := 0; i < len(tokens); i++ {
		temp = temp + tokens[i]
		if i < len(tokens)-1 {
			temp = temp + "_"
		}
	}
	temp = "https://en.wikipedia.org/wiki/" + temp
	return temp
}

func TurnToTitle(url string) string {
	url = url[30 : len(url)-0]
	temp := ""
	tokens := strings.Split(url, "_")
	for i := 0; i < len(tokens); i++ {
		temp = temp + tokens[i]
		if i < len(tokens)-1 {
			temp = temp + " "
		}
	}
	return temp
}

func Back_Main(inpute LinkInfo) ([]string, int64) {
	inpute.LinkValue = TurnToWikipedia(inpute.LinkValue)
	inpute.FinValue = TurnToWikipedia(inpute.FinValue)
	var listSolution []string

	t1 := time.Now()

	if inpute.IsOn {
		listSolution, _ = IDS(inpute.LinkValue, inpute.FinValue)
		// IDS(
	} else {
		listSolution, _ = IDS(inpute.LinkValue, inpute.FinValue)
	}
	t2 := time.Since(t1).Milliseconds()

	for i := 0; i < len(listSolution); i++ {
		listSolution[i] = TurnToTitle(listSolution[i])
	}

	return listSolution, t2
}
