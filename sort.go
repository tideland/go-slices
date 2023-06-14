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
	"math/rand"
	"runtime"

	"golang.org/x/exp/constraints"
)

//--------------------
// SORT
//--------------------

// Sort returns a sorted copy of the given slice of values. It
// uses a parallel quicksort. The type of the slice value has to
// fulfil the constraints.Ordered constraint.
func Sort[V constraints.Ordered](ivs []V) []V {
	less := func(vs []V, i, j int) bool {
		return vs[i] < vs[j]
	}

	return SortWith(ivs, less)
}

// SortsWith reurns a sorted copy of the given slice like Sort().
// Instead of having to fulfil a constraint any value type but can
// be used. The given less function must do the comparison of two
// values.
func SortWith[V any](ivs []V, less func(vs []V, i, j int) bool) []V {
	ovs := Copy(ivs)

	sort(ovs, less)

	return ovs
}

// IsSorted returns true if a slice is sorted in ascending order.
func IsSorted[V constraints.Ordered](vs []V) bool {
	for i := len(vs) - 1; i > 0; i-- {
		if vs[i] < vs[i-1] {
			return false
		}
	}
	return true
}

// IsSortedWith returns true if a slice is sorted in ascending order
// using less as comparison function.
func IsSortedWith[V any](vs []V, less func(a, b V) bool) bool {
	for i := len(vs) - 1; i > 0; i-- {
		if less(vs[i], vs[i-1]) {
			return false
		}
	}
	return true
}

// Shuffle returns a randomly shuffled copy of of the
// given slices of values.
func Shuffle[V any](vs []V) []V {
	ovs := Copy(vs)

	rand.Shuffle(len(ovs), func(i, j int) {
		ovs[i], ovs[j] = ovs[j], ovs[i]
	})

	return ovs
}

//--------------------
// PRIVATE
//--------------------

// sequentialThreshold for switching from sequential quick sort to insertion sort.
var sequentialThreshold = runtime.NumCPU()*4 - 1

// parallelThreshold for switching from parallel to sequential quick sort.
var parallelThreshold = runtime.NumCPU()*2048 - 1

// swap exchanges two values in a slice.
func swap[V any](vs []V, lo, hi int) {
	tmp := vs[lo]
	vs[lo] = vs[hi]
	vs[hi] = tmp
}

// insertionSort for smaller data collections.
func insertionSort[V any](vs []V, less func(vs []V, i, j int) bool, lo, hi int) {
	for i := lo + 1; i < hi+1; i++ {
		for j := i; j > lo && less(vs, j, j-1); j-- {
			swap(vs, j, j-1)
		}
	}
}

// median to caclculate the median based on Tukey's ninther.
func median[V any](vs []V, less func(vs []V, i, j int) bool, lo, hi int) int {
	m := (lo + hi) / 2
	d := (hi - lo) / 8
	// Move median into the middle.
	mot := func(ml, mm, mh int) {
		if less(vs, mm, ml) {
			swap(vs, mm, ml)
		}
		if less(vs, mh, mm) {
			swap(vs, mh, mm)
		}
		if less(vs, mm, ml) {
			swap(vs, mm, ml)
		}
	}
	// Get low, middle, and high median.
	if hi-lo > 40 {
		mot(lo+d, lo, lo+2*d)
		mot(m-d, m, m+d)
		mot(hi-d, hi, hi-2*d)
	}
	// Get combined median.
	mot(lo, m, hi)
	return m
}

// partition the data based on the median.
func partition[V any](vs []V, less func(vs []V, i, j int) bool, lo, hi int) (int, int) {
	med := median(vs, less, lo, hi)
	idx := lo
	swap(vs, med, hi)
	for i := lo; i < hi; i++ {
		if less(vs, i, hi) {
			swap(vs, i, idx)
			idx++
		}
	}
	swap(vs, idx, hi)
	return idx - 1, idx + 1
}

// sequentialQuickSort using itself recursively.
func sequentialQuickSort[V any](vs []V, less func(vs []V, i, j int) bool, lo, hi int) {
	if hi-lo > sequentialThreshold {
		// Use sequential quicksort.
		plo, phi := partition(vs, less, lo, hi)
		sequentialQuickSort(vs, less, lo, plo)
		sequentialQuickSort(vs, less, phi, hi)
	} else {
		// Use insertion sort.
		insertionSort(vs, less, lo, hi)
	}
}

// parallelQuickSort using itself recursively and concurrent.
func parallelQuickSort[V any](vs []V, less func(vs []V, i, j int) bool, lo, hi int, done chan struct{}) {
	if hi-lo > parallelThreshold {
		// Parallel QuickSort.
		plo, phi := partition(vs, less, lo, hi)
		partDone := make(chan struct{})
		go parallelQuickSort(vs, less, lo, plo, partDone)
		go parallelQuickSort(vs, less, phi, hi, partDone)
		// Wait for the end of both sorts.
		<-partDone
		<-partDone
	} else {
		// Sequential QuickSort.
		sequentialQuickSort(vs, less, lo, hi)
	}
	// Signal that it's done.
	done <- struct{}{}
}

// sort starts the parallel quick sort for the whole slice.
func sort[V any](vs []V, less func(vs []V, i, j int) bool) {
	done := make(chan struct{})

	go parallelQuickSort(vs, less, 0, len(vs)-1, done)

	<-done
}

// EOF
