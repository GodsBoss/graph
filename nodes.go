package graph

import (
	"fmt"
)

// Node is the node of a graph.
type Node struct {
	id *int
}

// Nodes returns itself as a NodeSet. Lets node implement NodesExposer.
func (node Node) Nodes() NodeSet {
	set := NewNodeSet()
	set.Add(node)
	return set
}

// NewNode creates a new node.
func NewNode() Node {
	var i int
	return Node{
		id: &i,
	}
}

// NodeList is list of nodes. Duplicates are allowed.
type NodeList []Node

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

// Nodes returns the node set. Lets NodeSet implement NodesExposer.
func (set NodeSet) Nodes() NodeSet {
	return set
}

// NodesExposer exposes its nodes.
type NodesExposer interface {
	// Nodes returns all nodes. The returned set must not be changed.
	Nodes() NodeSet
}
