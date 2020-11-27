package visit_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
)

func ExampleOnce() {

	// Create a graph with this structure and a mapping from nodes to
	// corresponding values (X is not a node, but two crossing lines):
	//
	//       1
	//      / \
	//     2   3
	//     |\ /|
	//     | X |
	//     |/ \|
	//     4   5
	//

	g := digraph.New()
	nodes := make([]digraph.Node, 5)
	for i := 0; i < 5; i++ {
		nodes[i] = g.NewNode()
	}

	values := map[digraph.Node]int{
		nodes[0]: 1,
		nodes[1]: 2,
		nodes[2]: 3,
		nodes[3]: 4,
		nodes[4]: 5,
	}

	g.Connect(nodes[0], nodes[1])
	g.Connect(nodes[0], nodes[2])

	g.Connect(nodes[1], nodes[3])
	g.Connect(nodes[1], nodes[4])

	g.Connect(nodes[2], nodes[3])
	g.Connect(nodes[2], nodes[4])

	addValues := func(sum *int) func(digraph.Node) (bool, error) {
		return func(node digraph.Node) (bool, error) {
			*sum += values[node]
			return true, nil
		}
	}

	withDuplicateVisitsSum := 0
	visit.Graph(g, nodes[0], addValues(&withDuplicateVisitsSum), visit.DepthFirst)
	fmt.Printf("Sum is %d (visiting nodes more than once).\n", withDuplicateVisitsSum)

	withoutDuplicateVisitsSum := 0
	visit.Graph(g, nodes[0], addValues(&withoutDuplicateVisitsSum), visit.Once(visit.DepthFirst))
	fmt.Printf("Sum is %d (visiting nodes at most once).\n", withoutDuplicateVisitsSum)

	// Output:
	// Sum is 24 (visiting nodes more than once).
	// Sum is 15 (visiting nodes at most once).
}
