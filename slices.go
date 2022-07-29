// Tideland Go Slices
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// SLICES
//--------------------

// Append appends the values of all slices to one new slice.
func Append[V any](vss ...[]V) []V {
	all := []V{}
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
	nvs := []V{}
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
	ovs := []O{}
	for _, iv := range ivs {
		if ov, ok := fun(iv); ok {
			ovs = append(ovs, ov)
		}
	}
	return ovs
}

// Join create a slice mixing a separator between each value of the slice.
func Join[V any](sep V, vs []V) []V {
	nvs := []V{}
	last := len(vs) - 1
	for i, v := range vs {
		nvs = append(nvs, v)
		if i < last {
			nvs = append(nvs, sep)
		}
	}
	return nvs
}

// Map creates a slice of output values from the input values and converted
// by the map function.
func Map[I, O any](fun func(I) O, ivs []I) []O {
	ovs := []O{}
	for _, iv := range ivs {
		ovs = append(ovs, fun(iv))
	}
	return ovs
}

// EOF
