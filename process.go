// Tideland Go Slices
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// COMPUTATIONS
//--------------------

// FoldL iterates over the slice from left to right. It calls fun() for
// each value passing the initial accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will
// be returned.
func FoldL[V, Acc any](vs []V, acc Acc, fun func(V, Acc) Acc) Acc {
	for _, v := range vs {
		acc = fun(v, acc)
	}
	return acc
}

// FoldLFirst iterates over the slice from left to right. It calls fun() for
// each value passing the first value as accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldLFirst[V any](vs []V, fun func(V, V) V) V {
	var first V
	if len(vs) == 0 {
		// Return default value.
		return first
	}
	first = vs[0]
	return FoldL(vs[1:], first, fun)
}

// FoldR iterates over the slice from right to left. It calls fun() for
// each value passing the initial accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldR[V, Acc any](vs []V, acc Acc, fun func(V, Acc) Acc) Acc {
	for i := len(vs) - 1; i >= 0; i-- {
		acc = fun(vs[i], acc)
	}
	return acc
}

// FoldRLast iterates over the slice from right to left. It calls fun() for
// each value passing the last value as accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldRLast[V any](vs []V, fun func(V, V) V) V {
	var last V
	if len(vs) == 0 {
		// Return default value.
		return last
	}
	last = vs[len(vs)-1]
	return FoldR(vs[:len(vs)-1], last, fun)
}

// MapFoldL combines the operations of Map() and FoldL() in one pass.
func MapFoldL[I, O, Acc any](ivs []I, acc Acc, fun func(I, Acc) (O, Acc)) ([]O, Acc) {
	var ov O
	var ovs []O
	if ivs != nil {
		ovs = make([]O, len(ivs))
	}
	for i, iv := range ivs {
		ov, acc = fun(iv, acc)
		ovs[i] = ov
	}
	return ovs, acc
}

// MapFoldR combines the operations of Map() and FoldR() in one pass.
func MapFoldR[I, O, Acc any](ivs []I, acc Acc, fun func(I, Acc) (O, Acc)) ([]O, Acc) {
	var ov O
	var ovs []O
	if ivs != nil {
		ovs = make([]O, len(ivs))
	}
	for i := len(ivs) - 1; i >= 0; i-- {
		ov, acc = fun(ivs[i], acc)
		ovs[i] = ov
	}
	return ovs, acc
}

// Partition checks all values of the slices and returns all where pred() returns
// true in one slice and false in another one. Their individual ordering will be
// the same of the original one.
func Partition[V any](vs []V, pred func(V) bool) ([]V, []V) {
	var satisfying []V
	var notSatisfying []V
	for _, v := range vs {
		if pred(v) {
			satisfying = append(satisfying, v)
		} else {
			notSatisfying = append(notSatisfying, v)
		}
	}
	return satisfying, notSatisfying
}

// EOF
