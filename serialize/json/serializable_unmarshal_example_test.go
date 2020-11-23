package json_test

import (
	"github.com/GodsBoss/go-digraph"
	jsonGraph "github.com/GodsBoss/go-digraph/serialize/json"

	"encoding/json"
	"fmt"
)

func ExampleSerializable_UnmarshalJSON() {
	input := `
  {
    "nodes": {
      "1": ["2"],
      "2": ["3"],
      "3": []
    },
		"values": {
			"1": "a",
			"2": "b",
			"3": "c"
		}
  }
  `

	s := &jsonGraph.Serializable{}
	err := json.Unmarshal([]byte(input), s)

	if err != nil {
		fmt.Printf("error on marshal: %+v\n", err)
		return
	}

	dg := s.Graph
	nodes := dg.Nodes()
	fmt.Printf("%d nodes found.\n", len(nodes))

	found := make(map[int]digraph.Node)

	for i := range nodes {
		ingoing, _ := dg.PointingTo(nodes[i])
		outgoing, _ := dg.PointedToFrom(nodes[i])

		switch true {
		case len(ingoing) == 0 && len(outgoing) == 1:
			found[1] = nodes[i]
		case len(ingoing) == 1 && len(outgoing) == 1:
			found[2] = nodes[i]
		case len(ingoing) == 1 && len(outgoing) == 0:
			found[3] = nodes[i]
		}
	}

	for i := 1; i <= 3; i++ {
		fmt.Printf("Found node '%s'.\n", s.Values[found[i]])
	}

	// Output:
	// 3 nodes found.
	// Found node 'a'.
	// Found node 'b'.
	// Found node 'c'.
}
