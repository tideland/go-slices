// Tideland Go Slices
//
// Copyright (C) 2022-2023 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// IMPORTS
//--------------------

import (
	"golang.org/x/exp/constraints"
)

//--------------------
// SLICES
//--------------------

// Append appends the values of all slices to one new slice.
func Append[V any](ivss ...[]V) []V {
	var ovs []V
	for _, vs := range ivss {
		if vs != nil {
			ovs = append(ovs, vs...)
		}
	}
	return ovs
}

// ContainsAll returns true if the function pred() returns true for all
// values of the slice.
func ContainsAll[V any](ivs []V, pred func(v V) bool) bool {
	for _, v := range ivs {
		if !pred(v) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the function pred() returns true for at least
// one value of the slice.
func ContainsAny[V any](ivs []V, pred func(v V) bool) bool {
	for _, v := range ivs {
		if pred(v) {
			return true
		}
	}
	return false
}

// Copy is simply a convenient combination of allocation and copying.
func Copy[V any](ivs []V) []V {
	if ivs == nil {
		return nil
	}
	ovs := make([]V, len(ivs))
	copy(ovs, ivs)
	return ovs
}

// Delete removes the first matching value of a slice.
func Delete[V comparable](dv V, ivs []V) []V {
	ovs := Copy(ivs)
	for i := range ovs {
		if ovs[i] == dv {
			ovs = append(ovs[:i], ovs[i+1:]...)
			return ovs
		}
	}
	return ovs
}

// DeleteAll removes all matching valus of a slice.
func DeleteAll[V comparable](dv V, ivs []V) []V {
	if ivs == nil {
		return nil
	}
	ovs := []V{}
	for _, v := range ivs {
		if v != dv {
			ovs = append(ovs, v)
		}
	}
	return ovs
}

// DeleteAllWith removes all values of a slice where pred returns true.
func DeleteAllWith[V any](ivs []V, pred func(V) bool) []V {
	_, nsvs := Partition(ivs, pred)
	if nsvs == nil && ivs != nil {
		return []V{}
	}
	return nsvs
}

// DeleteFirst removes the first value of a slice.
func DeleteFirst[V any](ivs []V) []V {
	return Subslice(ivs, 1, len(ivs)-1)
}

// DeleteLast removes the last value of a slice.
func DeleteLast[V any](ivs []V) []V {
	return Subslice(ivs, 0, len(ivs)-2)
}

// DeleteWhile removes all values as long pred() returns true.
func DeleteWhile[V any](ivs []V, pred func(V) bool) []V {
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
	return Subslice(ivs, dropped+1, len(ivs)-1)
}

// DeleteWith removes the first value of a slice where pred returns true.
func DeleteWith[V any](ivs []V, pred func(V) bool) []V {
	ovs := Copy(ivs)
	for i := range ovs {
		if pred(ovs[i]) {
			ovs = append(ovs[:i], ovs[i+1:]...)
			return ovs
		}
	}
	return ovs
}

// Filter creates a slice from all values where pred() returns true.
func Filter[V any](ivs []V, pred func(V) bool) []V {
	svs, _ := Partition(ivs, pred)
	if svs == nil && ivs != nil {
		return []V{}
	}
	return svs
}

// FilterMap creates a slice from of new values created by fun() where
// it also returns true.
func FilterMap[I, O any](ivs []I, fun func(I) (O, bool)) []O {
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
func IsMember[V comparable](v V, ivs []V) bool {
	return ContainsAny(ivs, func(iv V) bool {
		return iv == v
	})
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
func Map[I, O any](ivs []I, fun func(I) O) []O {
	if ivs == nil {
		return nil
	}
	ovs := make([]O, len(ivs))
	for i, iv := range ivs {
		ovs[i] = fun(iv)
	}
	return ovs
}

// Merge merges two slices in a sorted way together.
func Merge[V constraints.Ordered](vsa, vsb []V) []V {
	return Sort(Append(vsa, vsb))
}

// MergeWith merges two slices and uses a comparator function for sorting.
func MergeWith[V any, K constraints.Ordered](vsa, vsb []V, key func(V) K) []V {
	less := func(vs []V, i, j int) bool {
		ki := key(vs[i])
		kj := key(vs[j])
		return ki < kj
	}
	return SortWith(Append(vsa, vsb), less)
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
		return nil, Copy(ivs)
	case n >= len(ivs):
		return Copy(ivs), nil
	}
	return Subslice(ivs, 0, n), Subslice(ivs, n+1, len(ivs)-1)
}

// SplitWith returns the values while pred() returns true as first and the rest
// as second slice.
func SplitWith[V any](ivs []V, pred func(V) bool) ([]V, []V) {
	if len(ivs) == 0 {
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

// Subslice returns the values of slices from fpos to tpos as a new slice.
// Negative fpos as well as too high tpos are allowed and will be limited.
// Starting behind the slice or end before 0 returns nil.
func Subslice[V any](ivs []V, fpos, tpos int) []V {
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

// Subtract returns a new slice that is a copy of input slice, subjected to the following
// procedure: for each element in the subtract slice, its first occurrence in the input
// slice is deleted.
func Subtract[V comparable](ivs, svs []V) []V {
	if ivs == nil || svs == nil {
		return ivs
	}
	ovs := Copy(ivs)
	for _, sv := range svs {
		ovs = Delete(sv, ovs)
	}
	return ovs
}

// TakeWhile copies all values as long pred() returns true.
func TakeWhile[V any](ivs []V, pred func(V) bool) []V {
	if ivs == nil {
		return nil
	}
	taken := -1
	for i, v := range ivs {
		if pred(v) {
			taken = i
			continue
		}
		break
	}
	return Subslice(ivs, 0, taken)
}

// Unique returns a slice which contains each value only once. The second
// and further values are dropped.
func Unique[V comparable](ivs []V) []V {
	if ivs == nil {
		return nil
	}
	ovs := []V{}
	isContained := map[V]struct{}{}
	for _, v := range ivs {
		if _, ok := isContained[v]; !ok {
			ovs = append(ovs, v)
			isContained[v] = struct{}{}
		}
	}
	return ovs
}

// UniqueMerge merges two slices in a sorted way together. Duplicates are dropped.
func UniqueMerge[V constraints.Ordered](vsa, vsb []V) []V {
	return Unique(Sort(Append(vsa, vsb)))
}

// UniqueMergeWith merges two slices and uses a key function to get a sortable key
// of the values. This could e.g. be a field of a struct. Duplicate key values are
// dropped.
func UniqueMergeWith[V any, K constraints.Ordered](vsa, vsb []V, key func(V) K) []V {
	less := func(vs []V, i, j int) bool {
		ki := key(vs[i])
		kj := key(vs[j])
		return ki < kj
	}
	return UniqueWith(SortWith(Append(vsa, vsb), less), key)
}

// UniqueWith returns a slice which contains each value return by the key function
// only once.  The returned value could e.g. be a fiel of a struct.
func UniqueWith[V any, K comparable](ivs []V, key func(V) K) []V {
	if ivs == nil {
		return nil
	}
	ovs := []V{}
	isContained := map[K]struct{}{}
	for _, v := range ivs {
		k := key(v)
		if _, ok := isContained[k]; !ok {
			ovs = append(ovs, v)
			isContained[k] = struct{}{}
		}
	}
	return ovs
}

// Search returns the first value that satisfies the given predicate.
func Search[V any](pred func(v V) bool, ivs []V) (V, bool) {
	for _, v := range ivs {
		if pred(v) {
			return v, true
		}
	}
	// Return default value and false.
	var ov V
	return ov, false
}

// EOF
