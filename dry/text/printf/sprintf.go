/* This is free and unencumbered software released into the public domain. */

// Package printf provides printf(3) text formatting.
package printf

import (
	"fmt"
	. "github.com/dryproject/drylib.go/dry" // nolint: golint
)

////////////////////////////////////////////////////////////////////////////////
// Functions

// Sprintf implements <tt>dry:text/printf/sprintf</tt>.
func Sprintf(format String, args ...interface{}) String {
	return fmt.Sprintf(format, args...)
}
