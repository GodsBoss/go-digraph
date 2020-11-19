package operations

import (
	"github.com/GodsBoss/go-digraph"
)

// Duplicated duplicates a graph. It returns the duplicate graph and a mapping from
// nodes of the original graph to nodes of the duplicate.
func Duplicated(original digraph.Graph) (digraph.Graph, map[digraph.Node]digraph.Node) {
	duplicate := digraph.New()
	mapping := make(map[digraph.Node]digraph.Node)

	originalNodes := original.Nodes()
	for i := range originalNodes {
		n := duplicate.NewNode()
		mapping[originalNodes[i]] = n
	}

	edges := original.Edges()
	for i := range edges {
		duplicate.Connect(
			mapping[edges[i].Origin],
			mapping[edges[i].Destination],
		)
	}

	return duplicate, mapping
}
