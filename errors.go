package graph

import (
	"fmt"
)

type nodesNotContainedInGraphError map[string]Node

func (err nodesNotContainedInGraphError) nodes() map[string]Node {
	return map[string]Node(err)
}

// Error lets nodesNotContainedInGraphError implement error.
func (err nodesNotContainedInGraphError) Error() string {
	nodeNames := []string{}
	for name := range err {
		nodeNames = append(nodeNames, name)
	}
	return fmt.Sprintf("nodes not contained in graph: %+v", nodeNames)
}

// NodesNotContainedInGraphFromError checks if the error given is an error
// signaling that nodes are not contained in a graph. If it is, the nodes are
// returned. The second return value determines if it was such an error at all.
func NodesNotContainedInGraphFromError(err error) (map[string]Node, bool) {
	nodesProvider, ok := err.(interface {
		nodes() map[string]Node
	})
	if !ok {
		return nil, false
	}
	return nodesProvider.nodes(), true
}
