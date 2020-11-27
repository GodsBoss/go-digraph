package abort_test

import (
	"strconv"
	"strings"

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

	err := visit.Graph(g, a, f, strategy)

	ok, maximum := abort.IsMaximumExceededError(err)
	if ok {
		fmt.Printf("Maximum node visits before next visit fails was %d.\n", maximum)
	}
	if strings.Contains(err.Error(), strconv.Itoa(1000)) {
		fmt.Println("Error message contains maximum.")
	}

	fmt.Printf("There were %d visits.\n", visits)

	// Output:
	// Maximum node visits before next visit fails was 1000.
	// Error message contains maximum.
	// There were 1001 visits.
}
