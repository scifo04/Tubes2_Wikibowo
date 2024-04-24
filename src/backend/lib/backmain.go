package back

import (
	"github.com/gocolly/colly"
	"fmt"
	"strings"
	"time"
	"sync"
)

type LinkInfo struct {
	LinkValue string `json:"linkValue"`
	FinValue  string `json:"finValue"`
	IsOn      bool   `json:"isOn"`
	IsName    bool   `json:"isName"`
}

var namespace bool
var depth int
// var cache []string
var cache_checker map[string]bool
var end string
var start string

func Scrape(lenc []string, ch chan<- [][]string, wg *sync.WaitGroup,sem chan struct{}) { //ex: [start, first click] -> [start,first click, second click] , [start,first click, second click], ...
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
		if (!namespace) {
			if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") && !(strings.Contains(linkURL,"Category:") || strings.Contains(linkURL,"Wikipedia:") || strings.Contains(linkURL,"Special:") || strings.Contains(linkURL,"File:") || strings.Contains(linkURL,"Help:") || strings.Contains(linkURL,"Talk:") || strings.Contains(linkURL,"Portal:") || strings.Contains(linkURL,"Template:") || strings.Contains(linkURL,"Template_talk:") || strings.Contains(linkURL,"Main_Page")) {
				linkToAppend := append([]string{}, lenc...) // Create a new slice with the same elements as lenc
				if (!cache_checker[linkURL] || end == linkURL) && start != linkURL {
					cache_checker[linkURL] = true
					linkToAppend = append(linkToAppend, linkURL)
				}
				// fmt.Println(linkToAppend)
				if (len(linkToAppend) == depth+1) {
					links = append(links, linkToAppend)
				}
			}
		} else {
			if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") {
				linkToAppend := append([]string{}, lenc...) // Create a new slice with the same elements as lenc
				if (!cache_checker[linkURL] || end == linkURL) && start != linkURL {
					cache_checker[linkURL] = true
					linkToAppend = append(linkToAppend, linkURL)
				}
				// fmt.Println(linkToAppend)
				if (len(linkToAppend) == depth+1) {
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

func findElement(data [][]string,end string) int{
	for i := len(data) - 1; i >= 0; i-- {
		// fmt.Println(data[i][len(data[i])-1]," ",end)
		if (data[i][len(data[i]) - 1] == end){
			return i
		}
	}
	return -1
}

// func goPrint(data [][]string,end string) {
// 	for i := len(data) - 1; i >= 0; i-- {
// 		fmt.Println(data[i][len(data[i])-1]," ",end)
// 	}
// }

// func isInList(data []string, title string) bool {
// 	for i := 0; i < len(data); i++ {
// 		if data[i] == title {
// 			return true;
// 		}
// 	}
// 	return false
// }

func BFS(i int, start string, end string, listScrape [][]string, ch chan<- []string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	
	var tempScrape [][]string;

	if ((i >= 0 && len(listScrape[i]) == depth+1) || i == -1) {
		depth++
		fmt.Println("Layer ",depth)
	}
	fmt.Print(len(listScrape))
	
	if (i == -1){ //listScrape masih kosong
		fmt.Println(" ",start)
		chScrape := make(chan [][]string)
		go Scrape([]string{start},chScrape,wg,sem)
		tempScrape = <-chScrape
		listScrape = append(listScrape, tempScrape...)
	} else {
		fmt.Println(" ",listScrape[i])
		chScrape := make(chan[][] string)
		go Scrape(listScrape[i],chScrape,wg,sem)
		tempScrape = <-chScrape
		listScrape = append(listScrape, tempScrape...)
		// if (listScrape[i][len(listScrape[i])-1] == "https://en.wikipedia.org/wiki/Ichthyotitan") {
		// 	goPrint(tempScrape,end)
		// }
	}
	idxTemp := findElement(tempScrape,end)
	if (idxTemp != -1){
		wg.Add(1)
		fmt.Println("length : ",len(listScrape))
		ch <- tempScrape[idxTemp]
		return
	}
	sem <- struct{}{}
    defer func() { <-sem }()
    wg.Add(1)
    go BFS(i+1, start, end, listScrape, ch, wg, sem)
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
	if (strings.Contains(temp,"–")) {
		temp = strings.Replace(temp,"–","%E2%80%93",-1)
	}
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
	if (strings.Contains(temp,"%E2%80%93")) {
		temp = strings.Replace(temp,"%E2%80%93","–",-1)
	}
	fmt.Println(temp)
	return temp
}

func Back_Main(inpute LinkInfo, maxConcurrency int) ([]string, int64) {
	depth = 0
	fmt.Println("Layer ",depth)
	inpute.LinkValue = TurnToWikipedia(inpute.LinkValue)
	inpute.FinValue = TurnToWikipedia(inpute.FinValue)
	end = inpute.FinValue
	start = inpute.LinkValue
	namespace = inpute.IsName
	// cache = []string{}
	cache_checker = map[string]bool{}
	var listSolution []string
	ch := make(chan []string)
	var wg sync.WaitGroup
	sem := make(chan struct{},maxConcurrency)
	t1 := time.Now()
	if (inpute.IsOn) {
		go BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{},ch,&wg,sem)
		// IDS()
	} else {
		go BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{},ch,&wg,sem)
	}
	listSolution = <-ch
	for i := 0; i < len(listSolution); i++ {
		listSolution[i] = TurnToTitle(listSolution[i])
	}
	t3 := time.Since(t1)
	t2 := time.Since(t1).Milliseconds()
	// fmt.Println(cache)
	fmt.Println("Execution Time: ",t3)
	return listSolution,t2
}