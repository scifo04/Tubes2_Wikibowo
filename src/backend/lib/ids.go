package back

import (
	"fmt"
)

// Menjalankan fungsi dfs dengan kedalaman yang dibatasi
func DLS(g *Graph, src string, target string, limit int, isNameSpace bool) ([]string, bool) {
	// target dicapai
	if limit == 0 && src == target {
		return []string{src}, true
	}
	//target belum dicapai tetapi sudah mencapai batas depth
	if limit == 0 && src != target {
		return nil, false
	}

	// Jika nodes ada, maka tidak perlu di-scrap lagi
	if _, exist := g.Nodes[src]; !exist {
		children := ScrapeIds(src, isNameSpace)
		g.AddNode(src, children)
	}

	// Iterasi dari cache graph yang sudah dibuat
	for i := range g.Nodes[src].Child {
		// Jika ketemu, langsung return
		if g.Nodes[src].Child[i] == target {
			return []string{src, target}, true
		}

		// Jika belum ketemu, lakukan rekursi
		if g.Nodes[src].Child[i] != src {
			path, found := DLS(g, g.Nodes[src].Child[i], target, limit-1, isNameSpace)
			if found {
				return append([]string{src}, path...), true
			}
		}
	}

	return nil, false
}

// Algoritma IDS
func IDS(g *Graph, src string, target string, isNameSpace bool) ([]string, bool) {
	// Inisialisasi
	depth := 0
	var path []string
	found := false

	// Pemanggilan fungsi DLS dengan depth yang di - increment
	for {
		fmt.Printf("Depth : %d\n", depth)
		path, found = DLS(g, src, target, depth, isNameSpace)
		if found {
			return path, true
		}
		depth++
	}
}
