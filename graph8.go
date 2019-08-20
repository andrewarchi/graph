package graph

import "math/bits"

// Graph8 is a directed graph with 8 nodes.
type Graph8 uint64

// Add adds a directed edge from node i to j.
func (g Graph8) Add(i, j uint) Graph8 {
	return g | 1<<(i*8+j)
}

// AddUndirected adds an undirected edge between nodes i and j.
func (g Graph8) AddUndirected(i, j uint) Graph8 {
	return g | 1<<(i*8+j) | 1<<(i+j*8)
}

// Swap isomorphically swaps nodes i and j.
func (g Graph8) Swap(i, j uint) Graph8 {
	x := (g>>(i*8) ^ g>>(j*8)) & 0xff
	g ^= x<<(i*8) | x<<(j*8)
	x = (g>>i ^ g>>j) & 0x0101010101010101
	return g ^ (x<<i | x<<j)
}

// Has returns whether an edge connects node i to j.
func (g Graph8) Has(i, j uint) bool {
	return g&(1<<(i*8+j)) != 0
}

// OutDegree returns the number of edges leading out from the given node.
func (g Graph8) OutDegree(i uint) int {
	return bits.OnesCount8(uint8(g >> (i * 8)))
}

// InDegree returns the number of edges leading to the given node.
func (g Graph8) InDegree(i uint) int {
	// Equivalent to bits.OnesCount64((g >> i) & m0)
	const (
		m0 = 0x0101010101010101
		m1 = 0x00ff00ff00ff00ff
		m2 = 0x0000ffff0000ffff
		m3 = 0x00000000ffffffff
	)
	d := (g >> i) & m0
	d = ((d >> 8) + d) & m1
	d = ((d >> 16) + d) & m2
	d = ((d >> 32) + d) & m3
	return int(d)
}

// Len returns the number of nodes in the graph.
func (g Graph8) Len() int { return 8 }
