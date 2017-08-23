package graph_test

import (
	"github.com/GodsBoss/graph"

	"testing"
)

func TestNodeSetDoesNotContainUnknownNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()

	if set.Contains(node) {
		t.Errorf("Expected node set %+v not to contain %+v", set, node)
	}
}

func TestNodeSetDoesContainAddedNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()
	set.Add(node)

	if !set.Contains(node) {
		t.Errorf("Expected node set %+v to contain %+v", set, node)
	}
}

func TestNodeSetDoesNotContainRemovedNode(t *testing.T) {
	set := graph.NewNodeSet()
	node := graph.NewNode()
	set.Add(node)
	set.Remove(node)

	if set.Contains(node) {
		t.Errorf("Expected node set %+v not to contain %+v", set, node)
	}
}

func TestMergingNoNodeSetsCreatesEmptyNodeSet(t *testing.T) {
	set := graph.MergeNodeSets()

	if !set.Empty() {
		t.Errorf("Expected node set %+v to be empty", set)
	}
}

func TestMergedNodeSetContainsNodesOfAllSets(t *testing.T) {
	one, two := twoNodes()
	set1 := graph.NewNodeSet()
	set1.Add(one)
	set2 := graph.NewNodeSet()
	set2.Add(two)

	set := graph.MergeNodeSets(set1, set2)

	if !set.Contains(one) {
		t.Errorf("Expected %+v to contain %+v", set, one)
	}
	if !set.Contains(two) {
		t.Errorf("Expected %+v to contain %+v", set, two)
	}
}

func TestNodeSetDiffContainsNodesNotRemoved(t *testing.T) {
	one, two := twoNodes()
	set := graph.NewNodeSet()
	set.Add(one)
	set.Add(two)

	diffSet := graph.NewNodeSet()
	diffSet.Add(one)

	diffedSet := graph.DiffNodeSet(set, diffSet)

	if !diffedSet.Contains(two) {
		t.Errorf("Expected diffed set to contain %+v", two)
	}
	if diffedSet.Contains(one) {
		t.Errorf("Expected diffed set not to contain %+v", one)
	}
}

func TestNodeExposesItselfAsANodeSet(t *testing.T) {
	node := graph.NewNode()
	set := node.Nodes()

	if !set.Contains(node) {
		t.Errorf("Expected node set %+v to include %+v", set, node)
	}
}

func TestNewNodeListIsEmpty(t *testing.T) {
	list := graph.NewNodeList()

	if !list.Empty() {
		t.Errorf("Expected list %+v to be empty, but it wasn't", list)
	}
}
