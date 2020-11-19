package digraph

type nodesNotContainedError []Node

// NewNodesNotContainedError creates an error which signals that it was caused
// by providing nodes not contained in a graph. It implements
// NodesNotContainedProvider.
func NewNodesNotContainedError(nodes ...Node) error {
	return nodesNotContainedError(nodes)
}

func (err nodesNotContainedError) Error() string {
	if len(err) > 1 {
		return "nodes not contained in graph"
	}
	return "node not contained in graph"
}

func (err nodesNotContainedError) NodesNotContained() []Node {
	nodes := make([]Node, len(err))
	copy(nodes, err)
	return nodes
}

// NodesNotContainedProvider is a marker interface for errors returned by a graph
// caused by providing nodes not found in the graph. Instead of doing the type
// assertion yourself, use IsNodesNotContainedError for your convenience.
//
// You will need this if you implement errors for a graph implementation yourself.
// Most probably, just creating such errors with NewNodesNotContainedError is a
// better choice.
type NodesNotContainedProvider interface {
	// nodesNotContained returns the nodes causing the error. The returned slice
	// must be safe for changes.
	NodesNotContained() []Node
}

// IsNodesNotContainedError checks wether an error was caused by providing nodes not
// contained in a graph. If true, the nodes are returned, too. The returned
// slice is safe for changes, i.e. even after changing it, calling the function
// with the same error again will return the same nodes as before.
func IsNodesNotContainedError(err error) (bool, []Node) {
	if nodes, ok := err.(NodesNotContainedProvider); ok {
		return true, nodes.NodesNotContained()
	}
	return false, nil
}

type nodesAlreadyConnectedError struct {
	origin      Node
	destination Node
}

func (err nodesAlreadyConnectedError) Error() string {
	return "nodes already connected"
}

func (err nodesAlreadyConnectedError) Origin() Node {
	return err.origin
}

func (err nodesAlreadyConnectedError) Destination() Node {
	return err.destination
}

func (err nodesAlreadyConnectedError) AlreadyConnected() {}

// AlreadyConnected is a marker interface implemented by errors caused by
// connecting already connected nodes.
type AlreadyConnected interface {
	// Origin returns the origin node of the attempt.
	Origin() Node

	// Destination returns the destination node of the attempt.
	Destination() Node

	// AlreadyConnected is for marking and does nothing.
	AlreadyConnected()
}

// IsNodesAlreadyConnectedError checks wether an error was caused by connecting
// two already connected nodes. If true, the corresponding edge is returned as
// well.
// The edge is guaranteed to be safe for changes.
func IsNodesAlreadyConnectedError(err error) (bool, *Edge) {
	if alreadyConnected, ok := err.(AlreadyConnected); ok {
		return true, &Edge{
			Origin:      alreadyConnected.Origin(),
			Destination: alreadyConnected.Destination(),
		}
	}
	return false, nil
}
