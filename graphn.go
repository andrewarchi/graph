package graph

// GraphN is a directed graph with an arbitrary number of nodes.
type GraphN struct {
	g    []uint64
	rank uint
}

// NewGraphN constructs a graph with the given number of nodes.
func NewGraphN(rank uint) *GraphN {
	n := 1 + ((rank*rank - 1) / 64) // ceiling division
	return &GraphN{make([]uint64, n), rank}
}

// Add adds a directed edge from node i to j.
func (g *GraphN) Add(i, j uint) {
	p := i*g.rank + j
	g.g[p/64] |= 1 << (p % 64)
}

// AddUndirected adds an undirected edge between nodes i and j.
func (g *GraphN) AddUndirected(i, j uint) {
	g.Add(i, j)
	g.Add(j, i)
}

// Swap is unimplemented.
func (g *GraphN) Swap(i, j uint) {
	panic("graph: swap is unimplemented for GraphN")
}

// Has returns whether an edge connects node i to j.
func (g *GraphN) Has(i, j uint) bool {
	p := i*g.rank + j
	return g.g[p/64]&(1<<(p%64)) != 0
}

// Copy creates a copy of the graph.
func (g *GraphN) Copy() Graph {
	h := make([]uint64, len(g.g))
	copy(h, g.g)
	return &GraphN{h, g.rank}
}

// OutDegree is unimplemented.
func (g *GraphN) OutDegree(i uint) int {
	panic("graph: out degree is unimplemented for GraphN")
	// for p := i * g.rank; p < (i+1)*g.rank; p++ {
	// 	g.g[p/64]
	// }
}

// InDegree is unimplemented.
func (g *GraphN) InDegree(i uint) int {
	panic("graph: in degree is unimplemented for GraphN")
}

// Len returns the number of nodes in the graph.
func (g *GraphN) Len() int {
	return int(g.rank)
}

func (g *GraphN) String() string {
	return FormatList(g)
}
