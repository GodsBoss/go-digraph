package visit_test

import (
	"fmt"
	"testing"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
)

func TestOnceFailsIfWrappedStrategyFails(t *testing.T) {
	g := digraph.New()
	n := g.NewNode()

	inputErr := fmt.Errorf("i am an error")

	createStrategy := func(g digraph.Graph) visit.Strategy {
		return func(_ digraph.Node, _ []digraph.Node) ([]digraph.Node, error) {
			return nil, inputErr
		}
	}

	f, _ := createFunc(true, nil)

	outputErr := visit.Graph(g, n, f, visit.Once(createStrategy))

	if outputErr != inputErr {
		t.Errorf("expected %+v to equal %+v", outputErr, inputErr)
	}
}
