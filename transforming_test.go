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

// TestAppend verifies the appending of a number of slices.
func TestAppend(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	assert.True(slices.IsEqual(slices.Append(first, second, third), all))
	assert.True(slices.IsEqual(slices.Append(first), first))
	assert.True(slices.IsEqual(slices.Append[int](), []int{}))
}

// TestDelete verifies the deleting of a first matching value of a slice.
func TestDelete(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedOne := []int{2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedFive := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	deletedNine := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}

	assert.True(slices.IsEqual(slices.Delete(1, all), deletedOne))
	assert.True(slices.IsEqual(slices.Delete(5, all), deletedFive))
	assert.True(slices.IsEqual(slices.Delete(9, all), deletedNine))
	assert.True(slices.IsEqual(slices.Delete(10, all), all))
}

// EOF
