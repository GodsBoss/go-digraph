package operations

import (
	"github.com/GodsBoss/go-digraph"
)

// InvertedNodeMapping returns an inverted node mapping. The map passed to this
// function remains unchanged.
func InvertedNodeMapping(mapping map[digraph.Node]digraph.Node) map[digraph.Node]digraph.Node {
	result := make(map[digraph.Node]digraph.Node)
	for key := range mapping {
		result[mapping[key]] = key
	}
	return result
}
