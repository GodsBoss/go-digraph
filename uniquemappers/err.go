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
	if err.node == nil {
		return StringAlreadyTakenMessage
	}
	if err.s == nil {
		return NodeAlreadyTakenMessage
	}
	return NodeAndStringAlreadyTakenMessage
}

// These are the messages the error caused by trying to map already mapped
// will expose. Do not depend on these exact values, they may change in the
// future! These are merely here for information.
const (
	NodeAndStringAlreadyTakenMessage = "node and string already taken"
	NodeAlreadyTakenMessage          = "node already taken"
	StringAlreadyTakenMessage        = "string already taken"
)

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
