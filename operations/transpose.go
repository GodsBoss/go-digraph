package operations

import (
	"github.com/GodsBoss/go-digraph"
)

// Transposed (could also be called Converse or Reverse) takes a graph and returns
// a graph with basically the same nodes, but all edges reversed. The original
// graph remains unchanged.
// Returns a mapping from nodes of the original graph to the transposed graph
// as well.
func Transposed(original digraph.Graph) (digraph.Graph, map[digraph.Node]digraph.Node) {
	transposed := digraph.New()
	mapping := make(map[digraph.Node]digraph.Node)

	originalNodes := original.Nodes()
	for i := range originalNodes {
		n := transposed.NewNode()
		mapping[originalNodes[i]] = n
	}

	edges := original.Edges()
	for i := range edges {
		transposed.Connect(
			mapping[edges[i].Destination],
			mapping[edges[i].Origin],
		)
	}

	return transposed, mapping
}
