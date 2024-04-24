package back

import (
	"fmt"
	"time"
	"sync"
)

func Back_Main(inpute LinkInfo, maxConcurrency int) ([]string, int64) {
	var engine Engine
	fmt.Println("Layer ",engine.Depth)
	inpute.LinkValue = TurnToWikipedia(inpute.LinkValue)
	inpute.FinValue = TurnToWikipedia(inpute.FinValue)
	engine = createEngine(engine,inpute.LinkValue,inpute.FinValue,0,inpute.IsName)
	// cache = []string{}
	var listSolution []string
	ch := make(chan []string)
	var wg sync.WaitGroup
	sem := make(chan struct{},maxConcurrency)
	t1 := time.Now()
	if (inpute.IsOn) {
		go BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{},engine,ch,&wg,sem)
		// IDS()
	} else {
		go BFS(-1,inpute.LinkValue,inpute.FinValue,[][]string{},engine,ch,&wg,sem)
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