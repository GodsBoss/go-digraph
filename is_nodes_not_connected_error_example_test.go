package digraph_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
)

func ExampleIsNodesNotConnectedError() {
	g := digraph.New()
	origin, destination := g.NewNode(), g.NewNode()

	names := map[digraph.Node]string{
		origin:      "origin",
		destination: "destination",
	}

	err := g.Disconnect(origin, destination)

	fmt.Println(err)

	if ok, edge := digraph.IsNodesNotConnectedError(err); ok {
		fmt.Printf("tried disconnecting %s and %s\n", names[edge.Origin()], names[edge.Destination()])
	}

	if ok, _ := digraph.IsNodesNotConnectedError(fmt.Errorf("any error")); !ok {
		fmt.Println("not a 'nodes not connected' error")
	}

	// Output:
	// nodes not connected
	// tried disconnecting origin and destination
	// not a 'nodes not connected' error
}
