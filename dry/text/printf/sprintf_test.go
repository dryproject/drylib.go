/* This is free and unencumbered software released into the public domain. */

package printf

import (
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

func TestSprintf(t *testing.T) {
	var actual = Sprintf("%s=%d %s=%.2f", "answer", 42, "pi", 3.1415)
	var expected = "answer=42 pi=3.14"
	if actual != expected {
		t.Errorf("text/printf/sprintf: %q != %q", actual, expected)
	}
}
