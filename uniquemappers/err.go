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

// IsStringAlreadyTakenError checks wether an error was caused by trying to add
// a string that has already been mapped. The offending string is also exposed.
func IsStringAlreadyTakenError(err error) (bool, string) {
	if satErr, ok := err.(stringAlreadyTaken); ok {
		return satErr.stringAlreadyTaken()
	}
	return false, ""
}

type stringAlreadyTaken interface {
	stringAlreadyTaken() (bool, string)
}

type alreadyTakenError struct {
	node digraph.Node
	s    *string
}

func (err alreadyTakenError) Error() string {
	if err.node != nil && err.s != nil {
		return "node and string already taken"
	}
	if err.node != nil {
		return "node already taken"
	}
	if err.s != nil {
		return "string already taken"
	}
	return ""
}

func (err alreadyTakenError) nodeAlreadyTaken() (bool, digraph.Node) {
	if err.node == nil {
		return false, nil
	}
	return true, err.node
}

func (err alreadyTakenError) stringAlreadyTaken() (bool, string) {
	if err.s == nil {
		return false, ""
	}
	return true, *err.s
}
