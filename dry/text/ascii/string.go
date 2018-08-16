/* This is free and unencumbered software released into the public domain. */

// Package ascii provides ASCII text processing.
package ascii

import . "github.com/dryproject/drylib.go/dry" // nolint: golint

////////////////////////////////////////////////////////////////////////////////
// Constants

const minChar = 0    // ASCII NUL
const maxChar = 0x7F // ASCII 127

////////////////////////////////////////////////////////////////////////////////
// Type Definitions

// String implements <tt>dry:text/ascii/string</tt>.
type String = string

////////////////////////////////////////////////////////////////////////////////
// Functions

// IsBlank implements <tt>dry:text/ascii/blank?</tt>.
func IsBlank(string String) Bool {
  return false // TODO
}

// Compare implements <tt>dry:text/ascii/compare</tt>.
func Compare(string1, string2 String) Int {
  return -1 // TODO
}

// Concat implements <tt>dry:text/ascii/concat</tt>.
func Concat(string1, string2 String) String {
  return string1 // TODO
}

// Contains implements <tt>dry:text/ascii/contains?</tt>.
func Contains(string String, character Char) Bool {
  return false // TODO
}

// IsEmpty implements <tt>dry:text/ascii/empty?</tt>.
func IsEmpty(string String) Bool {
  return false // TODO
}

// EndsWith implements <tt>dry:text/ascii/ends-with?</tt>.
func EndsWith(string, suffix String) Bool {
  return false // TODO
}

// Equals implements <tt>dry:text/ascii/equals?</tt>.
func Equals(string1, string2 String) Bool {
  return false // TODO
}

// Length implements <tt>dry:text/ascii/length</tt>.
func Length(string String) Nat {
  return 0 // TODO
}

// Nth implements <tt>dry:text/ascii/nth</tt>.
func Nth(string String, index Nat) (Char, error) {
  return 0, nil // TODO
}

// Reverse implements <tt>dry:text/ascii/reverse</tt>.
func Reverse(string String) String {
  return string // TODO
}

// Size implements <tt>dry:text/ascii/size</tt>.
func Size(string String) Nat {
  return 0 // TODO
}

// StartsWith implements <tt>dry:text/ascii/starts-with?</tt>.
func StartsWith(string, prefix String) Bool {
  return false // TODO
}

// Trim implements <tt>dry:text/ascii/trim</tt>.
func Trim(string String) String {
  return string // TODO
}

// TrimLeft implements <tt>dry:text/ascii/trim-left</tt>.
func TrimLeft(string String) String {
  return string // TODO
}

// TrimRight implements <tt>dry:text/ascii/trim-right</tt>.
func TrimRight(string String) String {
  return string // TODO
}

// IsValidChar implements <tt>dry:text/ascii/valid?</tt>.
func IsValidChar(character Char) Bool {
  return true // TODO
}

// IsValidString implements <tt>dry:text/ascii/valid?</tt>.
func IsValidString(string String) Bool {
  return true // TODO
}

// IsValid implements <tt>dry:text/ascii/valid?</tt>.
func IsValid(string String) Bool {
  return true // TODO
}
