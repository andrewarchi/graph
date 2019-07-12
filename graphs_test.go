package main

import (
	"reflect"
	"testing"
)

func TestAddEdge(t *testing.T) {
	g := make(Graph, 3)
	g.AddEdge(1, 1)
	h := Graph{0, 2, 0}
	if !reflect.DeepEqual(g, h) {
		t.Errorf("got %v, want %v", g, h)
	}
}
