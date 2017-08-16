package graph

import (
	"fmt"
)

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

// NewNodeSet creates a new, empty node set.
func NewNodeSet() NodeSet {
	return make(NodeSet)
}

// Add adds a node to this set. Returns an error if that node was already in the set.
func (set NodeSet) Add(node Node) error {
	_, found := set[node]
	if found {
		return fmt.Errorf("node %+v already contained", node)
	}
	set[node] = true
	return nil
}

// Contains returns wether a node is contained in a set of nodes.
func (set NodeSet) Contains(node Node) bool {
	_, found := set[node]
	return found
}

// Remove removes the node from this set. Returns an error if that node was
// not contained in the set.
func (set NodeSet) Remove(node Node) error {
	if !set.Contains(node) {
		return fmt.Errorf("node %+v not contained", node)
	}
	delete(set, node)
	return nil
}

// NodesExposer exposes its nodes.
type NodesExposer interface {
	// Nodes returns all nodes. The returned list must not be changed.
	Nodes() NodeSet
}
