// Tideland Go Slices
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// IMPORTS
//--------------------

//--------------------
// CONSTANTS
//--------------------

//--------------------
// TRANSFORMING
//--------------------

// Append appends the values of all slices to one new slice.
func Append[V any](vss ...[]V) []V {
	var all []V
	for _, vs := range vss {
		all = append(all, vs...)
	}
	return all
}

// Delete removes the first matching value of the slice.
func Delete[V comparable](v V, vs []V) []V {
	for i := range vs {
		if vs[i] == v {
			var vsn []V
			vsn = append(append(vsn, vs[:i]...), vs[i+1:]...)
			return vsn
		}
	}
	return vs
}

// EOF
