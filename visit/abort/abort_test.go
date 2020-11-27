package abort_test

import (
	"github.com/GodsBoss/go-digraph/visit/abort"

	"fmt"
	"testing"
)

func TestMaximumExceededErrorDetection(t *testing.T) {
	ok, maximum := abort.IsMaximumExceededError(fmt.Errorf("some random error"))
	if ok {
		t.Errorf("expected error not to be a 'maximum exceeded' error")
	}
	if maximum != 0 {
		t.Errorf("expected maximum to be 0, but got %d", maximum)
	}
}
