package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	graphs := GenerateGraphs(3)
	fmt.Println("len", 3, len(graphs))
	for _, g := range graphs {
		fmt.Println(g)
	}
}

type Graph []Node
type Node []int

func (g Graph) AddEdge(i, j int) Graph {
	h := make(Graph, len(g))
	copy(h, g)
	h[i] = append(h[i], j)
	if len(h[j]) == 0 || h[j][len(h[j])-1] != i {
		h[j] = append(h[j], i)
	}
	return h
}

func GenerateGraphs(nodes int) []Graph {
	graphs := []Graph{make(Graph, nodes)}
	for i := 0; i < nodes; i++ {
		for j := i; j < nodes; j++ {
			l := len(graphs)
			for g := 0; g < l; g++ {
				graphs = append(graphs, graphs[g].AddEdge(i, j))
			}
		}
	}
	return graphs
}

func (g Graph) String() string {
	var b strings.Builder
	for i, node := range g {
		b.WriteString(strconv.FormatInt(int64(i), 10))
		b.WriteByte(':')
		for _, edge := range node {
			b.WriteByte(' ')
			b.WriteString(strconv.FormatInt(int64(edge), 10))
		}
		b.WriteByte('\n')
	}
	return b.String()
}
