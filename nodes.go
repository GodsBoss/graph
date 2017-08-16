package graph

// Node is the node of a graph.
type Node *int

// NewNode creates a new node.
func NewNode() Node {
	var i int
	return Node(&i)
}

// Nodes is list of nodes. Duplicates are allowed.
type Nodes []Node

// NodeSet is a non-ordered set of nodes. Nodes may appear only once in a node set.
type NodeSet map[Node]bool

// NodesExposer exposes its nodes.
type NodesExposer interface {
	// Nodes returns all nodes. The returned list must not be changed.
	Nodes() NodeSet
}
