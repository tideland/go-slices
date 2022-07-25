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

// TestContainsAll verifies the testing of all slice values.
func TestContainsAll(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fs := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	assert.True(slices.ContainsAll[int](func(i int) bool {
		return i > 0
	}, is))
	assert.True(slices.ContainsAll(func(i int) bool {
		return i > 0
	}, is))
	assert.False(slices.ContainsAll(func(i int) bool {
		return i != 5
	}, is))
	assert.True(slices.ContainsAll(func(f float64) bool {
		return f > 0.0
	}, fs))
}

// TestContainsAny verifies the testing of all slice values finding
// at least one value.
func TestContainsAny(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fs := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	assert.True(slices.ContainsAny[int](func(i int) bool {
		return i == 5
	}, is))
	assert.True(slices.ContainsAny(func(i int) bool {
		return i == 9
	}, is))
	assert.False(slices.ContainsAny(func(i int) bool {
		return i > 10
	}, is))
	assert.True(slices.ContainsAny(func(f float64) bool {
		return f > 5.0
	}, fs))
}

// TestIsEqual verifies the testing of the equality of two slices.
func TestIsEqual(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	isA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isB := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isC := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assert.True(slices.IsEqual(isA, isB))
	assert.False(slices.IsEqual(isA, isC))
	assert.False(slices.IsEqual(isC, isA))
}

// TestIsMember verifies the testing of the membership of a value
// in a slice.
func TestIsMember(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	is := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fs := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	assert.True(slices.IsMember[int](5, is))
	assert.True(slices.IsMember(9, is))
	assert.False(slices.IsMember(100, is))
	assert.True(slices.IsMember(5.0, fs))
}

// TestIsPrefix verifies the testing of prefix slices.
func TestIsPrefix(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	prefixA := []int{1, 2, 3, 4, 5}
	prefixB := []int{1, 2, 3, 99, 5}
	prefixC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.True(slices.IsPrefix(prefixA, all))
	assert.True(slices.IsPrefix(all, all))
	assert.False(slices.IsPrefix(prefixB, all))
	assert.False(slices.IsPrefix(prefixC, all))
}

// TestIsSuffix verifies the testing of suffix slices.
func TestIsSuffix(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	suffixA := []int{6, 7, 8, 9, 10}
	suffixB := []int{6, 7, 8, 99, 10}
	suffixC := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.True(slices.IsSuffix(suffixA, all))
	assert.True(slices.IsSuffix(all, all))
	assert.False(slices.IsSuffix(suffixB, all))
	assert.False(slices.IsSuffix(suffixC, all))
}

// EOF
