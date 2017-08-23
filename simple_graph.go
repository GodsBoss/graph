package graph

import (
	"fmt"
)

// SimpleGraph is a unidirectional non-weighted graph without multiple edges or loops.
type SimpleGraph struct {
	nodes                 NodeSet
	connectedNodesPerNode map[Node]NodeSet
	edges                 Edges
}

// NewSimpleGraph creates a new simple graph.
func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		nodes: make(NodeSet),
		connectedNodesPerNode: map[Node]NodeSet{},
		edges: NewEdges(),
	}
}

// AddNode adds a node to the graph. Signals an error if that node was already
// contained in the graph.
func (graph *SimpleGraph) AddNode(node Node) error {
	err := graph.nodes.Add(node)
	if err != nil {
		return err
	}
	graph.connectedNodesPerNode[node] = NewNodeSet()
	return nil
}

// MergeNodeSet merges all nodes of the node set into the graph. It is no error
// if nodes already included are added.
func (graph *SimpleGraph) MergeNodeSet(nodes NodeSet) {
	for node := range nodes {
		_ = graph.AddNode(node)
	}
}

// ContainsNode checks wether graph contains the node.
func (graph *SimpleGraph) ContainsNode(node Node) bool {
	return graph.nodes.Contains(node)
}

// RemoveNode removes node from the graph. Returns an error if the graph did
// not contain that node.
func (graph *SimpleGraph) RemoveNode(node Node) error {
	err := graph.nodes.Remove(node)
	if err != nil {
		return err
	}
	for edgeNode := range graph.connectedNodesPerNode[node] {
		graph.connectedNodesPerNode[edgeNode].Remove(node)
	}
	delete(graph.connectedNodesPerNode, node)
	return nil
}

// Nodes exposes the graph's nodes as a NodeSet.
func (graph *SimpleGraph) Nodes() NodeSet {
	return graph.nodes
}

// Connect creates an edge between the from and to node. Returns an error if
// the nodes are already connected or either node is not within the graph.
func (graph *SimpleGraph) Connect(from Node, to Node) error {
	connected, err := graph.Connected(from, to)
	if err != nil {
		return err
	}
	if connected {
		return fmt.Errorf("Nodes already connected")
	}
	graph.connectedNodesPerNode[from].Add(to)
	graph.connectedNodesPerNode[to].Add(from)
	graph.edges.Append(NewEdge(from, to))
	return nil
}

// Connected returns wether two nodes of the graph are connected. If one or both
// nodes are not part of the graph, an error is returned.
func (graph *SimpleGraph) Connected(from Node, to Node) (bool, error) {
	err := graph.failIfNotBothNodesInGraph(from, to)
	if err != nil {
		return false, err
	}
	if graph.connectedNodesPerNode[from].Contains(to) {
		return true, nil
	}
	return false, nil
}

// Disconnect removes and edge between two nodes. Returns an error if one or both
// nodes are not contained in the graph or if the nodes are not connected.
func (graph *SimpleGraph) Disconnect(from Node, to Node) error {
	connected, err := graph.Connected(from, to)
	if err != nil {
		return err
	}
	if !connected {
		return fmt.Errorf("Nodes are not connected")
	}
	graph.connectedNodesPerNode[from].Remove(to)
	graph.connectedNodesPerNode[to].Remove(from)
	graph.edges = graph.edges.Without(NewEdge(from, to))
	return nil
}

func (graph *SimpleGraph) failIfNotBothNodesInGraph(from Node, to Node) error {
	fromOK := graph.ContainsNode(from)
	toOK := graph.ContainsNode(to)
	nodesNotContained := map[string]Node{}
	if !fromOK {
		nodesNotContained["from"] = from
	}
	if !toOK {
		nodesNotContained["to"] = to
	}
	if len(nodesNotContained) > 0 {
		return nodesNotContainedError(nodesNotContained)
	}
	return nil
}

// Size exposes the size of the graph.
func (graph *SimpleGraph) Size() int {
	return graph.nodes.Size()
}

// Empty returns wether a graph is empty.
func (graph *SimpleGraph) Empty() bool {
	return Empty(graph)
}

// Edges returns the graph's edges. This list must not be changed.
func (graph *SimpleGraph) Edges() Edges {
	return graph.edges
}

// NodeEdges returns all edges for the given node. Signals an error if the node
// does not belong to the graph.
func (graph *SimpleGraph) NodeEdges(node Node) (Edges, error) {
	if !graph.ContainsNode(node) {
		return nil, fmt.Errorf("Node not contained in graph")
	}
	edges := NewEdges()
	for toNode := range graph.connectedNodesPerNode[node] {
		edges.Append(NewEdge(node, toNode))
	}
	return edges, nil
}
