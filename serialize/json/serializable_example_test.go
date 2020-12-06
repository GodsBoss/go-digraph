package json_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	jsonGraph "github.com/GodsBoss/go-digraph/serialize/json"

	"encoding/json"
)

func ExampleSerializable() {
	input := digraph.New()
	n1, n2, n3 := input.NewNode(), input.NewNode(), input.NewNode()

	input.Connect(n1, n2)
	input.Connect(n1, n3)
	input.Connect(n2, n3)

	serializer := &jsonGraph.Serializable{
		Graph: input,
		Values: map[digraph.Node]interface{}{
			n1: "a",
			n2: "b",
			n3: "c",
		},
	}

	data, _ := json.Marshal(serializer)

	unserializer := &jsonGraph.Serializable{}
	_ = json.Unmarshal([]byte(data), unserializer)

	nodesToStrings := make(map[digraph.Node]string)

	for _, node := range unserializer.Graph.Nodes() {
		nodesToStrings[node] = unserializer.Values[node].(string)
	}

	for _, edge := range unserializer.Graph.Edges() {
		fmt.Printf(
			"%s connects to %s\n",
			nodesToStrings[edge.Origin],
			nodesToStrings[edge.Destination],
		)
	}

	// Unordered output:
	// a connects to b
	// a connects to c
	// b connects to c
}
