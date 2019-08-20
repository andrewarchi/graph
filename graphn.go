package graph

type GraphN struct {
	v    []uint64
	rank uint
}

func NewGraphN(rank uint) *GraphN {
	return &GraphN{make([]uint64, rank*rank/64), rank}
}

func (g GraphN) Add(i, j uint) {
	p := i*g.rank + j
	g.v[p/64] |= 1 << (p % 64)
}

func (g GraphN) AddUndirected(i, j uint) {
	g.Add(i, j)
	g.Add(j, i)
}

func (g GraphN) Swap(i, j uint) {

}

func (g GraphN) Has(i, j uint) bool {
	p := i*g.rank + j
	return g.v[p/64]&(1<<(p%64)) != 0
}

func (g GraphN) OutDegree(i uint) int {
	for p := i * g.rank; p < (i+1)*g.rank; p++ {
		g.v[p/64]
	}
}
