package digraph_test

import (
	"fmt"
	"sort"

	"github.com/GodsBoss/go-digraph"
)

func ExampleGraph() {
	// Create a new graph.
	graph := digraph.New()

	// Create some nodes. They become part of the graph automatically.
	png := graph.NewNode()
	image := graph.NewNode()
	bufio := graph.NewNode()
	io := graph.NewNode()

	// Connect nodes.
	graph.Connect(png, io)
	graph.Connect(png, image)
	graph.Connect(image, bufio)
	graph.Connect(image, io)

	// Nodes can be used as map values.
	packageNames := map[digraph.Node]string{
		png:   "image/png",
		image: "image",
		bufio: "bufio",
		io:    "io",
	}

	edges := graph.Edges()
	lines := make([]string, len(edges))
	for i := range edges {
		lines[i] = fmt.Sprintf(
			"%s depends on %s",
			packageNames[edges[i].Origin],
			packageNames[edges[i].Destination],
		)
	}
	sort.Strings(lines) // Make output reproducible.

	for i := range lines {
		fmt.Println(lines[i])
	}

	// Output:
	// image depends on bufio
	// image depends on io
	// image/png depends on image
	// image/png depends on io
}
