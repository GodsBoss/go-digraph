package visit_test

import (
	"fmt"

	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"

	"testing"
)

func TestGraphFailsWithUnknownNode(t *testing.T) {
	g := digraph.New()
	n := digraph.New().NewNode()

	nopFunc := func(_ digraph.Node) (bool, error) {
		return true, nil
	}

	err := visit.Graph(g, n, nopFunc, visit.DepthFirst)
	if err == nil {
		t.Errorf("expected an error")
		return
	}

	ok, errNodes := digraph.IsNodesNotContainedError(err)
	if !ok {
		t.Errorf("expected 'node not contained' error")
	}
	if len(errNodes) != 1 {
		t.Errorf("expected 1 node, got %d", len(errNodes))
		return
	}
	if errNodes[0] != n {
		t.Errorf("unexpected node")
	}
}

func TestGraphReturnsFuncsError(t *testing.T) {
	g := digraph.New()
	n := g.NewNode()
	dest := g.NewNode()
	g.Connect(n, dest)

	inputErr := fmt.Errorf("some error")
	f, visitedNodes := createFunc(true, inputErr)

	outputErr := visit.Graph(g, n, f, visit.DepthFirst)

	if outputErr != inputErr {
		t.Errorf("expected %+v, but got %+v", inputErr, outputErr)
	}
	if _, ok := visitedNodes[dest]; ok {
		t.Errorf("expected destination node not to have been visited")
	}
}

func TestFuncReturnsFalse(t *testing.T) {
	g := digraph.New()
	n := g.NewNode()
	dest := g.NewNode()
	g.Connect(n, dest)

	f, visitedNodes := createFunc(false, nil)

	err := visit.Graph(g, n, f, visit.DepthFirst)

	if err != nil {
		t.Errorf("expected no error, but got %+v", err)
	}
	if _, ok := visitedNodes[dest]; ok {
		t.Errorf("expected destination node not to have been visited")
	}
}

func TestStrategyFails(t *testing.T) {
	g := digraph.New()
	n := g.NewNode()

	inputErr := fmt.Errorf("i am an error")

	createStrategy := func(g digraph.Graph) visit.Strategy {
		return func(_ digraph.Node, _ []digraph.Node) ([]digraph.Node, error) {
			return nil, inputErr
		}
	}

	f, _ := createFunc(true, nil)

	outputErr := visit.Graph(g, n, f, createStrategy)

	if outputErr != inputErr {
		t.Errorf("expected error to be %+v, but got %+v", inputErr, outputErr)
	}
}

func createFunc(ok bool, err error) (visit.Func, map[digraph.Node]struct{}) {
	visitedNodes := make(map[digraph.Node]struct{})
	return func(node digraph.Node) (bool, error) {
		visitedNodes[node] = struct{}{}
		return ok, err
	}, visitedNodes
}
