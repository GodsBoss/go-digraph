package visit

import (
	"github.com/GodsBoss/go-digraph"
)

// Once allows to create strategies which visits nodes at most once.
func Once(createStrategy CreateStrategy) CreateStrategy {

	return func(g digraph.Graph) Strategy {
		strategy := createStrategy(g)
		visitedNodes := make(map[digraph.Node]struct{})

		return func(current digraph.Node, rest []digraph.Node) ([]digraph.Node, error) {
			visitedNodes[current] = struct{}{}
			rest, err := strategy(current, rest)
			if err != nil {
				return nil, err
			}
			filtered := make([]digraph.Node, 0)
			for i := range rest {
				if _, ok := visitedNodes[rest[i]]; !ok {
					filtered = append(filtered, rest[i])
				}
			}
			return filtered, nil
		}
	}
}
