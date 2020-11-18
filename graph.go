package digraph

import "fmt"

// New creates an empty, mutable directed graph.
func New() Graph {
	id := 0
	return &graph{
		id:    &id,
		nodes: make(map[Node]struct{}),
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
}

type graph struct {
	id         *int
	lastNodeID int
	nodes      map[Node]struct{}
}

func (g *graph) NewNode() Node {
	g.lastNodeID++
	n := node{
		graphID: g.id,
		nodeID:  g.lastNodeID,
	}
	g.nodes[n] = struct{}{}
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
	delete(g.nodes, n)
	return nil
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
