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
