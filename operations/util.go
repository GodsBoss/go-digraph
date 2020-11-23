package operations

import (
	"github.com/GodsBoss/go-digraph"
)

// InvertedNodeMapping returns an inverted node mapping. The map passed to this
// function remains unchanged.
// If the input map contains duplicate values, is is undefined which of them will
// become key in the return value. For maps m containing every value just once,
// InvertedNodeMapping(InvertedNodeMapping(m)) == m
// holds true. In general, for an arbitrary map m,
// len(InvertedNodeMapping(m)) <= len(m)
// holds true.
func InvertedNodeMapping(mapping map[digraph.Node]digraph.Node) map[digraph.Node]digraph.Node {
	result := make(map[digraph.Node]digraph.Node)
	for key := range mapping {
		result[mapping[key]] = key
	}
	return result
}
