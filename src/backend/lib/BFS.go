package back

import (
	"fmt"
	"sync"
)

func BFS(i int, start string, end string, listScrape [][]string, eng Engine, ch chan<- []string, wg *sync.WaitGroup, sem chan struct{}) {
	defer wg.Done()
	
	var tempScrape [][]string;

	if ((i >= 0 && len(listScrape[i]) == eng.Depth+1) || i == -1) {
		eng.Depth++
		fmt.Println("Layer ",eng.Depth)
	}
	fmt.Print(len(listScrape))
	
	if (i == -1){ //listScrape masih kosong
		fmt.Println(" ",start)
		chScrape := make(chan [][]string)
		go Scrape([]string{start},eng,chScrape,wg,sem)
		tempScrape = <-chScrape
		listScrape = append(listScrape, tempScrape...)
	} else {
		fmt.Println(" ",listScrape[i])
		chScrape := make(chan[][] string)
		go Scrape(listScrape[i],eng,chScrape,wg,sem)
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
    go BFS(i+1, start, end, listScrape, eng, ch, wg, sem)
}