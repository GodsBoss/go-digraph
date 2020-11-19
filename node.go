package digraph

// Node is a node of a directed graph. It is a value object. Can be used as
// map key. It is also possible to compare nodes with ==.
type Node interface {
	internal()
}

type node struct {
	providerID *int
	nodeID     int
}

func (n node) internal() {}

// NewNodeProvider creates a function for creating unique nodes, i.e. nodes
// created by any function are guaranteed to never be equal to any node created
// by any function.
//
// You probably don't need this except when implementing your own Graph type.
func NewNodeProvider() func() Node {
	providerID := 0
	lastNodeID := 0

	return func() Node {
		lastNodeID++
		return node{
			providerID: &providerID,
			nodeID:     lastNodeID,
		}
	}
}
