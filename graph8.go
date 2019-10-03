package graph

import "math/bits"

// Graph8 is a directed graph with 8 nodes with operations in constant
// time.
type Graph8 struct {
	g uint64
}

var _ Graph = (*Graph8)(nil)

// Add adds a directed edge from node i to j.
func (g *Graph8) Add(i, j uint) {
	g.g |= 1 << (i*8 + j)
}

// AddUndirected adds an undirected edge between nodes i and j.
func (g *Graph8) AddUndirected(i, j uint) {
	g.g |= 1 << (i*8 + j)
	g.g |= 1 << (i + j*8)
}

// Swap isomorphically swaps nodes i and j.
func (g *Graph8) Swap(i, j uint) {
	x := (g.g>>(i*8) ^ g.g>>(j*8)) & 0xff
	g.g ^= x<<(i*8) | x<<(j*8)
	x = (g.g>>i ^ g.g>>j) & 0x0101010101010101
	g.g ^= (x<<i | x<<j)
}

// Has returns whether an edge connects node i to j.
func (g *Graph8) Has(i, j uint) bool {
	return g.g&(1<<(i*8+j)) != 0
}

// Copy creates a copy of the graph.
func (g *Graph8) Copy() Graph {
	return &Graph8{g.g}
}

// OutDegree returns the number of edges directed from the given node.
func (g *Graph8) OutDegree(i uint) int {
	return bits.OnesCount8(uint8(g.g >> (i * 8)))
}

// InDegree returns the number of edges directed to the given node.
func (g *Graph8) InDegree(i uint) int {
	// Equivalent to bits.OnesCount64((g >> i) & m0)
	const (
		m0 = 0x0101010101010101
		m1 = 0x00ff00ff00ff00ff
		m2 = 0x0000ffff0000ffff
		m3 = 0x00000000ffffffff
	)
	d := (g.g >> i) & m0
	d = ((d >> 8) + d) & m1
	d = ((d >> 16) + d) & m2
	d = ((d >> 32) + d) & m3
	return int(d)
}

// Len returns the number of nodes in the graph.
func (g *Graph8) Len() int { return 8 }

// String formats the graph as a list of edges on a single line.
func (g *Graph8) String() string {
	return FormatList(g)
}
