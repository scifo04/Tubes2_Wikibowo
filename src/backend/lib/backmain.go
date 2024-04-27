package back

import (
	"fmt"
	"sync"
	"time"
)

func Back_Main(inpute LinkInfo, maxConcurrency int) ([]string, int64, int, []string) {
	var engine Engine
	fmt.Println("Layer ", engine.Depth)
	inpute.LinkValue = TurnToWikipedia(inpute.LinkValue)
	inpute.FinValue = TurnToWikipedia(inpute.FinValue)
	engine = createEngine(engine, inpute.LinkValue, inpute.FinValue, 0, inpute.IsName)
	// cache = []string{}
	var listSolution []string
	var linkSolution []string
	ch := make(chan []string)
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency)
	t1 := time.Now()

	g := NewGraph()
	var cnt int = 0

	if inpute.IsOn {
		listSolution, _ = IDS(g, inpute.LinkValue, inpute.FinValue, inpute.IsName)
		for _, value := range g.Nodes {
			cnt += len(value.Child)
		}
	} else {
		go BFS(-1, inpute.LinkValue, inpute.FinValue, [][]string{}, engine, ch, &wg, sem)
		listSolution = <-ch
	}

	for i := 0; i < len(listSolution); i++ {
		listSolution[i] = TurnToTitle(listSolution[i])
	}

	for i := 0; i < len(listSolution); i++ {
		linkSolution = append(linkSolution, TurnToWikipedia(listSolution[i]))
	}

	t3 := time.Since(t1)
	t2 := time.Since(t1).Milliseconds()
	// fmt.Println(cache)

	fmt.Println("Execution Time: ", t3)
	if inpute.IsOn {
		fmt.Print(listSolution)
		fmt.Print(t2)
		fmt.Print(len(g.Nodes))
		return listSolution, t2, cnt, linkSolution
	} else {
		return listSolution, t2, len(engine.Cache), linkSolution
	}
}
