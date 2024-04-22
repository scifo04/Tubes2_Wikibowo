package main

import (

	"fmt"
	"strings"
	"time"
	"github.com/gocolly/colly"
)


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

func search(i int, start string, end string, listScrape [][]string) []string{
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
	return search(i+1,start,end,listScrape)
}

func main(){
	fmt.Print("Masukkan content Wikipedia awal yang mau discrape: ")
	var inputStart string
	var inputEnd string
	fmt.Scanln(&inputStart)

	fmt.Print("Masukkan content Wikipedia akhir uang mau discrape: ")
	fmt.Scanln(&inputEnd)

	startTime := time.Now()
	var template string = "https://en.wikipedia.org/wiki/"
	inputStart = template + inputStart
	inputEnd = template + inputEnd
	
	listSolution := search(-1,inputStart,inputEnd,[][]string{})
	for i := 0; i < len(listSolution); i++ {
		fmt.Println(listSolution[i])
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Program took %s to complete.", duration)


}
