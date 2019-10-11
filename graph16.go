package graph

import "math/bits"

// Graph16 is a directed graph with at most 16 nodes.
type Graph16 []uint16

var _ Graph = (Graph16)(nil)

// NewGraph16 constructs a graph with a given number of nodes.
func NewGraph16(rank uint) Graph16 {
	if rank > 16 {
		panic("graph: Graph16 rank out of bounds")
	}
	return make(Graph16, rank)
}

// Add adds a directed edge from node i to j.
func (g Graph16) Add(i, j uint) {
	g[i] |= 1 << j
}

// AddUndirected adds an undirected edge between nodes i and j.
func (g Graph16) AddUndirected(i, j uint) {
	g[i] |= 1 << j
	g[j] |= 1 << i
}

// Clear removes the directed edge from node i to j.
func (g Graph16) Clear(i, j uint) {
	g[i] &^= 1 << j
}

// Swap isomorphically swaps nodes i and j.
func (g Graph16) Swap(i, j uint) {
	g[i], g[j] = g[j], g[i]
	for n, node := range g {
		// Swap individual bits
		// http://graphics.stanford.edu/~seander/bithacks.html#SwappingBitsXOR
		x := (node>>i ^ node>>j) & 0x1
		g[n] = node ^ (x<<i | x<<j)
	}
}

// Has returns whether an edge connects node i to j.
func (g Graph16) Has(i, j uint) bool {
	return g[i]&(1<<j) != 0
}

// Copy creates a copy of the graph.
func (g Graph16) Copy() Graph {
	h := make(Graph16, len(g))
	copy(h, g)
	return h
}

// Reverse creates a graph with reversed edges.
func (g Graph16) Reverse() Graph {
	h := make(Graph16, len(g))
	for i, node := range g {
		for node != 0 {
			j := uint(bits.TrailingZeros16(node))
			h.Add(j, uint(i))
			node &^= 1 << j
		}
	}
	return h
}

// OutDegree returns the number of edges directed from the given node.
func (g Graph16) OutDegree(i uint) int {
	return bits.OnesCount16(uint16(g[i]))
}

// InDegree returns the number of edges directed to the given node.
func (g Graph16) InDegree(i uint) int {
	d := 0
	for j := 0; j < len(g); j++ {
		d += int((g[j] >> i) & 0x1)
	}
	return d
}

// Len returns the number of nodes in the graph.
func (g Graph16) Len() int {
	return len(g)
}

// String formats the graph as a list of edges on a single line.
func (g Graph16) String() string {
	return FormatList(g)
}
