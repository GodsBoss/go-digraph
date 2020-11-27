package abort_test

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
	"github.com/GodsBoss/go-digraph/visit/abort"

	"fmt"
)

func ExampleAfterMaximumExceeded() {
	g := digraph.New()

	a, b, c := g.NewNode(), g.NewNode(), g.NewNode()

	g.Connect(a, b)
	g.Connect(b, c)
	g.Connect(c, a)

	visits := 0

	f := func(_ digraph.Node) (bool, error) {
		visits++
		return true, nil
	}

	strategy := abort.Strategy(abort.AfterMaximumExceeded(1000))(visit.DepthFirst)

	visit.Graph(g, a, f, strategy)

	fmt.Printf("There were %d visits.\n", visits)

	// Output:
	// There were 1001 visits.
}
