package uniquemappers_test

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/uniquemappers"

	"sort"
	"testing"
)

func TestAddStringMappingErrs(t *testing.T) {
	g := digraph.New()
	n1, n2, n3 := g.NewNode(), g.NewNode(), g.NewNode()
	mapper := uniquemappers.NewString()
	mapper.Add(
		uniquemappers.StringNodePair{
			Node:   n1,
			String: "foo",
		},
	)

	testcases := map[string]struct {
		node     digraph.Node
		str      string
		checkErr func(*testing.T, error)
	}{
		"success": {
			n2,
			"bar",
			func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
			},
		},
		"fail (both exist)": {
			n1,
			"foo",
			func(t *testing.T, err error) {
				if err == nil {
					t.Fatal("expected error")
				}
				ok, node := uniquemappers.IsNodeAlreadyTakenError(err)
				if !ok {
					t.Errorf("expected error to be caused by trying to map a node already mapped")
				}
				if node != n1 {
					t.Errorf("expected n1 as node")
				}
        ok, s := uniquemappers.IsStringAlreadyTakenError(err)
        if !ok {
          t.Error("expected error to be caused by trying to map a string already mapped")
        }
        if s != "foo" {
          t.Error("expected string to be 'foo'")
        }
			},
		},
		"fail (node exists)": {
			n1,
			"baz",
			func(t *testing.T, err error) {
				if err == nil {
					t.Fatal("expected error")
				}
				ok, node := uniquemappers.IsNodeAlreadyTakenError(err)
				if !ok {
					t.Errorf("expected error to be caused by trying to map a node already mapped")
				}
				if node != n1 {
					t.Errorf("expected n1 as node")
				}
        if ok, _ := uniquemappers.IsStringAlreadyTakenError(err); ok {
          t.Errorf("expected error not to be caused by string already mapped")
        }
			},
		},
		"fail (string exists)": {
			n3,
			"foo",
			func(t *testing.T, err error) {
				if err == nil {
					t.Fatal("expected error")
				}
        ok, s := uniquemappers.IsStringAlreadyTakenError(err)
        if !ok {
          t.Error("expected error to be caused by trying to map a string already mapped")
        }
        if s != "foo" {
          t.Error("expected string to be 'foo'")
        }
        if ok, _ := uniquemappers.IsNodeAlreadyTakenError(err); ok {
          t.Errorf("expected error not to be caused by node already mapped")
        }
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				err := mapper.Add(
					uniquemappers.StringNodePair{
						Node:   testcase.node,
						String: testcase.str,
					},
				)
				testcase.checkErr(t, err)
			},
		)
	}
}

func TestStringMappingRemoval(t *testing.T) {
	g := digraph.New()
	n1, n2, n3 := g.NewNode(), g.NewNode(), g.NewNode()

	pair1 := uniquemappers.StringNodePair{Node: n1, String: "1"}
	pair2 := uniquemappers.StringNodePair{Node: n2, String: "2"}
	pair3 := uniquemappers.StringNodePair{Node: n3, String: "3"}

	mapper := uniquemappers.NewString()
	mapper.Add(pair1)
	mapper.Add(pair2)
	mapper.Add(pair3)

	testcases := map[string]struct {
		f    func() error
		pair uniquemappers.StringNodePair
	}{
		"pair 1": {
			pair: pair1,
			f: func() error {
				return mapper.Remove(pair1)
			},
		},
		"pair 2": {
			pair: pair2,
			f: func() error {
				return mapper.RemoveNode(pair2.Node)
			},
		},
		"pair 3": {
			pair: pair3,
			f: func() error {
				return mapper.RemoveString(pair3.String)
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				err := testcase.f()
				if err != nil {
					t.Errorf("expected no error, got %+v", err)
				}
				if mapper.HasString(testcase.pair.String) {
					t.Error("expected mapper not to contain string")
				}
				if mapper.HasNode(testcase.pair.Node) {
					t.Error("expected mapper not to contain node")
				}
			},
		)
	}
}

func TestStringMappingRemovalErrs(t *testing.T) {
	g := digraph.New()
	n1, n2, n3 := g.NewNode(), g.NewNode(), g.NewNode()

	mapper := uniquemappers.NewString()
	mapper.Add(uniquemappers.StringNodePair{Node: n1, String: "foo"})

	testcases := map[string]func() error{
		"invalid node": func() error {
			return mapper.RemoveNode(n2)
		},
		"invalid string": func() error {
			return mapper.RemoveString("bar")
		},
		"invalid node (pair)": func() error {
			return mapper.Remove(uniquemappers.StringNodePair{Node: n2, String: "foo"})
		},
		"invalid string (pair)": func() error {
			return mapper.Remove(uniquemappers.StringNodePair{Node: n1, String: "baz"})
		},
		"invalid pair": func() error {
			return mapper.Remove(uniquemappers.StringNodePair{Node: n3, String: "bax"})
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				err := testcase()
				if err == nil {
					t.Errorf("expected removal to fail")
				}
			},
		)
	}
}

func TestStringMapperExposesContents(t *testing.T) {
	g := digraph.New()
	n1, n2, n3 := g.NewNode(), g.NewNode(), g.NewNode()

	mapper := uniquemappers.NewString()
	mapper.Add(uniquemappers.StringNodePair{Node: n1, String: "foo"})
	mapper.Add(uniquemappers.StringNodePair{Node: n2, String: "bar"})
	mapper.Add(uniquemappers.StringNodePair{Node: n3, String: "baz"})

	t.Run(
		"strings",
		func(t *testing.T) {
			strings := mapper.Strings()
			sort.Strings(strings)
			expected := []string{"bar", "baz", "foo"}

			if len(strings) != len(expected) {
				t.Fatalf("expected %d strings, got %+v", len(expected), strings)
			}
			for i := range expected {
				if expected[i] != strings[i] {
					t.Errorf("expected to find '%s'", expected[i])
				}
			}
		},
	)
	t.Run(
		"nodes",
		func(t *testing.T) {
			nodes := mapper.Nodes()
			if len(nodes) != 3 {
				t.Fatalf("expected 3 nodes, got %d", len(nodes))
			}
			found := make(map[digraph.Node]struct{})
			for i := range nodes {
				if _, ok := found[nodes[i]]; ok {
					t.Errorf("duplicate node")
					continue
				}
				found[nodes[i]] = struct{}{}
			}
			for node := range found {
				switch node {
				case n1:
				case n2:
				case n3:
				default:
					t.Error("unknown node found")
				}
			}
		},
	)
}
