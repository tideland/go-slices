// Tideland Go Testing - Unit Tests
//
// Copyright (C) 2022 Frank Mueller / Tideland / Oldenburg / Germany
//
// MatchesAll rights reserved. Use of this source code is governed
// by the new BSD license.

package slices_test // import "tideland.dev/go/slices"

//--------------------
// IMPORTS
//--------------------

import (
	"testing"

	"tideland.dev/go/audit/asserts"

	"tideland.dev/go/slices"
)

//--------------------
// TESTS
//--------------------

// TestFoldL verifies the left folding of a slice.
func TestFoldL(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	in := 0
	out := 45

	assert.Equal(slices.FoldL(func(v int, acc int) int {
		return acc + v
	}, in, all), out)
}

// EOF
