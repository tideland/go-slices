// Tideland Go Slices
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package slices // import "tideland.dev/go/slices"

//--------------------
// CALCS
//--------------------

// FoldL iterates over the slice from left to right. It calls fun() for
// each value passing the initial accumulator. The accumulator returned
// by each function call is passed to the next one. The last one will be
// returned.
func FoldL[V, Acc any](fun func(V, Acc) Acc, acc Acc, vs []V) Acc {
	out := acc
	for _, v := range vs {
		out = fun(v, out)
	}
	return out
}

// EOF
