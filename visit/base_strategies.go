package visit

import (
	"github.com/GodsBoss/go-digraph"
)

// DepthFirst visits nodes depth-first. See Wikipedia for more details:
// https://en.wikipedia.org/wiki/Depth-first_search
// This strategy will visit the same node twice or more if the corresponding
// connections exist. It may also go into an infinite loop if the graph contains
// a cycle reachable from the start node given to Graph().
func DepthFirst(g digraph.Graph) Strategy {
	return func(current digraph.Node, rest []digraph.Node) ([]digraph.Node, error) {
		outgoing, _ := g.PointedToFrom(current)
		rest = append(outgoing, rest...)
		return rest, nil
	}
}

// BreadthFirst visits nodes breadth-first. See Wikipedia for more details:
// https://en.wikipedia.org/wiki/Breadth-first_search
// This strategy will visit the same node twice or more if the corresponding
// connections exist. It may also go into an infinite loop if the graph contains
// a cycle reachable from the start node given to Graph().
func BreadthFirst(g digraph.Graph) Strategy {
	return func(current digraph.Node, rest []digraph.Node) ([]digraph.Node, error) {
		outgoing, _ := g.PointedToFrom(current)
		rest = append(rest, outgoing...)
		return rest, nil
	}
}
