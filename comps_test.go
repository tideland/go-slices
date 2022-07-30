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

	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	one := []int{1}
	none := []int{}
	in := "0"
	outAll := "0123456789"
	outOne := "01"
	outNone := "0"
	stringer := func(v int, acc string) string { return fmt.Sprintf("%s%d", acc, v) }

	assert.Equal(slices.FoldL(stringer, in, all), outAll)
	assert.Equal(slices.FoldL(stringer, in, one), outOne)
	assert.Equal(slices.FoldL(stringer, in, none), outNone)
}

// TestFoldLFirst verifies the left folding of a slice with first as
// accumulator.
func TestFoldLFirst(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	one := []int{1}
	none := []int{}
	outAll := 123456789
	outOne := 1
	outNone := 0
	potentiator := func(v, acc int) int { return acc*10 + v }

	assert.Equal(slices.FoldLFirst(potentiator, all), outAll)
	assert.Equal(slices.FoldLFirst(potentiator, one), outOne)
	assert.Equal(slices.FoldLFirst(potentiator, none), outNone)
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

// EOF
