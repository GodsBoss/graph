package graph

// Edge represents a connection between two nodes.
type Edge struct {
	from Node
	to   Node
}

// From returns the edge's from node.
func (edge Edge) From() Node {
	return edge.from
}

// To returns the edge's to node.
func (edge Edge) To() Node {
	return edge.to
}

// Nodes returns the edge's two nodes as a NodeSet. Lets Edge implement
// NodeExposer.
func (edge Edge) Nodes() NodeSet {
	set := NewNodeSet()
	set.Add(edge.From())
	set.Add(edge.To())
	return set
}
