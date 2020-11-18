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

	n1a := g1.NewNode()
	n1b := g1.NewNode()
	n2 := g2.NewNode()

	_ = g1.Remove(n1b)

	testcases := []struct {
		graph          digraph.Graph
		graphName      string
		node           digraph.Node
		nodeName       string
		expectedResult bool
	}{
		{g1, "g1", n1a, "n1a", true},
		{g1, "g1", n2, "n2", false},
		{g1, "g1", n1b, "n1b", false},
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

func TestRemoveNodes(t *testing.T) {
	g1 := digraph.New()
	g2 := digraph.New()

	n1a := g1.NewNode()
	n1b := g1.NewNode()
	g1.Remove(n1b)
	n2 := g2.NewNode()

	testcases := []struct {
		graph        digraph.Graph
		graphName    string
		node         digraph.Node
		nodeName     string
		expectsError bool
	}{
		{g1, "g1", n1a, "n1a", false},
		{g1, "g1", n1b, "n1b", true},
		{g1, "g1", n2, "n2", true},
	}

	for i := range testcases {
		testcase := testcases[i]
		t.Run(
			fmt.Sprintf("remove %s from %s", testcase.node, testcase.graph),
			func(t *testing.T) {
				err := testcase.graph.Remove(testcase.node)

				if testcase.expectsError && err == nil {
					t.Errorf("expected error")
				}
				if !testcase.expectsError && err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			},
		)
	}
}

func TestConnect(t *testing.T) {
	testcases := map[string]struct {
		f            func(g digraph.Graph) error
		expectsError bool
	}{
		"success": {
			f: func(g digraph.Graph) error {
				return g.Connect(g.NewNode(), g.NewNode())
			},
			expectsError: false,
		},
		"connects_twice": {
			f: func(g digraph.Graph) error {
				n1 := g.NewNode()
				n2 := g.NewNode()
				_ = g.Connect(n1, n2)
				return g.Connect(n1, n2)
			},
			expectsError: true,
		},
		"origin_not_contained": {
			f: func(g digraph.Graph) error {
				n := g.NewNode()
				g.Remove(n)

				return g.Connect(n, g.NewNode())
			},
			expectsError: true,
		},
		"destination_not_contained": {
			f: func(g digraph.Graph) error {
				n := g.NewNode()
				g.Remove(n)

				return g.Connect(g.NewNode(), n)
			},
			expectsError: true,
		},
	}

	for name := range testcases {
		testcase := testcases[name]
		t.Run(
			name,
			func(t *testing.T) {
				err := testcase.f(digraph.New())

				if testcase.expectsError && err == nil {
					t.Errorf("expected error")
				}
				if !testcase.expectsError && err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			},
		)
	}
}
