package graph_test

import (
	"github.com/godsboss/graph"

	"testing"
)

func TestEdgeExposesNodesAsASet(t *testing.T) {
	from, to := twoNodes()
	edge := graph.NewEdge(from, to)
	nodes := edge.Nodes()

	if !nodes.Contains(from) {
		t.Errorf("Nodes %+v exposes by edge %+v does not contain %+v", nodes, edge, from)
	}
	if !nodes.Contains(to) {
		t.Errorf("Nodes %+v exposes by edge %+v does not contain %+v", nodes, edge, to)
	}
}

func TestEdgesCanBeAppendedToEdgeList(t *testing.T) {
	edges := graph.NewEdges()
	n11, n12 := twoNodes()
	n21, n22 := twoNodes()

	edge1 := graph.NewEdge(n11, n12)
	edge2 := graph.NewEdge(n21, n22)

	edges.Append(edge1, edge2)

	if edges.Size() != 2 {
		t.Errorf("Expected size of %+v to be %d, but was %d", edges, 2, edges.Size())
	}
}
