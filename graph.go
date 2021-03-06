package digraph

import (
	"fmt"
)

// New creates an empty, mutable directed graph. That graph is not safe for
// concurrent writes or read/write. It is safe for concurrent reads.
func New() Graph {
	return &graph{
		nodeProvider:        NewNodeProvider(),
		nodes:               make(map[Node]struct{}),
		originToDestination: make(map[Node]map[Node]struct{}),
		destinationToOrigin: make(map[Node]map[Node]struct{}),
	}
}

// Graph represents a directed graph.
type Graph interface {
	// NewNode creates a node inside the graph and returns it.
	NewNode() Node

	// Nodes returns all nodes.
	Nodes() []Node

	// Contains checks wether node is contained in this graph.
	Contains(Node) bool

	// Remove removes the given node from this graph. Returns an error if
	// node is not contained in this graph.
	Remove(Node) error

	// Connect creates an edge from origin to destination. Returns an error if
	// origin or destination or both are not contained in the graph. Also returns
	// an error if there is already a connection from origin to destination (but
	// not if there is a connection from destination to origin).
	Connect(origin, destination Node) error

	// Disconnect removes an existing edge. Returns an error if there is no edge
	// from origin to destination.
	Disconnect(origin, destination Node) error

	// Edges returns all edges of the graph.
	Edges() []Edge

	// PointingTo returns a list of all nodes pointing to the given node.
	// Returns an error if that node does not belong to the graph.
	PointingTo(Node) ([]Node, error)

	// PointedToFrom returns a list of all nodes pointed to from the given node.
	// Returns an error if that node does not belong to the graph.
	PointedToFrom(Node) ([]Node, error)
}

type graph struct {
	// nodeProvider creates new, unique nodes.
	nodeProvider func() Node

	nodes map[Node]struct{}

	// originToDestination is a map from origins to destinations.
	originToDestination map[Node]map[Node]struct{}

	// destinationToOrigin is a map from destinations to origins.
	destinationToOrigin map[Node]map[Node]struct{}
}

func (g *graph) NewNode() Node {
	n := g.nodeProvider()
	g.nodes[n] = struct{}{}
	g.originToDestination[n] = make(map[Node]struct{})
	g.destinationToOrigin[n] = make(map[Node]struct{})
	return n
}

func (g *graph) Nodes() []Node {
	nodes := make([]Node, 0, len(g.nodes))
	for n := range g.nodes {
		nodes = append(nodes, n)
	}
	return nodes
}

func (g *graph) Contains(n Node) bool {
	_, ok := g.nodes[n]
	return ok
}

func (g *graph) nodesNotContainedError(ns ...Node) error {
	notContained := make([]Node, 0)
	for i := range ns {
		if !g.Contains(ns[i]) {
			notContained = append(notContained, ns[i])
		}
	}
	if len(notContained) > 0 {
		return NewNodesNotContainedError(notContained...)
	}
	return nil
}

func (g *graph) Remove(n Node) error {
	if err := g.nodesNotContainedError(n); err != nil {
		return err
	}
	if len(g.originToDestination[n]) > 0 {
		return fmt.Errorf("cannot remove node still connected")
	}
	if len(g.destinationToOrigin[n]) > 0 {
		return fmt.Errorf("cannot remove node still connected")
	}
	delete(g.nodes, n)
	delete(g.originToDestination, n)
	delete(g.destinationToOrigin, n)
	return nil
}

func (g *graph) Connect(origin, destination Node) error {
	if err := g.nodesNotContainedError(origin, destination); err != nil {
		return err
	}
	if _, ok := g.originToDestination[origin][destination]; ok {
		return NewNodesAlreadyConnectedError(origin, destination)
	}
	g.originToDestination[origin][destination] = struct{}{}
	g.destinationToOrigin[destination][origin] = struct{}{}
	return nil
}

func (g *graph) Disconnect(origin, destination Node) error {
	if err := g.nodesNotContainedError(origin, destination); err != nil {
		return err
	}
	if _, ok := g.originToDestination[origin][destination]; !ok {
		return NewNodesNotConnectedError(origin, destination)
	}
	delete(g.originToDestination[origin], destination)
	delete(g.destinationToOrigin[destination], origin)
	return nil
}

func (g *graph) Edges() []Edge {
	edges := make([]Edge, 0)
	for origin := range g.originToDestination {
		for destination := range g.originToDestination[origin] {
			edges = append(
				edges,
				Edge{
					Origin:      origin,
					Destination: destination,
				},
			)
		}
	}
	return edges
}

func (g *graph) PointingTo(n Node) ([]Node, error) {
	if err := g.nodesNotContainedError(n); err != nil {
		return nil, err
	}
	origins := make([]Node, 0, len(g.destinationToOrigin[n]))
	for origin := range g.destinationToOrigin[n] {
		origins = append(origins, origin)
	}
	return origins, nil
}

func (g *graph) PointedToFrom(n Node) ([]Node, error) {
	if err := g.nodesNotContainedError(n); err != nil {
		return nil, err
	}
	destinations := make([]Node, 0, len(g.originToDestination[n]))
	for destination := range g.originToDestination[n] {
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

// Edge is a connection from an origin node to a destination node.
type Edge struct {
	Origin      Node
	Destination Node
}
