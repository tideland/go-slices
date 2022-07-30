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
	"fmt"
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

	acc := "0"
	stringer := func(v int, acc string) string { return fmt.Sprintf("%s%d", acc, v) }
	tests := []struct {
		values []int
		out    string
	}{
		{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    "0123456789",
		}, {
			values: []int{1},
			out:    "01",
		}, {
			values: []int{},
			out:    "0",
		}, {
			values: nil,
			out:    "0",
		},
	}

	for _, test := range tests {
		assert.Equal(slices.FoldL(stringer, acc, test.values), test.out)
	}
}

// TestFoldLFirst verifies the left folding of a slice with first as
// accumulator.
func TestFoldLFirst(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	potentiator := func(v, acc int) int { return acc*10 + v }
	tests := []struct {
		values []int
		out    int
	}{
		{
			values: []int{1, 2, 3, 4, 5},
			out:    12345,
		}, {
			values: []int{1},
			out:    1,
		}, {
			values: []int{},
			out:    0,
		}, {
			values: nil,
			out:    0,
		},
	}

	for _, test := range tests {
		assert.Equal(slices.FoldLFirst(potentiator, test.values), test.out)
	}
}

// TestFoldR verifies the right folding of a slice.
func TestFoldR(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	acc := "0"
	stringer := func(v int, acc string) string { return fmt.Sprintf("%s%d", acc, v) }
	tests := []struct {
		values []int
		out    string
	}{
		{
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    "0987654321",
		}, {
			values: []int{1},
			out:    "01",
		}, {
			values: []int{},
			out:    "0",
		}, {
			values: nil,
			out:    "0",
		},
	}

	for _, test := range tests {
		assert.Equal(slices.FoldR(stringer, acc, test.values), test.out)
	}
}

// TestFoldRLast verifies the right folding of a slice.
func TestFoldRLast(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	potentiator := func(v, acc int) int { return acc*10 + v }
	tests := []struct {
		values []int
		out    int
	}{
		{
			values: []int{1, 2, 3, 4, 5},
			out:    54321,
		}, {
			values: []int{1},
			out:    1,
		}, {
			values: []int{},
			out:    0,
		}, {
			values: nil,
			out:    0,
		},
	}

	for _, test := range tests {
		assert.Equal(slices.FoldRLast(potentiator, test.values), test.out)
	}
}

// TestMapFoldL verifies the left combined mapping and folding.
func MapFoldL(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	in := "0"
	plusStringer := func(v int, acc string) (int, string) { return v + 1, fmt.Sprintf("%s%d", acc, v) }
	tests := []struct {
		values []int
		mapped []int
		out    string
	}{
		{
			values: []int{1, 2, 3, 4, 5},
			mapped: []int{2, 3, 4, 5, 6},
			out:    "012345",
		}, {
			values: []int{1},
			mapped: []int{2},
			out:    "01",
		}, {
			values: []int{},
			mapped: []int{},
			out:    "0",
		},
	}

	for _, test := range tests {
		mapped, out := slices.MapFoldL(plusStringer, in, test.values)
		assert.Equal(mapped, test.mapped)
		assert.Equal(out, test.out)
	}
}

// TestMapFoldR verifies the right combined mapping and folding.
func TestMapFoldR(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	in := "0"
	plusStringer := func(v int, acc string) (int, string) { return v + 1, fmt.Sprintf("%s%d", acc, v) }
	tests := []struct {
		values []int
		mapped []int
		out    string
	}{
		{
			values: []int{1, 2, 3, 4, 5},
			mapped: []int{2, 3, 4, 5, 6},
			out:    "054321",
		}, {
			values: []int{1},
			mapped: []int{2},
			out:    "01",
		}, {
			values: []int{},
			mapped: []int{},
			out:    "0",
		}, {
			values: nil,
			mapped: nil,
			out:    "0",
		},
	}

	for _, test := range tests {
		mapped, out := slices.MapFoldR(plusStringer, in, test.values)
		assert.Equal(mapped, test.mapped)
		assert.Equal(out, test.out)
	}
}

// TestPartition verifies the partitioning of slices.
func TestPartition(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	isMod := func(v int) bool { return v%2 == 0 }
	tests := []struct {
		values        []int
		satisfying    []int
		notSatisfying []int
	}{
		{
			values:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			satisfying:    []int{2, 4, 6, 8},
			notSatisfying: []int{1, 3, 5, 7, 9},
		}, {
			values:        []int{1},
			satisfying:    nil,
			notSatisfying: []int{1},
		}, {
			values:        []int{},
			satisfying:    nil,
			notSatisfying: nil,
		}, {
			values:        nil,
			satisfying:    nil,
			notSatisfying: nil,
		},
	}

	for _, test := range tests {
		satisfying, notSatisfying := slices.Partition(isMod, test.values)
		assert.Equal(satisfying, test.satisfying)
		assert.Equal(notSatisfying, test.notSatisfying)
	}
}

// EOF
