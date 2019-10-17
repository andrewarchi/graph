package graph // import "github.com/andrewarchi/graph"

import "testing"

func TestAdd8(t *testing.T) {
	g := &Graph8{0, 3}
	g.Add(0, 0)
	g.Add(0, 2)
	g.Add(1, 0)
	g.Add(1, 1)
	g.Add(2, 1)
	g.Add(2, 2)
	h := &Graph8{0x060305, 3}
	if g.g != h.g {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestAddUndirected8(t *testing.T) {
	g := &Graph8{0, 2}
	h := &Graph8{0x000102, 2}
	g.AddUndirected(1, 0)
	if g.g != h.g {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestSwap8(t *testing.T) {
	g := &Graph8{0x060305, 3}
	h := &Graph8{0x050603, 3}
	g.Swap(0, 1)
	if g.g != h.g {
		t.Errorf("got %v, want %v", g, h)
	}
}
