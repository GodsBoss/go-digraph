// Package json allows for directional graphs to be stored as JSON and restored
// back.
package json

import (
	"github.com/GodsBoss/go-digraph"

	"encoding/json"
)

// Graph is a wrapper type for a directional graph which can be marshaled
// into JSON and unmarshaled back.
type Graph interface {
	json.Marshaler
	json.Unmarshaler

	// Graph returns a directional graph represented by this instance. Subsequent
	// calls return new, distinct graphs, so changes in them do not effect this
	// source of graphs.
	Graph() digraph.Graph
}
