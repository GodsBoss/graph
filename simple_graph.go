package graph

import (
	"fmt"
)

// SimpleGraph is a unidirectional non-weighted graph.
type SimpleGraph struct {
	nodes NodeSet
	edges map[Node]NodeSet
}

// NewSimpleGraph creates a new simple graph.
func NewSimpleGraph() *SimpleGraph {
	return &SimpleGraph{
		nodes: make(NodeSet),
		edges: map[Node]NodeSet{},
	}
}

// AddNode adds a node to the graph. Signals an error if that node was already
// contained in the graph.
func (graph *SimpleGraph) AddNode(node Node) error {
	err := graph.nodes.Add(node)
	if err != nil {
		return err
	}
	graph.edges[node] = NewNodeSet()
	return nil
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
	graph.edges[from].Add(to)
	graph.edges[to].Add(from)
	return nil
}

// Connected returns wether two nodes of the graph are connected. If one or both
// nodes are not part of the graph, an error is returned.
func (graph *SimpleGraph) Connected(from Node, to Node) (bool, error) {
	err := graph.failIfNotBothNodesInGraph(from, to)
	if err != nil {
		return false, err
	}
	if graph.edges[from].Contains(to) {
		return true, nil
	}
	return false, nil
}

func (graph *SimpleGraph) failIfNotBothNodesInGraph(from Node, to Node) error {
	fromOK := graph.ContainsNode(from)
	toOK := graph.ContainsNode(to)
	if (!fromOK) && (!toOK) {
		return fmt.Errorf("Both nodes not contained in graph")
	}
	if !fromOK {
		return fmt.Errorf("from node not contained in graph")
	}
	if !toOK {
		return fmt.Errorf("to node not contained in graph")
	}
	return nil
}
