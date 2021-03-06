package graph_test

import (
	"github.com/GodsBoss/graph"

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

func TestFreshSimpleGraphIsEmpty(t *testing.T) {
	if !graph.NewSimpleGraph().Empty() {
		t.Errorf("Expected new graph to be empty, but it was not")
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

func TestSimpleGraphDetectsIfTwoNodesAreNotConnected(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	connected, err := gr.Connected(from, to)

	if connected {
		t.Errorf("Expected %+v and %+v in %+v not to be connected", from, to, gr)
	}
	if err != nil {
		t.Errorf("Expected error not to be nil, but got %+v", err)
	}
}

func TestSimpleGraphDetectsIfTwoNodesAreConnected(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	gr.Connect(from, to)
	connected, err := gr.Connected(from, to)

	if !connected {
		t.Errorf("Expected %+v and %+v in %+v to be connected", from, to, gr)
	}
	if err != nil {
		t.Errorf("Expected error not to be nil, but got %+v", err)
	}
}

func TestSimpleGraphDetectsNodesAsDisconnectedAfterDisconnectingThem(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	gr.Connect(from, to)
	gr.Disconnect(from, to)
	connected, err := gr.Connected(from, to)

	if connected {
		t.Errorf("Expected %+v and %+v in %+v not to be connected", from, to, gr)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, but got %+v", err)
	}
}

func TestRemovingNodeFromSimpleGraphRemovesAllEdges(t *testing.T) {
	gr := graph.NewSimpleGraph()
	from, to := twoNodes()
	gr.AddNode(from)
	gr.AddNode(to)
	gr.Connect(from, to)
	gr.RemoveNode(to)
	gr.AddNode(to) // So SimpleGraph.Connected() will not return an error.

	connected, err := gr.Connected(from, to)

	if connected {
		t.Errorf("Expected %+v and %+v in %+v not to be connected", from, to, gr)
	}
	if err != nil {
		t.Errorf("Expected error to be nil, but got %+v", err)
	}
}

func TestSimpleGraphExposesEdges(t *testing.T) {
	one := graph.NewNode()
	two := graph.NewNode()
	three := graph.NewNode()

	gr := graph.NewSimpleGraph()
	gr.AddNode(one)
	gr.AddNode(two)
	gr.AddNode(three)

	gr.Connect(one, two)
	gr.Connect(one, three)
	gr.Connect(two, three)
	gr.Disconnect(one, three)

	edges := gr.Edges()

	if edges.Size() != 2 {
		t.Errorf("Expected graph to contain %d edges, but got %d edges", 2, edges.Size())
	}
}

func TestSimpleGraphExposesEdgesForANode(t *testing.T) {
	one, two := twoNodes()
	three, four := twoNodes()

	gr := graph.NewSimpleGraph()
	gr.AddNode(one)
	gr.AddNode(two)
	gr.AddNode(three)
	gr.AddNode(four)

	gr.Connect(one, two)
	gr.Connect(four, one)
	gr.Connect(three, four)

	edges, err := gr.NodeEdges(one)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %+v", err)
	}
	if edges.Size() != 2 {
		t.Errorf("Expected number of edges to be %d, but got %d", 2, edges.Size())
	}

	for _, edge := range edges {
		if edge.From() != one && edge.To() != one {
			t.Errorf("Expected only edges with the given node, but got %+v", edge)
		}
	}

	for _, edge := range edges {
		if edge.From() == three || edge.To() == three {
			t.Errorf("Expected no edges with node %+v, but %+v contained it", three, edge)
		}
	}
}

func twoNodes() (graph.Node, graph.Node) {
	return graph.NewNode(), graph.NewNode()
}
