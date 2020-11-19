package operations_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/operations"
)

func ExampleTransposed() {
	original := digraph.New()

	a, b, c := original.NewNode(), original.NewNode(), original.NewNode()

	names := map[digraph.Node]string{
		a: "a",
		b: "b",
		c: "c",
	}

	original.Connect(a, b)
	original.Connect(b, c)

	transposed, originalToTransposed := operations.Transposed(original)

	// A mapping as returned by Copy() maps nodes from the original to the clone.
	// We need the opposite, too, so we invert it.
	transposedToOriginal := operations.InvertedNodeMapping(originalToTransposed)

	origins, _ := transposed.PointingTo(originalToTransposed[b])

	fmt.Printf(
		"%s points to %s\n",
		names[transposedToOriginal[origins[0]]],
		names[b],
	)

	destinations, _ := transposed.PointedToFrom(originalToTransposed[b])

	fmt.Printf(
		"%s points to %s\n",
		names[b],
		names[transposedToOriginal[destinations[0]]],
	)

	// Output:
	// c points to b
	// b points to a
}
