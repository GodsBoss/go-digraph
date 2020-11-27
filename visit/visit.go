package visit

import (
	"github.com/GodsBoss/go-digraph"
)

// Graph visits nodes of a graph according to the given strategy, with start
// being the first node visited. f is called for the currently visited node.
// After that, the current list of nodes to be visited (with the current node
// being the very first element of that list) is replaced by calling strategy
// with said list.
// If the strategy fails, its error is returned.
// Visiting ends if strategy returns no nodes, or if f returns false or an error.
// In the latter case, that error is returned.
// Graph fails with an error if the start node does not belong to the graph.
// The graph should not be modified while visiting, else the behaviour is
// undefined.
func Graph(g digraph.Graph, start digraph.Node, f Func, createStrategy CreateStrategy) error {
	if !g.Contains(start) {
		return digraph.NewNodesNotContainedError(start)
	}

	nodesToVisit := []digraph.Node{
		start,
	}
	strategy := createStrategy(g)

	for {
		current := nodesToVisit[0]
		ok, err := f(current)
		if err != nil {
			return err
		}
		if !ok {
			return nil
		}
		nodesToVisit, err = strategy(current, nodesToVisit[1:])
		if err != nil {
			return err
		}
		if len(nodesToVisit) == 0 {
			return nil
		}
	}
}

// Strategy is every function which takes a current node and a list of nodes
// yet to be visited, and returns a new list of nodes.
type Strategy func(current digraph.Node, rest []digraph.Node) ([]digraph.Node, error)

// CreateStrategy takes a graph and creates a strategy for that graph.
type CreateStrategy func(g digraph.Graph) Strategy

// Func is usually provided by the caller of the visiting functions. It is
// called for every node visited. If returning false and/or an error, the
// visiting stops (and the error, if non-nil, is returned).
type Func func(node digraph.Node) (bool, error)
