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
func Append[V any](ivss ...[]V) []V {
	if ivss == nil {
		return nil
	}
	ovs := []V{}
	for _, vs := range ivss {
		ovs = append(ovs, vs...)
	}
	return ovs
}

// Delete removes the first matching value of the slice.
func Delete[V comparable](v V, ivs []V) []V {
	for i := range ivs {
		if ivs[i] == v {
			var ovs []V
			ovs = append(ivs[:i], ivs[i+1:]...)
			return ovs
		}
	}
	return ivs
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
func Filter[V any](pred func(V) bool, ivs []V) []V {
	if ivs == nil {
		return nil
	}
	ovs := []V{}
	for _, v := range ivs {
		if pred(v) {
			ovs = append(ovs, v)
		}
	}
	return ovs
}

// FilterMap creates a slice from of new values created by fun() where
// it also returns true.
func FilterMap[I, O any](fun func(I) (O, bool), ivs []I) []O {
	if ivs == nil {
		return nil
	}
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

// Split returns the first n values of a slice as first slice and the rest
// as second.
func Split[V any](n int, ivs []V) ([]V, []V) {
	switch {
	case ivs == nil:
		return nil, nil
	case n < 0:
		n = 0
	case n >= len(ivs):
		n = len(ivs) - 1
	}
	lovs := make([]V, n+1)
	rovs := make([]V, len(ivs)-n-1)
	copy(lovs, ivs[:n+1])
	copy(rovs, ivs[n+1:])
	return lovs, rovs
}

// EOF
