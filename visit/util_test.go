package visit_test

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"
)

func createFunc(ok bool, err error) (visit.Func, map[digraph.Node]struct{}) {
	visitedNodes := make(map[digraph.Node]struct{})
	return func(node digraph.Node) (bool, error) {
		visitedNodes[node] = struct{}{}
		return ok, err
	}, visitedNodes
}
