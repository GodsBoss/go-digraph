package uniquemappers

import (
	"github.com/GodsBoss/go-digraph"
)

// IsNodeAlreadyTakenError checks wether an error was caused by trying to add
// a node that has already been mapped. The offending node is also exposed.
func IsNodeAlreadyTakenError(err error) (bool, digraph.Node) {
	if natErr, ok := err.(nodeAlreadyTaken); ok {
		return natErr.nodeAlreadyTaken()
	}
	return false, nil
}

type nodeAlreadyTaken interface {
	nodeAlreadyTaken() (bool, digraph.Node)
}

type alreadyTakenError struct {
	node digraph.Node
}

func (err alreadyTakenError) Error() string {
	return "node already taken"
}

func (err alreadyTakenError) nodeAlreadyTaken() (bool, digraph.Node) {
	return true, err.node
}
