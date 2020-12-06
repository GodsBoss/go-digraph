// Package uniquemappers contains bi-directional mappers from digraph.Nodes to
// some basic types. These are useful for mapping between graph nodes and your
// objects, identified by strings (e.g. UUIDs) or integer IDs.
//
// Mappers do not check for graph plausibility. It is perfectly valid for a
// mapper to contain nodes from non-identical graphs.
package uniquemappers
