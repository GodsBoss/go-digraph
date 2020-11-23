package json_test

import (
	"github.com/GodsBoss/go-digraph"
	jsonGraph "github.com/GodsBoss/go-digraph/serialize/json"

	"encoding/json"
	"fmt"
	"sort"
)

func ExampleSerializable_MarshalJSON() {
	g := digraph.New()
	n1, n2, n3 := g.NewNode(), g.NewNode(), g.NewNode()

	g.Connect(n1, n2)
	g.Connect(n1, n3)
	g.Connect(n2, n3)

	s := &jsonGraph.Serializable{
		Graph: g,
	}

	data, _ := json.Marshal(s)

	myGraph := struct {
		M map[string][]string `json:"nodes"`
	}{}

	_ = json.Unmarshal(data, &myGraph)

	nameMapping := make(map[string]string)

	keys := []string{"1", "2", "3"}
	for _, key := range keys {
		switch len(myGraph.M[key]) {
		case 2:
			nameMapping["a"] = key
		case 1:
			nameMapping["b"] = key
		case 0:
			nameMapping["c"] = key
		}
	}

	invertedNameMapping := make(map[string]string)
	for key, value := range nameMapping {
		invertedNameMapping[value] = key
	}

	namedKeys := []string{"a", "b", "c"}
	for _, namedKey := range namedKeys {
		origin := nameMapping[namedKey]
		mappedDestinations := make([]string, 0)

		for _, destination := range myGraph.M[origin] {
			mappedDestinations = append(mappedDestinations, invertedNameMapping[destination])
		}
		sort.Strings(mappedDestinations)

		for _, mappedDestination := range mappedDestinations {
			fmt.Printf("%s points to %s.\n", namedKey, mappedDestination)
		}
	}

	// Output:
	// a points to b.
	// a points to c.
	// b points to c.
}
