package json_test

import (
	jsonGraph "github.com/GodsBoss/go-digraph/serialize/json"

	"encoding/json"
	"fmt"
	"sort"
)

func ExampleSerializable_UnmarshalJSON() {
	input := `
  {
    "nodes": {
      "1": ["2"],
      "2": ["3"],
      "3": []
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

	found := make([]int, 3)

	for i := range nodes {
		ingoing, _ := dg.PointingTo(nodes[i])
		outgoing, _ := dg.PointedToFrom(nodes[i])

		switch true {
		case len(ingoing) == 0 && len(outgoing) == 1:
			found[i] = 1
		case len(ingoing) == 1 && len(outgoing) == 1:
			found[i] = 2
		case len(ingoing) == 1 && len(outgoing) == 0:
			found[i] = 3
		}
	}

	sort.Ints(found)

	for i := range found {
		fmt.Printf("Found node '%d'.\n", found[i])
	}

	// Output:
	// 3 nodes found.
	// Found node '1'.
	// Found node '2'.
	// Found node '3'.
}
