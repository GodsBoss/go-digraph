package digraph_test

import (
	"fmt"

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

func TestGraphsCheckWetherTheyContainNodes(t *testing.T) {
	g1 := digraph.New()
	g2 := digraph.New()

	n1 := g1.NewNode()
	n2 := g2.NewNode()

	testcases := []struct {
		graph          digraph.Graph
		graphName      string
		node           digraph.Node
		nodeName       string
		expectedResult bool
	}{
		{g1, "g1", n1, "n1", true},
		{g1, "g1", n2, "n2", false},
	}

	for i := range testcases {
		testcase := testcases[i]
		t.Run(
			fmt.Sprintf("does %s contain %s", testcase.graphName, testcase.nodeName),
			func(t *testing.T) {
				actualResult := testcase.graph.Contains(testcase.node)

				if actualResult != testcase.expectedResult {
					t.Errorf("expected %t", testcase.expectedResult)
				}
			},
		)
	}
}
