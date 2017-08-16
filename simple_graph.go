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

// AddNode adds a node to the graph and returns that node.
func (graph *SimpleGraph) AddNode() Node {
	node := NewNode()
	graph.nodes.Add(node)
	return node
}

// ContainsNode checks wether graph contains the node.
func (graph *SimpleGraph) ContainsNode(node Node) bool {
	return graph.nodes.Contains(node)
}
