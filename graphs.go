package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	edges := uint8(3)
	graphs := GenerateGraphs(edges)
	fmt.Println("len", edges, len(graphs))
	for _, g := range graphs {
		fmt.Println(g)
	}
}

type Graph []Node
type Node int16

func (g Graph) Copy() Graph {
	h := make(Graph, len(g))
	copy(h, g)
	return h
}

func (g Graph) AddEdge(i, j uint8) {
	g[i] |= 1 << j
	g[j] |= 1 << i
}

func (g Graph) HasEdge(i, j uint8) bool {
	return g[i]&(1<<j) != 0
}

func GenerateGraphs(nodes uint8) []Graph {
	graphs := []Graph{make(Graph, nodes)}
	for i := uint8(0); i < nodes; i++ {
		for j := uint8(i); j < nodes; j++ {
			l := len(graphs)
			for k := 0; k < l; k++ {
				g := graphs[k].Copy()
				g.AddEdge(i, j)
				graphs = append(graphs, g)
			}
		}
	}
	return graphs
}

func (g Graph) AdjacencyString() string {
	var b strings.Builder
	nodes := len(g)
	for i := 0; i < nodes; i++ {
		b.WriteString(strconv.FormatInt(int64(i), 10))
		b.WriteByte(':')
		for j := 0; j < nodes; j++ {
			if g.HasEdge(uint8(i), uint8(j)) {
				b.WriteByte(' ')
				b.WriteString(strconv.FormatInt(int64(j), 10))
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func (g Graph) String() string {
	var b strings.Builder
	nodes := len(g)
	first := true
	b.WriteByte('[')
	for i := 0; i < nodes; i++ {
		if g[i] == 0 {
			continue
		}
		for j := 0; j < nodes; j++ {
			if g.HasEdge(uint8(i), uint8(j)) {
				if !first {
					b.WriteByte(' ')
				}
				first = false
				b.WriteByte('(')
				b.WriteString(strconv.FormatInt(int64(i), 10))
				b.WriteByte(' ')
				b.WriteString(strconv.FormatInt(int64(j), 10))
				b.WriteByte(')')
			}
		}
	}
	b.WriteByte(']')
	return b.String()
}
