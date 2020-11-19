package operations_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/operations"
)

func ExampleCopy() {
	original := digraph.New()

	a, b, c := original.NewNode(), original.NewNode(), original.NewNode()

	names := map[digraph.Node]string{
		a: "a",
		b: "b",
		c: "c",
	}

	original.Connect(a, b)
	original.Connect(b, c)

	duplicate, originalToCopy := operations.Copy(original)

	// A mapping as returned by Copy() maps nodes from the original to the clone.
	// We need the opposite, too, so we invert it.
	copyToOriginal := operations.InvertedNodeMapping(originalToCopy)

	origins, _ := duplicate.PointingTo(originalToCopy[b])

	fmt.Printf(
		"%s points to %s\n",
		names[copyToOriginal[origins[0]]],
		names[b],
	)

	destinations, _ := duplicate.PointedToFrom(originalToCopy[b])

	fmt.Printf(
		"%s points to %s\n",
		names[b],
		names[copyToOriginal[destinations[0]]],
	)

	// Output:
	// a points to b
	// b points to c
}
