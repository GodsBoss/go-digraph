package digraph

type nodesNotContainedError []Node

func newNodesNotContainedError(nodes ...Node) error {
	return nodesNotContainedError(nodes)
}

func (err nodesNotContainedError) Error() string {
	if len(err) > 1 {
		return "nodes not contained in graph"
	}
	return "node not contained in graph"
}

func (err nodesNotContainedError) nodesNotContained() []Node {
	nodes := make([]Node, len(err))
	copy(nodes, err)
	return nodes
}

// nodesNotContainedProvider is a marker interface for errors returned by a graph
// caused by providing nodes not found in the graph.
type nodesNotContainedProvider interface {
	// nodesNotContained returns the nodes causing the error. The returned slice
	// must be safe for changes.
	nodesNotContained() []Node
}

// IsNodesNotContainedError checks wether an error was caused by providing nodes not
// contained in a graph. If true, the nodes are returned, too. The returned
// slice is safe for changes, i.e. even after changing it, calling the function
// with the same error again will return the same nodes as before.
func IsNodesNotContainedError(err error) (bool, []Node) {
	if nodes, ok := err.(nodesNotContainedProvider); ok {
		return true, nodes.nodesNotContained()
	}
	return false, nil
}
