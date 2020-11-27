package visit_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
)

func ExampleGraph() {

	// Create a graph with this structure and a mapping from nodes to
	// corresponding names:
	//
	//          a
	//         / \
	//        /   \
	//       b     b
	//      / \   / \
	//     c   c c   c
	//

	g := digraph.New()

	nodes := make([]digraph.Node, 7)
	for i := range nodes {
		nodes[i] = g.NewNode()
	}

	g.Connect(nodes[0], nodes[1])
	g.Connect(nodes[0], nodes[2])
	g.Connect(nodes[1], nodes[3])
	g.Connect(nodes[1], nodes[4])
	g.Connect(nodes[2], nodes[5])
	g.Connect(nodes[2], nodes[6])

	names := map[digraph.Node]string{
		nodes[0]: "a",
		nodes[1]: "b",
		nodes[2]: "b",
		nodes[3]: "c",
		nodes[4]: "c",
		nodes[5]: "c",
		nodes[6]: "c",
	}

	appendToResult := func(result *string) func(digraph.Node) (bool, error) {
		return func(n digraph.Node) (bool, error) {
			*result += names[n]
			return true, nil
		}
	}

	depthFirstResult := ""
	visit.Graph(g, nodes[0], appendToResult(&depthFirstResult), visit.DepthFirst)
	fmt.Println(depthFirstResult)

	breadthFirstResult := ""
	visit.Graph(g, nodes[0], appendToResult(&breadthFirstResult), visit.BreadthFirst)
	fmt.Println(breadthFirstResult)

	// Output:
	// abccbcc
	// abbcccc
}
