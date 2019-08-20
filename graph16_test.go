package graph

import (
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	g := NewGraph16(3)
	g.Add(0, 0)
	g.Add(0, 2)
	g.Add(1, 0)
	g.Add(1, 1)
	g.Add(2, 1)
	g.Add(2, 2)
	h := Graph16{0x5, 0x3, 0x6}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestAddUndirected(t *testing.T) {
	g := NewGraph16(3)
	g.AddUndirected(1, 0)
	h := Graph16{0x2, 0x1, 0x0}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestSwap(t *testing.T) {
	g := Graph16{0x5, 0x3, 0x6}
	g.Swap(0, 1)
	h := Graph16{0x3, 0x6, 0x5}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}
