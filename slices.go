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

// Copy copies the values of slices from fpos to tpos into a new slice.
func Copy[V any](ivs []V, fpos, tpos int) []V {
	if ivs == nil || fpos > tpos {
		return nil
	}
	if fpos < 0 {
		fpos = 0
	} else if fpos >= len(ivs) {
		return nil
	}
	if tpos < 0 {
		return nil
	} else if tpos >= len(ivs) {
		tpos = len(ivs) - 1
	}
	ovs := make([]V, tpos-fpos+1)
	copy(ovs, ivs[fpos:tpos+1])
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
		return nil, Copy(ivs, 0, n)
	case n >= len(ivs):
		return Copy(ivs, 0, n), nil
	}
	return Copy(ivs, 0, n), Copy(ivs, n+1, len(ivs))
}

// SplitWith returns the values while pred() returns true as first and the rest
// as second slice.
func SplitWith[V any](pred func(V) bool, ivs []V) ([]V, []V) {
	if ivs == nil || len(ivs) == 0 {
		return nil, nil
	}
	n := -1
	for _, v := range ivs {
		if !pred(v) {
			break
		}
		n++
	}
	return Split(n, ivs)
}

// EOF
