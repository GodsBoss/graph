package graph_test

import (
	"github.com/godsboss/graph"

	"testing"
)

func TestSimpleGraphDoesNotContainUnknownNode(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()

	if gr.ContainsNode(node) {
		t.Errorf("Expected %+v not to contain %+v", gr, node)
	}
}

func TestSimpleGraphContainsAddedNode(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()
	gr.AddNode(node)

	if !gr.ContainsNode(node) {
		t.Errorf("Expected %+v to contain %+v", gr, node)
	}
}

func TestCannotNodesAlreadyAddedToSimpleGraph(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()
	gr.AddNode(node)
	err := gr.AddNode(node)

	if err == nil {
		t.Errorf("Expected adding %+v to %+v to return an error", node, gr)
	}
}

func TestSimpleGraphDoesNotContainRemovedNodes(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()
	gr.AddNode(node)
	gr.RemoveNode(node)

	if gr.ContainsNode(node) {
		t.Errorf("Expected %+v not to contain %+v", gr, node)
	}
}
