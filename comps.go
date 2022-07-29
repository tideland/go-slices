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
func FoldL[V, Acc any](fun func(V, Acc) Acc, acc Acc, vs []V) Acc {
	out := acc
	for _, v := range vs {
		out = fun(v, out)
	}
	return out
}

// FoldLFirst iterates over the slice from left to right. It calls fun() for
// each value passing the first value as accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldLFirst[V any](fun func(V, V) V, vs []V) V {
	var first V
	if len(vs) == 0 {
		// Return default value.
		return first
	}
	first = vs[0]
	return FoldL(fun, first, vs[1:])
}

// FoldR iterates over the slice from right to left. It calls fun() for
// each value passing the initial accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldR[V, Acc any](fun func(V, Acc) Acc, acc Acc, vs []V) Acc {
	out := acc
	for i := len(vs) - 1; i >= 0; i-- {
		out = fun(vs[i], out)
	}
	return out
}

// FoldRLast iterates over the slice from right to left. It calls fun() for
// each value passing the last value as accumulator. The accumulator returned
// by each function call is used as input at the next call. The last one will be
// returned.
func FoldRLast[V any](fun func(V, V) V, vs []V) V {
	var last V
	if len(vs) == 0 {
		// Return default value.
		return last
	}
	last = vs[len(vs)-1]
	return FoldR(fun, last, vs[:len(vs)-1])
}

// MapFoldL combines the operations of Map() and FoldL() in one pass.
func MapFoldL[I, O, Acc any](fun func(I, Acc) (O, Acc), acc Acc, ivs []I) ([]O, Acc) {
	var ov O
	var ovs []O
	for _, iv := range ivs {
		ov, acc = fun(iv, acc)
		ovs = append(ovs, ov)
	}
	return ovs, acc
}

// EOF