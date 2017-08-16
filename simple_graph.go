package graph

// SimpleGraph is a unidirectional non-weighted graph.
type SimpleGraph struct {
	nodes NodeSet
}

// NewSimpleGraph creates a new simple graph.
func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		nodes: make(NodeSet),
	}
}

// AddNode adds a node to the graph. Signals an error if that node was already
// contained in the graph.
func (graph *SimpleGraph) AddNode(node Node) error {
	return graph.nodes.Add(node)
}

// ContainsNode checks wether graph contains the node.
func (graph *SimpleGraph) ContainsNode(node Node) bool {
	return graph.nodes.Contains(node)
}

// RemoveNode removes node from the graph. Returns an error if the graph did
// not contain that node.
func (graph *SimpleGraph) RemoveNode(node Node) error {
	return graph.nodes.Remove(node)
}
