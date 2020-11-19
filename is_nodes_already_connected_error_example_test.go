package digraph_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
)

func ExampleIsNodesAlreadyConnectedError() {
	g := digraph.New()
	origin, destination := g.NewNode(), g.NewNode()

	names := map[digraph.Node]string{
		origin:      "origin",
		destination: "destination",
	}

	g.Connect(origin, destination)
	err := g.Connect(origin, destination)

	fmt.Println(err)

	ok, edge := digraph.IsNodesAlreadyConnectedError(err)
	if ok {
		fmt.Println("error caused by connecting two already connected nodes")
		fmt.Printf("tried connecting from %s to %s\n", names[edge.Origin()], names[edge.Destination()])
	}

	if ok, _ := digraph.IsNodesAlreadyConnectedError(fmt.Errorf("some other error")); !ok {
		fmt.Println("not all errors were caused by connecting already connected nodes")
	}

	// Output:
	// nodes already connected
	// error caused by connecting two already connected nodes
	// tried connecting from origin to destination
	// not all errors were caused by connecting already connected nodes
}
