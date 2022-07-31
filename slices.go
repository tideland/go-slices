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
func DropWhile[V any](pred func(V) bool, ivs []V) []V {
	if ivs == nil {
		return nil
	}
	dropped := -1
	for i, v := range ivs {
		if pred(v) {
			dropped = i
			continue
		}
		break
	}
	if dropped == len(ivs)-1 {
		return nil
	}
	ovs := make([]V, len(ivs)-dropped-1)
	copy(ovs, ivs[dropped+1:])
	return ovs
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
func Join[V any](sep V, ivs []V) []V {
	if ivs == nil {
		return nil
	}
	ovs := []V{}
	last := len(ivs) - 1
	for i, v := range ivs {
		ovs = append(ovs, v)
		if i < last {
			ovs = append(ovs, sep)
		}
	}
	return ovs
}

// Map creates a slice of output values from the input values and converted
// by the map function.
func Map[I, O any](fun func(I) O, ivs []I) []O {
	if ivs == nil {
		return nil
	}
	ovs := make([]O, len(ivs))
	for i, iv := range ivs {
		ovs[i] = fun(iv)
	}
	return ovs
}

// Reverse returns the slice in reverse order.
func Reverse[V any](ivs []V) []V {
	if ivs == nil {
		return nil
	}
	l := len(ivs)
	ovs := make([]V, l)
	for i := range ivs {
		l--
		ovs[i] = ivs[l]
	}
	return ovs
}

// EOF
