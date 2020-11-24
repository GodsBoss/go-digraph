package visit

import (
	"github.com/GodsBoss/go-digraph"
)

// Func is usually provided by the caller of the visiting functions. It is
// called for every node visited. If returning false and/or an error, the
// visiting stops (and the error, if non-nil, is returned).
type Func func(node digraph.Node) (bool, error)
