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

func TestRemovingNodeNotInGraphFails(t *testing.T) {
	gr := graph.NewSimpleGraph()
	node := graph.NewNode()
	err := gr.RemoveNode(node)

	if err == nil {
		t.Errorf("Expected removing %+v from %+v to fail", node, gr)
	}
}

func TestSimpleGraphAllowsConnectingNodes(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	err := gr.Connect(from, to)

	if err != nil {
		t.Errorf("Expected connecting %+v and %+v in %+v not to return error, but got %+v", from, to, gr, err)
	}
}

func TestSimpleGraphFailsConnectingNodesIfBothAreNotContained(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	err := gr.Connect(from, to)

	if err == nil {
		t.Errorf("Expected connecting %+v and %+v in %+v to fail", from, to, gr)
	}
}

func TestSimpleGraphFailsConnectingNodesIfFromIsNotContained(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(to)
	err := gr.Connect(from, to)

	if err == nil {
		t.Errorf("Expected connecting %+v and %+v in %+v to fail", from, to, gr)
	}
}

func TestSimpleGraphFailsConnectingNodesIfToIsNotContained(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	err := gr.Connect(from, to)

	if err == nil {
		t.Errorf("Expected connecting %+v and %+v in %+v to fail", from, to, gr)
	}
}

func TestSimpleGraphFailsConnectingAlreadyConnectedNodes(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	gr.Connect(from, to)
	err := gr.Connect(from, to)

	if err == nil {
		t.Errorf("Expected connecting %+v and %+v in %+v to fail", from, to, gr)
	}
}

func twoNodes() (graph.Node, graph.Node) {
	return graph.NewNode(), graph.NewNode()
}
