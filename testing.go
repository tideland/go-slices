// Tideland Go Slices
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// TESTING
//--------------------

// ContainsAll returns true if the function pred() returns true for all
// values of the slice.
func ContainsAll[V any](pred func(v V) bool, vs []V) bool {
	for _, v := range vs {
		if !pred(v) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the function pred() returns true for at least
// one value of the slice.
func ContainsAny[V any](pred func(v V) bool, vs []V) bool {
	for _, v := range vs {
		if pred(v) {
			return true
		}
	}
	return false
}

// IsEqual returns true if both slices are equal.
func IsEqual[V comparable](first, second []V) bool {
	if len(first) != len(second) {
		return false
	}
	for i, v := range first {
		if v != second[i] {
			return false
		}
	}
	return true
}

// IsMember returns true if the slice contains the value v.
func IsMember[V comparable](v V, vs []V) bool {
	return ContainsAny(func(iv V) bool {
		return iv == v
	}, vs)
}

// IsPrefix returns true if the first slice is the prefix of the second one.
func IsPrefix[V comparable](prefix, all []V) bool {
	if len(prefix) > len(all) {
		return false
	}
	for i, v := range prefix {
		if v != all[i] {
			return false
		}
	}
	return true
}

// IsSuffix returns true if the first slice is the suffix of the second one.
func IsSuffix[V comparable](suffix, all []V) bool {
	if len(suffix) > len(all) {
		return false
	}
	diff := len(all) - len(suffix)
	for i, v := range suffix {
		if v != all[i+diff] {
			return false
		}
	}
	return true
}

// EOF
