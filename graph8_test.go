package graph

import "testing"

func TestGraph8Add(t *testing.T) {
	g := Graph8(0)
	g = g.Add(0, 0)
	g = g.Add(0, 2)
	g = g.Add(1, 0)
	g = g.Add(1, 1)
	g = g.Add(2, 1)
	g = g.Add(2, 2)
	h := Graph8(0x060305)
	if g != h {
		t.Errorf("got %x, want %x", g, h)
	}
}

func TestGraph8AddUndirected(t *testing.T) {
	g := Graph8(0).AddUndirected(1, 0)
	h := Graph8(0x000102)
	if g != h {
		t.Errorf("got %x, want %x", g, h)
	}
}

func TestGraph8Swap(t *testing.T) {
	g := Graph8(0x060305).Swap(0, 1)
	h := Graph8(0x050603)
	if g != h {
		t.Errorf("got %x, want %x", g, h)
	}
}
