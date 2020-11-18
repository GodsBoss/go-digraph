package digraph

import (
	"fmt"
)

// New creates an empty, mutable directed graph.
func New() Graph {
	id := 0
	return &graph{
		id:                  &id,
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
	id         *int
	lastNodeID int
	nodes      map[Node]struct{}

	// originToDestination is a map from origins to destinations.
	originToDestination map[Node]map[Node]struct{}

	// destinationToOrigin is a map from destinations to origins.
	destinationToOrigin map[Node]map[Node]struct{}
}

func (g *graph) NewNode() Node {
	g.lastNodeID++
	n := node{
		graphID: g.id,
		nodeID:  g.lastNodeID,
	}
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

func (g *graph) Remove(n Node) error {
	if !g.Contains(n) {
		return fmt.Errorf("graph did not contain given node")
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
	if !g.Contains(origin) {
		return fmt.Errorf("origin not contained in graph")
	}
	if !g.Contains(destination) {
		return fmt.Errorf("destination not contained in graph")
	}
	if _, ok := g.originToDestination[origin][destination]; ok {
		return fmt.Errorf("already connected")
	}
	g.originToDestination[origin][destination] = struct{}{}
	g.destinationToOrigin[destination][origin] = struct{}{}
	return nil
}

func (g *graph) Disconnect(origin, destination Node) error {
	if !g.Contains(origin) {
		return fmt.Errorf("origin not contained in graph")
	}
	if !g.Contains(destination) {
		return fmt.Errorf("destination not contained in graph")
	}
	if _, ok := g.originToDestination[origin][destination]; !ok {
		return fmt.Errorf("not connected")
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
	if !g.Contains(n) {
		return nil, fmt.Errorf("node not contained in graph")
	}
	origins := make([]Node, 0, len(g.destinationToOrigin[n]))
	for origin := range g.destinationToOrigin[n] {
		origins = append(origins, origin)
	}
	return origins, nil
}

func (g *graph) PointedToFrom(n Node) ([]Node, error) {
	if !g.Contains(n) {
		return nil, fmt.Errorf("node not contained in graph")
	}
	destinations := make([]Node, 0, len(g.originToDestination[n]))
	for destination := range g.originToDestination {
		destinations = append(destinations, destination)
	}
	return destinations, nil
}

// Node is a node of a directed graph. It is a value object. Can be used as
// map key. It is also possible to compare nodes with ==.
type Node interface {
	internal()
}

type node struct {
	graphID *int
	nodeID  int
}

func (n node) internal() {}

// Edge is a connection from an origin node to a destination node.
type Edge struct {
	Origin      Node
	Destination Node
}
