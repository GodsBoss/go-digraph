package abort_test

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
	"github.com/GodsBoss/go-digraph/visit/abort"

	"context"
	"fmt"
)

func ExampleOnContextCancelation() {
	g := digraph.New()

	a, b, c := g.NewNode(), g.NewNode(), g.NewNode()

	g.Connect(a, b)
	g.Connect(b, c)
	g.Connect(c, a)

	ctx, cancel := context.WithCancel(context.Background())

	visits := 0

	f := func(_ digraph.Node) (bool, error) {
		visits++
		if visits > 1000 {
			cancel()
		}
		return true, nil
	}

	strategy := abort.Strategy(abort.OnContextCancelation(ctx))(visit.DepthFirst)

	visit.Graph(g, a, f, strategy)

	fmt.Printf("There were %d visits.\n", visits)

	// Output:
	// There were 1001 visits.
}
