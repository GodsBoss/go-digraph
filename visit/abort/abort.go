// Package abort complements the visit package by providing means to abort
// traversal.
package abort

import (
	"github.com/GodsBoss/go-digraph"
	"github.com/GodsBoss/go-digraph/visit"

	"context"
	"fmt"
)

// Strategy wraps a graph visiting strategy and returns a new strategy which
// aborts if condition returns an error.
func Strategy(condition Condition) func(visit.CreateStrategy) visit.CreateStrategy {
	return func(createStrategy visit.CreateStrategy) visit.CreateStrategy {
		return func(g digraph.Graph) visit.Strategy {
			strategy := createStrategy(g)

			return func(current digraph.Node, rest []digraph.Node) ([]digraph.Node, error) {
				if err := condition(); err != nil {
					return nil, err
				}
				return strategy(current, rest)
			}
		}
	}
}

// Condition is used by the abortion strategy to determine wether visiting
// should fail.
type Condition func() error

// AfterMaximumExceeded is a condition which fails after n nodes have been visited.
func AfterMaximumExceeded(n int) Condition {
	current := 0
	return func() error {
		current++
		if current > n {
			return maximumExceededError(n)
		}
		return nil
	}
}

type maximumExceeded interface {
	maximum() int
}

type maximumExceededError int

func (err maximumExceededError) Error() string {
	return fmt.Sprintf("node visitation maximum of %d exceeded", err.maximum())
}

func (err maximumExceededError) maximum() int {
	return int(err)
}

// IsMaximumExceededError checks wether an error was caused by exceeding the
// maximum number of nodes to visit when using AfterMaximumExceeded() as abort
// condition.
func IsMaximumExceededError(err error) (ok bool, maximum int) {
	if mErr, ok := err.(maximumExceededError); ok {
		return true, mErr.maximum()
	}
	return false, 0
}

// OnContextCancelation is a condition which fails if a context has been canceled.
// The context's error is returned.
func OnContextCancelation(ctx context.Context) Condition {
	return func() error {
		if err := ctx.Err(); err != nil {
			return err
		}
		return nil
	}
}
