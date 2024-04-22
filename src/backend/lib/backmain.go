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

func Scrape(lenc []string) [][]string { //ex: [start, first click] -> [start,first click, second click] , [start,first click, second click], ...
	var links [][]string
    c := colly.NewCollector()
		// colly.AllowedDomains("en.wikipedia.org","https://en.wikipedia.org","en.wikipedia.org/","https://en.wikipedia.org/",))

    c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ",lenc[len(lenc) - 1])
        r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Response Code:", r.StatusCode)
    })

    // to get the "a" tag
    c.OnHTML("a", func(e *colly.HTMLElement) {
		// Extract the text and href attribute of the <a> tag
		linkURL := e.Attr("href")
		if (!strings.Contains(linkURL,"https://en.wikipedia.org/wiki/")) {
			linkURL = "https://en.wikipedia.org" + linkURL
		}

		// Print the extracted information
		if (strings.Contains(linkURL,"https://en.wikipedia.org/wiki/")) {
			linkToAppend := append(lenc, linkURL)
			links = append(links, linkToAppend)
		}
	})

    c.Visit(lenc[len(lenc) - 1])

	return links
}

func findElement(data [][]string,end string) int{
	for i := len(data) - 1; i >= 0; i-- {
		if (data[i][len(data[i]) - 1] == end){
			return i
		}
	}
	return -1
}

func BFS(i int, start string, end string, listScrape [][]string) []string{
	var tempScrape [][]string;
	
	if (i == -1){ //listScrape masih kosong
		tempScrape = Scrape([]string{start})
		listScrape = append(listScrape, tempScrape...)
	} else {
		tempScrape = Scrape(listScrape[i])
		listScrape = append(listScrape, tempScrape...)
	}
	idxTemp := findElement(tempScrape,end)
	if (idxTemp != -1){
		fmt.Println("length : ",len(listScrape))
		return tempScrape[idxTemp]
	}
	return BFS(i+1,start,end,listScrape)
}

func TurnToWikipedia(title string) string {
	tokens := strings.Split(title, " ")
	var temp string = ""
	for i := 0; i < len(tokens); i++ {
		temp = temp + tokens[i]
		if (i < len(tokens)-1) {
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
		if (i < len(tokens)-1) {
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
	if (inpute.IsOn) {
		listSolution = BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{})
		// IDS()
	} else {
		listSolution = BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{})
	}
	t2 := time.Since(t1).Milliseconds()
	for i := 0; i < len(listSolution); i++ {
		listSolution[i] = TurnToTitle(listSolution[i])
	}
	return listSolution,t2
}