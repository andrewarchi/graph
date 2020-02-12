package graph // import "github.com/andrewarchi/graph"

import (
	"strconv"
	"strings"
)

// Graph is a directed graph.
type Graph interface {
	Add(i, j uint)
	AddUndirected(i, j uint)
	Clear(i, j uint)
	Swap(i, j uint)
	Has(i, j uint) bool
	Copy() Graph
	Reverse() Graph
	OutDegree(i uint) int
	InDegree(i uint) int
	Len() int
	String() string
}

// NewGraph constructs a graph with a given number of nodes, selecting
// the type closest to the rank.
func NewGraph(rank uint) Graph {
	switch {
	case rank == 0:
		return nil
	case rank <= 8:
		return &Graph8{0, uint8(rank)}
	case rank <= 16:
		return make(Graph16, rank)
	default:
		return NewGraphN(rank)
	}
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

// GenerateUndirected generates all possible undirected graphs of a
// maximum rank starting with a given empty graph.
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

// FormatAdjacency formats the graph as an adjacency list on multiple
// lines.
func FormatAdjacency(g Graph) string {
	if g == nil {
		return "<nil>"
	}
	var b strings.Builder
	l := g.Len()
	for i := 0; i < l; i++ {
		b.WriteString(strconv.FormatInt(int64(i), 10))
		b.WriteByte(':')
		for j := 0; j < l; j++ {
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
	if g == nil {
		return "<nil>"
	}
	var b strings.Builder
	l := g.Len()
	first := true
	b.WriteByte('[')
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
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

// FormatGrid formats the graph as a grid on multiple lines.
func FormatGrid(g Graph) string {
	if g == nil {
		return "<nil>"
	}
	var b strings.Builder
	l := g.Len()
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if g.Has(uint(i), uint(j)) {
				b.WriteByte('*')
			} else {
				b.WriteByte('-')
			}
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// FormatMatrix formats a graph as a matrix on multiple lines using
// extending Unicode brackets.
func FormatMatrix(g Graph) string {
	if g == nil {
		return "<nil>"
	}
	var b strings.Builder
	l := g.Len()
	for i := 0; i < l; i++ {
		if i == 0 {
			b.WriteRune('⎡')
		} else if i == l-1 {
			b.WriteRune('⎣')
		} else {
			b.WriteRune('⎢')
		}
		for j := 0; j < l; j++ {
			if g.Has(uint(i), uint(j)) {
				b.WriteByte('1')
			} else {
				b.WriteByte('-')
			}
		}
		if i == 0 {
			b.WriteRune('⎤')
		} else if i == l-1 {
			b.WriteRune('⎦')
		} else {
			b.WriteRune('⎥')
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// FormatGridLabeled formats a graph as an ASCII grid with row and
// column labels.
func FormatGridLabeled(g Graph, labels []string) string {
	if g == nil {
		return "<nil>"
	}
	l := g.Len()
	if l != len(labels) {
		panic("graph: unequal number of edges and labels")
	}

	var maxLen int
	runeLabels := make([][]rune, l)
	for i := range labels {
		label := []rune(labels[i])
		if len(label) > maxLen {
			maxLen = len(label)
		}
		runeLabels[i] = label
	}
	sb := make([]byte, maxLen)
	for i := range sb {
		sb[i] = ' '
	}
	spaces := string(sb)

	// Print bottom aligned column labels
	var b strings.Builder
	for i := 0; i < maxLen; i++ {
		b.WriteString(spaces)
		b.WriteString("  ")
		for _, label := range runeLabels {
			r := ' '
			if d := maxLen - len(label); i >= d {
				r = label[i-d]
			}
			b.WriteByte(' ')
			b.WriteRune(r)
		}
		b.WriteByte('\n')
	}

	// Print horizontal line
	b.WriteString(spaces)
	b.WriteString("  ")
	for i := 0; i < l; i++ {
		b.WriteString("--")
	}
	b.WriteByte('\n')

	for i := 0; i < l; i++ {
		// Print right aligned row labels
		d := maxLen - len(runeLabels[i])
		b.WriteString(spaces[:d])
		b.WriteString(labels[i])
		b.WriteString(" |")

		// Print adjacency list
		for j := 0; j < l; j++ {
			c := byte('.')
			if g.Has(uint(i), uint(j)) {
				c = 'X'
			}
			b.WriteByte(' ')
			b.WriteByte(c)
		}
		b.WriteByte('\n')
	}
	return b.String()
}
