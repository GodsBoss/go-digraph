// Package kosaraju provides Kosaraju's algorithm for finding a directed graph's
// strongly connected components.
//
// For more details, see: https://en.wikipedia.org/wiki/Kosaraju%27s_algorithm
package kosaraju

import (
	"github.com/GodsBoss/go-digraph"
)

// StronglyConnectedComponents finds strongly connected components.
//
// The return value is a slice of strongly connected components, each of which
// is a slice of nodes together forming that component.
func StronglyConnectedComponents(g digraph.Graph) [][]digraph.Node {
	vertices := g.Nodes()
	L := make([]digraph.Node, 0)
	visited := make(map[digraph.Node]struct{})

	for i := range vertices {
		visit(g, vertices[i], visited, &L)
	}

	assigned := make(map[digraph.Node]struct{})
	components := make(map[digraph.Node][]digraph.Node)
	for i := range L {
		u := L[i]
		assign(g, u, u, assigned, components)
	}

	result := make([][]digraph.Node, 0, len(components))
	for key := range components {
		result = append(result, components[key])
	}
	return result
}

func visit(
	g digraph.Graph,
	u digraph.Node,
	visited map[digraph.Node]struct{},
	L *[]digraph.Node,
) {
	if _, ok := visited[u]; ok {
		return
	}
	visited[u] = struct{}{}
	outNeighbours, _ := g.PointedToFrom(u)
	for i := range outNeighbours {
		v := outNeighbours[i]
		visit(g, v, visited, L)
	}
	*L = append([]digraph.Node{u}, *L...)
}

func assign(
	g digraph.Graph,
	u digraph.Node,
	root digraph.Node,
	assigned map[digraph.Node]struct{},
	components map[digraph.Node][]digraph.Node,
) {
	if _, ok := assigned[u]; ok {
		return
	}
	components[root] = append(components[root], u)
	assigned[u] = struct{}{}
	inNeighbors, _ := g.PointingTo(u)
	for i := range inNeighbors {
		v := inNeighbors[i]
		assign(g, v, root, assigned, components)
	}
}
