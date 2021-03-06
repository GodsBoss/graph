package graph

// Edge represents a connection between two nodes.
type Edge struct {
	from Node
	to   Node
}

// NewEdge creates a new edge pointing from a node to another.
func NewEdge(from Node, to Node) *Edge {
	return &Edge{
		from: from,
		to:   to,
	}
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

// Edges is a list of edges.
type Edges []*Edge

// NewEdges creates a new, empty list of edges.
func NewEdges() Edges {
	return make(Edges, 0)
}

// Append appends edges to this list of edges.
func (edges *Edges) Append(additionalEdges ...*Edge) {
	*edges = append(*edges, additionalEdges...)
}

// Without returns a list without all occurrences of the edge.
func (edges *Edges) Without(edge *Edge) Edges {
	result := NewEdges()
	for _, edgeInList := range *edges {
		if (edge.From() != edgeInList.From() || edge.To() != edgeInList.To()) &&
			(edge.From() != edgeInList.To() || edge.To() != edgeInList.From()) {
			result.Append(edgeInList)
		}
	}
	return result
}

// Size exposes the number of edges.
func (edges Edges) Size() int {
	return len(edges)
}

// Empty checks wether edges is empty.
func (edges Edges) Empty() bool {
	return Empty(edges)
}
