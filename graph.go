package graph

import (
	"strconv"
	"strings"
)

// Graph is a directed graph.
type Graph interface {
	Add(i, j uint)
	AddUndirected(i, j uint)
	Swap(i, j uint)
	Has(i, j uint) bool
	Copy() Graph
	OutDegree(i uint) int
	InDegree(i uint) int
	Len() int
}

// Generate generates all possible directed graphs of a maximum rank
// starting with a given empty graph.
func Generate(g Graph) []Graph {
	graphs := []Graph{g}
	maxRank := g.Len()
	for i := 0; i < maxRank; i++ {
		for j := 0; j < maxRank; j++ {
			l := len(graphs)
			for k := 0; k < l; k++ {
				g2 := graphs[k].Copy()
				g2.Add(uint(i), uint(j))
				graphs = append(graphs, g2)
			}
		}
	}
	return graphs
}

// GenerateUndirected generates all possible undirected graphs of a maximum rank
// starting with a given empty graph.
func GenerateUndirected(g Graph) []Graph {
	graphs := []Graph{g}
	maxRank := g.Len()
	for i := 0; i < maxRank; i++ {
		for j := i; j < maxRank; j++ {
			l := len(graphs)
			for k := 0; k < l; k++ {
				g2 := graphs[k].Copy()
				g2.AddUndirected(uint(i), uint(j))
				graphs = append(graphs, g2)
			}
		}
	}
	return graphs
}

// FormatAdjacency formats the graph as an adjacency list on multiple lines.
func FormatAdjacency(g Graph) string {
	var b strings.Builder
	nodes := g.Len()
	for i := 0; i < nodes; i++ {
		b.WriteString(strconv.FormatInt(int64(i), 10))
		b.WriteByte(':')
		for j := 0; j < nodes; j++ {
			if g.Has(uint(i), uint(j)) {
				b.WriteByte(' ')
				b.WriteString(strconv.FormatInt(int64(j), 10))
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// FormatList formats the graph as a list of edges on a single line.
func FormatList(g Graph) string {
	var b strings.Builder
	nodes := g.Len()
	first := true
	b.WriteByte('[')
	for i := 0; i < nodes; i++ {
		for j := 0; j < nodes; j++ {
			if g.Has(uint(i), uint(j)) {
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
