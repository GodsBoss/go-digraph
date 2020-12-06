// Package json allows for directional graphs to be stored as JSON and restored
// back.
package json

import (
	"github.com/GodsBoss/go-digraph"

	"encoding/json"
	"fmt"
	"strconv"
)

// Serializable is a wrapper type for a directional graph which can be marshaled
// into JSON and unmarshaled back. The concrete serialization format used, e.g.
// JSON keys and structures, is NOT part of the interface of the package. It may
// change at any point in the future.
type Serializable struct {
	Graph digraph.Graph

	// Values stores arbitrary values alongside the graph. It must not contain any
	// nodes not contained in the graph.
	Values map[digraph.Node]interface{}
}

// MarshalJSON lets Serializable implement json.Marshaler.
func (s *Serializable) MarshalJSON() ([]byte, error) {
	dg := s.Graph
	if dg == nil {
		dg = digraph.New()
	}
	values := s.Values
	if values == nil {
		values = make(map[digraph.Node]interface{})
	}
	for n := range values {
		if !dg.Contains(n) {
			return nil, fmt.Errorf("values contains node not found in graph")
		}
	}
	j := &jsonGraph{
		Nodes:  make(map[string][]string),
		Values: make(map[string]interface{}),
	}
	nodes := dg.Nodes()
	nodeKeyMapping := make(map[digraph.Node]string)
	for i := range nodes {
		nodeKey := strconv.Itoa(i + 1)
		j.Nodes[nodeKey] = make([]string, 0)
		nodeKeyMapping[nodes[i]] = nodeKey
		if val, ok := values[nodes[i]]; ok {
			j.Values[nodeKey] = val
		}
	}
	edges := dg.Edges()
	for i := range edges {
		origin := nodeKeyMapping[edges[i].Origin]
		destination := nodeKeyMapping[edges[i].Destination]
		j.Nodes[origin] = append(j.Nodes[origin], destination)
	}
	return json.Marshal(j)
}

// UnmarshalJSON lets Serializable implement json.Unmarshaler.
func (s *Serializable) UnmarshalJSON(data []byte) error {
	j := &jsonGraph{}
	err := json.Unmarshal(data, j)
	if err != nil {
		return err
	}

	dg := digraph.New()
	values := make(map[digraph.Node]interface{})

	jNodesToNodes := make(map[string]digraph.Node)
	for i := range j.Nodes {
		jNodesToNodes[i] = dg.NewNode()
		if val, ok := j.Values[i]; ok {
			values[jNodesToNodes[i]] = val
		}
	}
	for i := range j.Nodes {
		for k := range j.Nodes[i] {
			to, ok := jNodesToNodes[j.Nodes[i][k]]
			if !ok {
				return fmt.Errorf("invalid edge target node id %s", j.Nodes[i][k])
			}
			err := dg.Connect(jNodesToNodes[i], to)
			if err != nil {
				return err
			}
		}
	}
	s.Graph = dg
	s.Values = values
	return nil
}

type jsonGraph struct {
	Nodes  map[string][]string    `json:"nodes"`
	Values map[string]interface{} `json:"values,omitempty"`
}
