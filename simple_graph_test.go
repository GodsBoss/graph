package graph_test

import (
	"github.com/godsboss/graph"

	"testing"
)

func TestGraphDoesNotContainUnknownNode(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()

	if gr.ContainsNode(node) {
		t.Errorf("Expected %+v not to contain %+v", gr, node)
	}
}

func TestGraphContainsAddedNode(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := gr.AddNode()

	if !gr.ContainsNode(node) {
		t.Errorf("Expected %+v to contain %+v", gr, node)
	}
}
