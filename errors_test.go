package graph_test

import (
	"github.com/godsboss/graph"

	"fmt"
	"testing"
)

func TestNodesNotContainedFromErrorReturnsNoNodesForArbitraryErrors(t *testing.T) {
	nodes, _ := graph.NodesNotContainedFromError(fmt.Errorf("Some error"))

	if len(nodes) > 0 {
		t.Errorf("Expected no nodes, but got %+v", nodes)
	}
}
