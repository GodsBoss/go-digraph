package uniquemappers

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
)

// String maps nodes to strings bijectively.
type String interface {
	// Add maps node and a string value from the pair to each other. It errors if node
	// or the string are mapped already, even if it is the same pair.
	Add(pair StringNodePair) error

	// Remove removes the connection between a node and a string. Returns an error
	// if node or string (or both) are not mapped.
	Remove(pair StringNodePair) error

	// RemoveNode unmaps the given node, ignoring which string value it is mapped
	// to. Returns an error if the node is not mapped.
	RemoveNode(node digraph.Node) error

	// RemoveString unmaps the given string, ignoring which node it is mapped to.
	// Returns an error if the string is not mapped.
	RemoveString(s string) error

	// GetNodeFor returns the node for a string. It returns an error if no such
	// node exists.
	GetNodeFor(s string) (digraph.Node, error)

	// GetStringFor returns the string for a node. It returns an error if no such
	// string exists.
	GetStringFor(node digraph.Node) (string, error)

	// HasString checks wether the string s is mapped to a node.
	HasString(s string) bool

	// HasNode checks wether node is mapped to a string.
	HasNode(node digraph.Node) bool

	// Nodes returns all nodes this mapper contains in no particular order.
	// The return value is free for modification.
	Nodes() []digraph.Node

	// Strings returns all strings this mapper contains in no particular order.
	// The return value is free for modification.
	Strings() []string
}

// StringNodePair represents a node and its corresponding string (usually some
// form of ID, e.g. a UUID).
type StringNodePair struct {
	Node   digraph.Node
	String string
}

// NewString returns a new unique mapper from nodes to strings (and vice versa).
func NewString() String {
	return &stringMapper{
		nodeToString: make(map[digraph.Node]string),
		stringToNode: make(map[string]digraph.Node),
	}
}

type stringMapper struct {
	nodeToString map[digraph.Node]string
	stringToNode map[string]digraph.Node
}

func (m *stringMapper) Add(pair StringNodePair) error {
	if m.HasNode(pair.Node) {
		return alreadyTakenError{
			node: pair.Node,
		}
	}
	if m.HasString(pair.String) {
		return fmt.Errorf("string already taken")
	}
	m.nodeToString[pair.Node] = pair.String
	m.stringToNode[pair.String] = pair.Node
	return nil
}

func (m *stringMapper) Remove(pair StringNodePair) error {
	if !m.HasNode(pair.Node) || !m.HasString(pair.String) {
		return fmt.Errorf("invalid pair")
	}
	m.remove(pair.Node, pair.String)
	return nil
}

func (m *stringMapper) RemoveNode(node digraph.Node) error {
	s, err := m.GetStringFor(node)
	if err != nil {
		return err
	}
	m.remove(node, s)
	return nil
}

func (m *stringMapper) RemoveString(s string) error {
	node, err := m.GetNodeFor(s)
	if err != nil {
		return err
	}
	m.remove(node, s)
	return nil
}

// remove is called by the Remove* methods after checking wether removal is possible.
func (m *stringMapper) remove(node digraph.Node, s string) {
	delete(m.nodeToString, node)
	delete(m.stringToNode, s)

}

func (m *stringMapper) GetNodeFor(s string) (digraph.Node, error) {
	node, ok := m.stringToNode[s]
	if !ok {
		return nil, fmt.Errorf("string not found")
	}
	return node, nil
}

func (m *stringMapper) GetStringFor(node digraph.Node) (string, error) {
	s, ok := m.nodeToString[node]
	if !ok {
		return "", fmt.Errorf("node not found")
	}
	return s, nil
}

func (m *stringMapper) HasString(s string) bool {
	_, ok := m.stringToNode[s]
	return ok
}

func (m *stringMapper) HasNode(node digraph.Node) bool {
	_, ok := m.nodeToString[node]
	return ok
}

func (m *stringMapper) Nodes() []digraph.Node {
	nodes := make([]digraph.Node, 0, len(m.nodeToString))
	for node := range m.nodeToString {
		nodes = append(nodes, node)
	}
	return nodes
}

func (m *stringMapper) Strings() []string {
	ss := make([]string, 0, len(m.stringToNode))
	for s := range m.stringToNode {
		ss = append(ss, s)
	}
	return ss
}
