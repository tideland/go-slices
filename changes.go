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
// CHANGES
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

// DropWhile removes all values as long pred() returns true.
func DropWhile[V any](pred func(V) bool, vs []V) []V {
	ni := -1
	for i, v := range vs {
		if !pred(v) {
			break
		}
		ni = i
	}
	ni++
	nvs := make([]V, len(vs)-ni)
	copy(nvs, vs[ni:])
	return nvs
}

// Filter creates a slice from all values where pred() returns true.
func Filter[V any](pred func(V) bool, vs []V) []V {
	var nvs []V
	for _, v := range vs {
		if pred(v) {
			nvs = append(nvs, v)
		}
	}
	return nvs
}

// FilterMap creates a slice from of new values created by fun() where
// it also returns true.
func FilterMap[I, O any](fun func(I) (O, bool), ivs []I) []O {
	var novs []O
	for _, iv := range ivs {
		if ov, ok := fun(iv); ok {
			novs = append(novs, ov)
		}
	}
	return novs
}

// EOF