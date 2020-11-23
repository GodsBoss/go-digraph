package json_test

import (
	jsonGraph "github.com/GodsBoss/go-digraph/serialize/json"

	"encoding/json"
	"testing"
)

func TestUnmarshalErrors(t *testing.T) {
	inputs := map[string]string{
		"invalid_json":               `{ "nodes": { "7": 3 } }`,
		"invalid_edge_target_node":   `{ "nodes": { "1": ["2"] } }`,
		"duplicate_edge_target_node": `{"nodes": { "1": [], "2": ["1", "1"] } }`,
	}

	for name := range inputs {
		data := []byte(inputs[name])
		t.Run(
			name,
			func(t *testing.T) {
				j := &jsonGraph.Serializable{}
				err := json.Unmarshal(data, j)
				t.Log(err)
				if err == nil {
					t.Errorf("expected error")
				}
			},
		)
	}
}
