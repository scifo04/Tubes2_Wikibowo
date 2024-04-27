package back

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type Node struct {
	Child []string
}

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(value string, child []string) {
	if _, exist := g.Nodes[value]; !exist {
		newNode := &Node{
			Child: child,
		}
		g.Nodes[value] = newNode
	}
}

func getLinks(lenc string) []string {
	var links []string
	c := colly.NewCollector(
		colly.DisallowedURLFilters(
			regexp.MustCompile("Category:"),
			regexp.MustCompile("Wikipedia:"),
			regexp.MustCompile("Special:"),
			regexp.MustCompile("File:"),
			regexp.MustCompile("Help:"),
			regexp.MustCompile("Talk:"),
			regexp.MustCompile("Portal:"),
			regexp.MustCompile("Template:"),
			regexp.MustCompile("Template_talk:"),
			regexp.MustCompile("Main_Page"),
		),
		colly.URLFilters(
			regexp.MustCompile("en.wikipedia.org/wiki"),
		),
	)
	// colly.AllowedDomains("en.wikipedia.org","https://en.wikipedia.org","en.wikipedia.org/","https://en.wikipedia.org/",))

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting: ", lenc)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println("Response Code:", r.StatusCode)
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

func DLS(g *Graph, src string, target string, limit int) ([]string, bool) {
	if limit == 0 && src == target {
		return []string{src}, true
	} else if limit == 0 && src != target {
		return nil, false
	}

	if _, exist := g.Nodes[src]; !exist {
		children := getLinks(src)
		g.AddNode(src, children)
	}

	for i := range g.Nodes[src].Child {
		if g.Nodes[src].Child[i] == target {
			return []string{src, target}, true
		}

		if g.Nodes[src].Child[i] != src {
			path, found := DLS(g, g.Nodes[src].Child[i], target, limit-1)
			if found {
				return append([]string{src}, path...), true
			}
		}
	}

	return nil, false
}

func IDS(g *Graph, src string, target string) ([]string, bool) {
	depth := 0
	var path []string
	found := false

	for {
		fmt.Printf("Depth : %d\n", depth)
		path, found = DLS(g, src, target, depth)
		if found {
			return path, true
		}
		depth++
	}
}
