package graph

import (
	"fmt"
)

type nodesNotContainedError map[string]Node

func (err nodesNotContainedError) nodes() map[string]Node {
	return map[string]Node(err)
}

// Error lets nodesNotContainedError implement error.
func (err nodesNotContainedError) Error() string {
	nodeNames := []string{}
	for name := range err {
		nodeNames = append(nodeNames, name)
	}
	return fmt.Sprintf("nodes not contained: %+v", nodeNames)
}

// NodesNotContainedFromError checks if the error given is an error
// signaling that nodes are not contained in a graph. If it is, the nodes are
// returned. The second return value determines if it was such an error at all.
func NodesNotContainedFromError(err error) (map[string]Node, bool) {
	nodesProvider, ok := err.(interface {
		nodes() map[string]Node
	})
	if !ok {
		return nil, false
	}
	return nodesProvider.nodes(), true
}
