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
}

func Scrape(lenc []string, ch chan<- [][]string) { //ex: [start, first click] -> [start,first click, second click] , [start,first click, second click], ...
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
		if strings.Contains(linkURL, "https://en.wikipedia.org/wiki/") {
			linkToAppend := append([]string{}, lenc...) // Create a new slice with the same elements as lenc
			linkToAppend = append(linkToAppend, linkURL)
			links = append(links, linkToAppend)
		}
	})

    c.Visit(lenc[len(lenc) - 1])

	// if (lenc[len(lenc)-1] == "https://en.wikipedia.org/wiki/Ichthyotitan") {
	// 	fmt.Println(links)
	// }

	ch <- links
}

func findElement(data [][]string,end string) int{
	for i := len(data) - 1; i >= 0; i-- {
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

func BFS(i int, start string, end string, listScrape [][]string, ch chan<- []string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	var tempScrape [][]string;

	// fmt.Print(len(listScrape))
	
	if (i == -1){ //listScrape masih kosong
		// fmt.Println(" ",start)
		chScrape := make(chan [][]string)
		go Scrape([]string{start},chScrape)
		tempScrape = <-chScrape
		listScrape = append(listScrape, tempScrape...)
	} else {
		// fmt.Println(" ",listScrape[i])
		chScrape := make(chan[][] string)
		go Scrape(listScrape[i],chScrape)
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

func Back_Main(inpute LinkInfo, maxConcurrency int) ([]string, int64) {
	inpute.LinkValue = TurnToWikipedia(inpute.LinkValue)
	inpute.FinValue = TurnToWikipedia(inpute.FinValue)
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
	fmt.Println("Execution Time: ",t3,"ms")
	return listSolution,t2
}