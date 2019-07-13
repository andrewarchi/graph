package main

import (
	"reflect"
	"testing"
)

func TestAddUndirectedEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddUndirectedEdge(1, 0)
	h := Graph{2, 1, 0}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestAddDirectedEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddDirectedEdge(0, 0)
	g.AddDirectedEdge(0, 2)
	g.AddDirectedEdge(1, 0)
	g.AddDirectedEdge(1, 1)
	g.AddDirectedEdge(2, 1)
	g.AddDirectedEdge(2, 2)
	h := Graph{5, 3, 6}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}

func TestSwapNodes(t *testing.T) {
	g := Graph{5, 3, 6}
	g.SwapNodes(0, 1)
	h := Graph{3, 6, 5}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}
