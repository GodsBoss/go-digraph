package kosaraju_test

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/algorithms/kosaraju"

	"testing"
)

func TestStronglyConnectedComponents(t *testing.T) {
	dg := digraph.New()

	a, b, c, d := dg.NewNode(), dg.NewNode(), dg.NewNode(), dg.NewNode()
	e, f, g, h := dg.NewNode(), dg.NewNode(), dg.NewNode(), dg.NewNode()
	x := dg.NewNode()

	// Strongly connected components

	dg.Connect(a, b)
	dg.Connect(b, e)
	dg.Connect(e, a)

	dg.Connect(f, g)
	dg.Connect(g, f)

	dg.Connect(c, d)
	dg.Connect(d, c)
	dg.Connect(d, h)
	dg.Connect(h, d)
	dg.Connect(x, h)
	dg.Connect(h, x)

	// Some additional connections

	dg.Connect(b, c)
	dg.Connect(b, f)
	dg.Connect(e, f)
	dg.Connect(c, g)
	dg.Connect(h, g)

	components := kosaraju.StronglyConnectedComponents(dg)

	if len(components) != 3 {
		t.Fatalf("expected 3 strongly connected components, but got %d", len(components))
	}

	expectedComponents := make(map[int]struct{})

	for i := range components {
		componentAsMap := nodeSliceToMap(components[i])
		expectedComponents[len(componentAsMap)] = struct{}{}

		switch len(componentAsMap) {
		case 2:
			assertNodeInMap(t, "f", f, "2", componentAsMap)
			assertNodeInMap(t, "g", g, "2", componentAsMap)
		case 3:
			assertNodeInMap(t, "a", a, "3", componentAsMap)
			assertNodeInMap(t, "b", b, "3", componentAsMap)
			assertNodeInMap(t, "e", e, "3", componentAsMap)
		case 4:
			assertNodeInMap(t, "c", c, "4", componentAsMap)
			assertNodeInMap(t, "d", d, "4", componentAsMap)
			assertNodeInMap(t, "h", h, "4", componentAsMap)
			assertNodeInMap(t, "x", x, "4", componentAsMap)
		default:
			t.Errorf("unexpected component size %d", len(componentAsMap))
		}
	}

	for i := 2; i <= 4; i++ {
		if _, ok := expectedComponents[i]; !ok {
			t.Errorf("expected component with size %d", i)
		}
	}
}

func assertNodeInMap(t *testing.T, nodeName string, node digraph.Node, mapName string, m map[digraph.Node]struct{}) {
	if _, ok := m[node]; !ok {
		t.Errorf("expected node %s to be in map %s", nodeName, mapName)
	}
}

func nodeSliceToMap(nodes []digraph.Node) map[digraph.Node]struct{} {
	m := make(map[digraph.Node]struct{})
	for i := range nodes {
		m[nodes[i]] = struct{}{}
	}
	return m
}
