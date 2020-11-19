package digraph_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
)

func ExampleIsNodesNotContainedError() {
	g := digraph.New()

	n1 := g.NewNode()
	n2 := g.NewNode()

	g.Remove(n1)
	g.Remove(n2)

	err := g.Remove(n1)
	fmt.Println(err)
	ok, nodes := digraph.IsNodesNotContainedError(err)
	if ok {
		fmt.Printf("there is %d node not contained in the graph\n", len(nodes))
	}

	err = g.Connect(n1, n2)
	fmt.Println(err)

	ok, _ = digraph.IsNodesNotContainedError(fmt.Errorf("some random error"))
	if !ok {
		fmt.Println("some random error did not provide any nodes")
	}

	// Output:
	// node not contained in graph
	// there is 1 node not contained in the graph
	// nodes not contained in graph
	// some random error did not provide any nodes
}
