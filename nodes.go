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

// Add adds a node to this set and returns wether the node was successfully added
// or not, i.e. if that node was not already contained in the set.
func (set NodeSet) Add(node Node) bool {
	_, found := set[node]
	set[node] = true
	return !found
}

// Contains returns wether a node is contained in a set of nodes.
func (set NodeSet) Contains(node Node) bool {
	_, found := set[node]
	return found
}

// NodesExposer exposes its nodes.
type NodesExposer interface {
	// Nodes returns all nodes. The returned list must not be changed.
	Nodes() NodeSet
}
