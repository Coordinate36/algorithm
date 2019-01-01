package main

import (
	"fmt"
	"os"
	"sort"
)

const INFILE = "/Users/coordinate36/program/algorithm/exp2/kruskal.txt"

type Edge struct {
	X, Y   int
	Weight int
}

type Edges []Edge

var edges []Edge
var depth []int
var root []int

func (e Edges) Len() int           { return len(e) }
func (e Edges) Less(i, j int) bool { return e[i].Weight < e[j].Weight }
func (e Edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func find(x int) int {
	if root[x] == -1 {
		return x
	}

	root[x] = find(root[x])
	return root[x]
}

func Union(x, y int) {
	fx := find(x)
	fy := find(y)

	if depth[fx] < depth[fy] {
		root[fx] = fy
	} else {
		root[fy] = fx
		if depth[fx] == depth[fy] {
			depth[fx]++
		}
	}
}

func kruskal() {
	for _, e := range edges {
		fx := find(e.X)
		fy := find(e.Y)
		if fx != fy {
			fmt.Printf("%d -> %d: %d\n", e.X, e.Y, e.Weight)
			Union(fx, fy)
		}
	}
}

func input() {
	var m, n int
	fmt.Scanf("%d%d", &n, &m)
	edges = make([]Edge, m)
	depth = make([]int, n+1)
	root = make([]int, n+1)
	for i := range root {
		root[i] = -1
	}

	for i := 0; i < m; i++ {
		fmt.Scanf("%d%d%d", &edges[i].X, &edges[i].Y, &edges[i].Weight)
	}

	fmt.Scanln()
}

func main() {

	// redirect stdin
	inFile, err := os.Open(INFILE)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	os.Stdin = inFile

	var numTests int
	fmt.Scanf("%d", &numTests)
	fmt.Scanln()

	for i := 0; i < numTests; i++ {
		input()
		sort.Sort(Edges(edges))
		fmt.Printf("Test%d\n", i)
		kruskal()
		fmt.Println()
	}
}
