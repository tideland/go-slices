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

	container := func(v int) bool { return v < 10 }
	tests := []struct {
		descr  string
		values []int
		out    bool
	}{
		{
			descr:  "Predicate matches all values",
			values: []int{1, 2, 3, 4, 5},
			out:    true,
		}, {
			descr:  "Predicate not matches all values",
			values: []int{8, 9, 10, 11, 12, 13, 14, 15},
			out:    false,
		}, {
			descr:  "Single value slice matches",
			values: []int{5},
			out:    true,
		}, {
			descr:  "Single value slice does not match",
			values: []int{15},
			out:    false,
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    true,
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    true,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.ContainsAll(container, test.values), test.out)
	}
}

// TestContainsAny verifies the testing of all slice values finding
// at least one value.
func TestContainsAny(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	container := func(v int) bool { return v < 3 }
	tests := []struct {
		descr  string
		values []int
		out    bool
	}{
		{
			descr:  "Predicate matches values",
			values: []int{3, 4, 5, 1, 6, 7, 8, 2},
			out:    true,
		}, {
			descr:  "Predicate not matches any values",
			values: []int{8, 9, 10, 11, 12, 13, 14, 15},
			out:    false,
		}, {
			descr:  "Single value slice matches",
			values: []int{1},
			out:    true,
		}, {
			descr:  "Single value slice does not match",
			values: []int{5},
			out:    false,
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    false,
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    false,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.ContainsAny(container, test.values), test.out)
	}
}

// TestIsEqual verifies the testing of the equality of two slices.
func TestIsEqual(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr         string
		first, second []int
		out           bool
	}{
		{
			descr:  "First and second equal",
			first:  []int{1, 2, 3, 4, 5},
			second: []int{1, 2, 3, 4, 5},
			out:    true,
		}, {
			descr:  "First and second empty",
			first:  []int{},
			second: []int{},
			out:    true,
		}, {
			descr:  "First and second nil",
			first:  nil,
			second: nil,
			out:    true,
		}, {
			descr:  "First and second different in size",
			first:  []int{1, 2, 3, 4, 5},
			second: []int{1, 2, 3},
			out:    false,
		}, {
			descr:  "First and second different in one value",
			first:  []int{1, 2, 3, 4, 5},
			second: []int{1, 2, 3, 5, 5},
			out:    false,
		}, {
			descr:  "One filled and one empty",
			first:  []int{1, 2, 3},
			second: []int{},
			out:    false,
		}, {
			descr:  "One filled and one nil",
			first:  []int{1, 2, 3},
			second: nil,
			out:    false,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.IsEqual(test.first, test.second), test.out)
		assert.Equal(slices.IsEqual(test.second, test.first), test.out)
	}
}

// TestIsMember verifies the testing of the membership of a value
// in a slice.
func TestIsMember(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		value  int
		out    bool
	}{
		{
			descr:  "Slice contains value",
			values: []int{1, 2, 3, 4, 5},
			value:  3,
			out:    true,
		}, {
			descr:  "Slice does not contain value",
			values: []int{1, 2, 3, 4, 5},
			value:  10,
			out:    false,
		}, {
			descr:  "Empty slice does not contain value",
			values: []int{},
			value:  10,
			out:    false,
		}, {
			descr:  "Nil slice does not contain value",
			values: nil,
			value:  10,
			out:    false,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.IsMember(test.value, test.values), test.out)
	}
}

// TestIsPrefix verifies the testing of prefix slices.
func TestIsPrefix(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		prefix []int
		out    bool
	}{
		{
			descr:  "Slice has prefix",
			values: []int{1, 2, 3, 4, 5},
			prefix: []int{1, 2, 3},
			out:    true,
		}, {
			descr:  "Slice and prefix are idendtical",
			values: []int{1, 2, 3, 4, 5},
			prefix: []int{1, 2, 3, 4, 5},
			out:    true,
		}, {
			descr:  "Prefix does not match",
			values: []int{1, 2, 3, 4, 5},
			prefix: []int{1, 2, 4},
			out:    false,
		}, {
			descr:  "Prefix is too long",
			values: []int{1, 2, 3, 4, 5},
			prefix: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    false,
		}, {
			descr:  "Prefix is empty slice",
			values: []int{1, 2, 3, 4, 5},
			prefix: []int{},
			out:    true,
		}, {
			descr:  "Both are empty slices",
			values: []int{},
			prefix: []int{},
			out:    true,
		}, {
			descr:  "Prefix is nil",
			values: []int{1, 2, 3, 4, 5},
			prefix: nil,
			out:    true,
		}, {
			descr:  "Slice and prefix are nil",
			values: nil,
			prefix: nil,
			out:    true,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.IsPrefix(test.prefix, test.values), test.out)
	}
}

// TestIsSuffix verifies the testing of suffix slices.
func TestIsSuffix(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		suffix []int
		out    bool
	}{
		{
			descr:  "Slice has suffix",
			values: []int{1, 2, 3, 4, 5},
			suffix: []int{3, 4, 5},
			out:    true,
		}, {
			descr:  "slice and suffix are identical",
			values: []int{1, 2, 3, 4, 5},
			suffix: []int{1, 2, 3, 4, 5},
			out:    true,
		}, {
			descr:  "Suffix does not match",
			values: []int{1, 2, 3, 4, 5},
			suffix: []int{3, 4, 6},
			out:    false,
		}, {
			descr:  "Suffix is too long",
			values: []int{1, 2, 3, 4, 5},
			suffix: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    false,
		}, {
			descr:  "Suffix is empty slice",
			values: []int{1, 2, 3, 4, 5},
			suffix: []int{},
			out:    true,
		}, {
			descr:  "Both are empty slices",
			values: []int{},
			suffix: []int{},
			out:    true,
		}, {
			descr:  "Suffix is nil",
			values: []int{1, 2, 3, 4, 5},
			suffix: nil,
			out:    true,
		}, {
			descr:  "Slice and suffix are nil",
			values: nil,
			suffix: nil,
			out:    true,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.IsSuffix(test.suffix, test.values), test.out)
	}
}

// TestSearch verifies the search inside a slice.
func TestSearch(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	searcher := func(vs []int) bool { return vs[0] == 5 }
	tests := []struct {
		descr    string
		values   [][]int
		outValue []int
		outOK    bool
	}{
		{
			descr:    "Many values containg one hit",
			values:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 4}, {6, 6}},
			outValue: []int{5, 4},
			outOK:    true,
		}, {
			descr:    "Many values containg multiple hits",
			values:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 5}, {5, 4}, {5, 5}, {6, 6}},
			outValue: []int{5, 4},
			outOK:    true,
		}, {
			descr:    "Many values containg no hit",
			values:   [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 5}, {6, 6}},
			outValue: nil,
			outOK:    false,
		}, {
			descr:    "One value containg one hit",
			values:   [][]int{{5, 6}},
			outValue: []int{5, 6},
			outOK:    true,
		}, {
			descr:    "One value containg no hit",
			values:   [][]int{{6, 5}},
			outValue: nil,
			outOK:    false,
		}, {
			descr:    "Empty slice",
			values:   [][]int{},
			outValue: nil,
			outOK:    false,
		}, {
			descr:    "Nil slice",
			values:   nil,
			outValue: nil,
			outOK:    false,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		v, ok := slices.Search(searcher, test.values)
		assert.Equal(v, test.outValue)
		assert.Equal(ok, test.outOK)
	}
}

// EOF
