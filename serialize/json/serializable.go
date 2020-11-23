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
// into JSON and unmarshaled back.
type Serializable struct {
	Graph digraph.Graph
}

// MarshalJSON lets Serializable implement json.Marshaler.
func (s *Serializable) MarshalJSON() ([]byte, error) {
	dg := s.Graph
	if dg == nil {
		dg = digraph.New()
	}
	j := &jsonGraph{
		Nodes: make(map[string][]string),
	}
	nodes := dg.Nodes()
	nodeKeyMapping := make(map[digraph.Node]string)
	for i := range nodes {
		nodeKey := strconv.Itoa(i + 1)
		j.Nodes[nodeKey] = make([]string, 0)
		nodeKeyMapping[nodes[i]] = nodeKey
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

	jNodesToNodes := make(map[string]digraph.Node)
	for i := range j.Nodes {
		jNodesToNodes[i] = dg.NewNode()
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
	return nil
}

type jsonGraph struct {
	Nodes map[string][]string `json:"nodes"`
}
