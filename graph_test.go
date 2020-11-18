package digraph_test

import (
	"github.com/GodsBoss/go-digraph"

	"testing"
)

func TestGraphNodes(t *testing.T) {
	g := digraph.New()

	addedNodes := []digraph.Node{
		g.NewNode(),
		g.NewNode(),
		g.NewNode(),
	}
	nodesFromGraph := g.Nodes()

	if len(nodesFromGraph) != 3 {
		t.Errorf("expected 3 nodes, not %d", len(nodesFromGraph))
	}

	for i := range addedNodes {
		found := false
		for i := range nodesFromGraph {
			if addedNodes[i] == nodesFromGraph[i] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("node with index %d was created, but cannot be found in graph's nodes", i)
		}
	}
}

func TestGraphNodesFromDifferentGraphsAreDistinct(t *testing.T) {
	if digraph.New().NewNode() == digraph.New().NewNode() {
		t.Errorf("expected nodes created in two graphs to be distinct")
	}
}
