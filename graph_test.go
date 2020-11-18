package digraph_test

import (
	"github.com/GodsBoss/go-digraph"

	"fmt"
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
		for j := range nodesFromGraph {
			if addedNodes[i] == nodesFromGraph[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("node %+v (index %d) was created from %+v, but cannot be found in %+v", addedNodes[i], i, g, nodesFromGraph)
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
	ne1 := g1.NewNode()
	ne2 := g1.NewNode()
	g1.Connect(ne1, ne2)
	ned1 := g2.NewNode()
	ned2 := g2.NewNode()
	g2.Connect(ned1, ned2)
	g2.Disconnect(ned1, ned2)

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
		{g1, "g1", ne1, "ne1", true},
		{g1, "g1", ne2, "ne2", true},
		{g2, "g2", ned1, "ned1", false},
		{g2, "g2", ned2, "ned2", false},
	}

	for i := range testcases {
		testcase := testcases[i]
		t.Run(
			fmt.Sprintf("remove %s from %s", testcase.nodeName, testcase.graphName),
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

func TestEdges(t *testing.T) {
	g := digraph.New()
	n1 := g.NewNode()
	n2 := g.NewNode()
	n3 := g.NewNode()

	g.Connect(n1, n2)
	g.Connect(n3, n3)
	g.Connect(n1, n3)
	g.Connect(n3, n2)
	g.Disconnect(n3, n2)

	expectedEdges := []digraph.Edge{
		{
			Origin:      n1,
			Destination: n2,
		},
		{
			Origin:      n3,
			Destination: n3,
		},
		{
			Origin:      n1,
			Destination: n3,
		},
	}

	edgesFromGraph := g.Edges()

	if len(edgesFromGraph) != 3 {
		t.Errorf("expected 3 edges from graph, got %d", len(edgesFromGraph))
	}

	for i := range expectedEdges {
		found := false
		for j := range edgesFromGraph {
			if expectedEdges[i] == edgesFromGraph[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("edge %+v not found in %+v", expectedEdges[i], edgesFromGraph)
		}
	}
}

func TestDisconnectEdgeErrors(t *testing.T) {
	testcases := map[string]struct {
		f func(g digraph.Graph) (origin, destination digraph.Node)
	}{
		"origin_not_contained": {
			f: func(g digraph.Graph) (origin, destination digraph.Node) {
				origin = g.NewNode()
				g.Remove(origin)
				return origin, g.NewNode()
			},
		},
		"destination_not_contained": {
			f: func(g digraph.Graph) (origin, destination digraph.Node) {
				destination = g.NewNode()
				g.Remove(destination)
				return g.NewNode(), destination
			},
		},
		"not_connected": {
			f: func(g digraph.Graph) (origin, destination digraph.Node) {
				return g.NewNode(), g.NewNode()
			},
		},
		"connection_removed": {
			f: func(g digraph.Graph) (origin, destination digraph.Node) {
				origin = g.NewNode()
				destination = g.NewNode()
				g.Connect(origin, destination)
				g.Disconnect(origin, destination)
				return origin, destination
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]
		t.Run(
			name,
			func(t *testing.T) {
				g := digraph.New()
				origin, destination := testcase.f(g)
				err := g.Disconnect(origin, destination)

				if err == nil {
					t.Errorf("expected error")
				}
			},
		)
	}
}

func TestPointX(t *testing.T) {
	g := digraph.New()

	nodeWithoutConnections := g.NewNode()
	nodeConnectedToItself := g.NewNode()
	nodeMututallyConnected1 := g.NewNode()
	nodeMututallyConnected2 := g.NewNode()
	nodeOrigin := g.NewNode()
	nodeDestination1 := g.NewNode()
	nodeDestination2 := g.NewNode()
	nodeOrigin1 := g.NewNode()
	nodeOrigin2 := g.NewNode()
	nodeDestination := g.NewNode()
	nodeStart := g.NewNode()
	nodeMiddle := g.NewNode()
	nodeEnd := g.NewNode()

	mustConnect := func(origin, destination digraph.Node) {
		err := g.Connect(origin, destination)
		if err != nil {
			t.Fatalf("could not connect nodes: %+v", err)
		}
	}

	mustConnect(nodeConnectedToItself, nodeConnectedToItself)
	mustConnect(nodeMututallyConnected1, nodeMututallyConnected2)
	mustConnect(nodeMututallyConnected2, nodeMututallyConnected1)
	mustConnect(nodeOrigin, nodeDestination1)
	mustConnect(nodeOrigin, nodeDestination2)
	mustConnect(nodeOrigin1, nodeDestination)
	mustConnect(nodeOrigin2, nodeDestination)
	mustConnect(nodeStart, nodeMiddle)
	mustConnect(nodeMiddle, nodeEnd)

	testcases := map[string]struct {
		node                 digraph.Node
		expectedOrigins      []digraph.Node
		expectedDestinations []digraph.Node
	}{
		"nodeWithoutConnections": {
			node:                 nodeWithoutConnections,
			expectedOrigins:      make([]digraph.Node, 0),
			expectedDestinations: make([]digraph.Node, 0),
		},
		"nodeConnectedToItself": {
			node: nodeConnectedToItself,
			expectedOrigins: []digraph.Node{
				nodeConnectedToItself,
			},
			expectedDestinations: []digraph.Node{
				nodeConnectedToItself,
			},
		},
		"nodeMututallyConnected": {
			node: nodeMututallyConnected2,
			expectedOrigins: []digraph.Node{
				nodeMututallyConnected1,
			},
			expectedDestinations: []digraph.Node{
				nodeMututallyConnected1,
			},
		},
		"nodeDestination1": {
			node: nodeDestination1,
			expectedOrigins: []digraph.Node{
				nodeOrigin,
			},
			expectedDestinations: make([]digraph.Node, 0),
		},
		"nodeOrigin1": {
			node:            nodeOrigin1,
			expectedOrigins: make([]digraph.Node, 0),
			expectedDestinations: []digraph.Node{
				nodeDestination,
			},
		},
		"nodeDestination": {
			node: nodeDestination,
			expectedOrigins: []digraph.Node{
				nodeOrigin1,
				nodeOrigin2,
			},
			expectedDestinations: make([]digraph.Node, 0),
		},
		"nodeOrigin": {
			node:            nodeOrigin,
			expectedOrigins: make([]digraph.Node, 0),
			expectedDestinations: []digraph.Node{
				nodeDestination1,
				nodeDestination2,
			},
		},
		"nodeMiddle": {
			node: nodeMiddle,
			expectedOrigins: []digraph.Node{
				nodeStart,
			},
			expectedDestinations: []digraph.Node{
				nodeEnd,
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]
		t.Run(
			name,
			func(t *testing.T) {
				t.Run(
					"origins",
					func(t *testing.T) {
						origins, err := g.PointingTo(testcase.node)
						if err != nil {
							t.Fatalf("failed getting origins for node: %+v", err)
						}
						assertNodesContainNodes(t, testcase.expectedOrigins, origins)
					},
				)
				t.Run(
					"destinations",
					func(t *testing.T) {
						destinations, err := g.PointedToFrom(testcase.node)
						if err != nil {
							t.Fatalf("failed getting destinations for node: %+v", err)
						}
						assertNodesContainNodes(t, testcase.expectedDestinations, destinations)
					},
				)
			},
		)
	}
}

func assertNodesContainNodes(t *testing.T, expected []digraph.Node, actual []digraph.Node) {
	for i := range expected {
		found := false
		for j := range actual {
			if expected[i] == actual[j] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("did not find node %+v (index %d) in %+v", expected[i], i, actual)
		}
	}
}

func TestPointXNodeNotContained(t *testing.T) {
	g := digraph.New()
	n := g.NewNode()
	g.Remove(n)

	t.Run(
		"PointingTo",
		func(t *testing.T) {
			origins, err := g.PointingTo(n)
			if err == nil {
				t.Errorf("expected error")
			}
			if len(origins) > 0 {
				t.Errorf("expected no nodes, got %+v", origins)
			}
		},
	)

	t.Run(
		"PointedToFrom",
		func(t *testing.T) {
			destinations, err := g.PointedToFrom(n)
			if err == nil {
				t.Errorf("expected error")
			}
			if len(destinations) > 0 {
				t.Errorf("expected no nodes, got %+v", destinations)
			}
		},
	)
}
