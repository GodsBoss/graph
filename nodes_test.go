package graph_test

import (
	"github.com/godsboss/graph"

	"testing"
)

func TestNodeSetDoesNotContainUnknownNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()

	if set.Contains(node) {
		t.Errorf("Expected node set %+v not to contain %+v", set, node)
	}
}

func TestNodeSetDoesContainAddedNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()
	set.Add(node)

	if !set.Contains(node) {
		t.Errorf("Expected node set %+v to contain %+v", set, node)
	}
}

func TestNodeSetDoesNotContainRemovedNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()
	set.Add(node)
	set.Remove(node)

	if set.Contains(node) {
		t.Errorf("Expected node set %+v not to contain %+v", set, node)
	}
}

func TestNodeExposesItselfAsANodeSet(t *testing.T) {
	node := graph.NewNode()
	set := node.Nodes()

	if !set.Contains(node) {
		t.Errorf("Expected node set %+v to include %+v", set, node)
	}
}
