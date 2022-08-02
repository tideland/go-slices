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

// TestAppend verifies the appending of a number of slices.
func TestAppend(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values [][]int
		out    []int
	}{
		{
			descr:  "One slice",
			values: [][]int{{1, 2, 3}},
			out:    []int{1, 2, 3},
		}, {
			descr:  "Two slices",
			values: [][]int{{1, 2, 3}, {4, 5, 6}},
			out:    []int{1, 2, 3, 4, 5, 6},
		}, {
			descr:  "More slices",
			values: [][]int{{1, 2, 3}, {4, 5, 6}, {7}, {8}, {9, 10}},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		}, {
			descr:  "No slices",
			values: [][]int{},
			out:    []int{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Append(test.values...), test.out)
	}
}

// TestDelete verifies the deleting of a first matching value of a slice.
func TestDelete(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		value  int
		values []int
		out    []int
	}{
		{
			descr:  "Delete first",
			value:  1,
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete last",
			value:  9,
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8},
		}, {
			descr:  "Delete single value",
			value:  2,
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete one of double",
			value:  5,
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete uncontained value",
			value:  10,
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete in empty input",
			value:  1,
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Delete in nil input",
			value:  1,
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Delete(test.value, test.values), test.out)
	}
}

// TestDropWhile verifies the dropping of the slice elements as long
// as a test returns true.
func TestDropWhile(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	dropper := func(v int) bool { return v <= 5 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{6, 7, 8, 9},
		}, {
			descr:  "Longer slice without drop",
			values: []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			out:    []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		}, {
			descr:  "Single value slice with drop",
			values: []int{0},
			out:    nil,
		}, {
			descr:  "Single value slice without drop",
			values: []int{9},
			out:    []int{9},
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    nil,
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.DropWhile(dropper, test.values), test.out)
	}
}

// TestFilter verifies the filtering of slice values.
func TestFilter(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	filter := func(v int) bool { return v%2 == 0 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Many values, some filtered",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{2, 4, 6, 8},
		}, {
			descr:  "Many values, none filtered",
			values: []int{2, 4, 6, 8, 10, 12, 14, 16},
			out:    []int{2, 4, 6, 8, 10, 12, 14, 16},
		}, {
			descr:  "Many values, all filtered",
			values: []int{1, 3, 5, 7, 11, 13, 15, 17},
			out:    []int{},
		}, {
			descr:  "One value, not filtered",
			values: []int{2},
			out:    []int{2},
		}, {
			descr:  "One value, filtered",
			values: []int{1},
			out:    []int{},
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Filter(filter, test.values), test.out)
	}

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	even := []int{2, 4, 6, 8}
	none := []int{}

	assert.Equal(slices.Filter(func(v int) bool { return v%2 == 0 }, all), even)
	assert.Equal(slices.Filter(func(v int) bool { return v > 100 }, all), none)
}

// TestFilterMap verifies the filtering and mapping of slice values.
func TestFilterMap(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	evenMap := func(v int) (string, bool) {
		if v%2 == 0 {
			return fmt.Sprintf("%v", v), true
		}
		return "", false
	}
	even := []string{"2", "4", "6", "8"}
	noneMap := func(v int) (bool, bool) {
		return false, false
	}
	none := []bool{}

	assert.Equal(slices.FilterMap(evenMap, all), even)
	assert.Equal(slices.FilterMap(noneMap, all), none)
}

// TestJoin verifies the joining of a separator value and the slice values.
func TestJoin(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	sep := 0
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5},
			out:    []int{1, 0, 2, 0, 3, 0, 4, 0, 5},
		}, {
			descr:  "Single value slice",
			values: []int{0},
			out:    []int{0},
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Join(sep, test.values), test.out)
	}
}

// TestMap verifies the mapping of slice values.
func TestMap(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	mapper := func(v int) string {
		itoa := map[int]string{
			0: "zero",
			1: "one",
			2: "two",
			3: "three",
			4: "four",
			5: "five",
			6: "six",
			7: "seven",
			8: "eight",
			9: "nine",
		}
		return itoa[v]
	}
	tests := []struct {
		descr  string
		values []int
		out    []string
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5},
			out:    []string{"one", "two", "three", "four", "five"},
		}, {
			descr:  "Single value slice",
			values: []int{0},
			out:    []string{"zero"},
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    []string{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Map(mapper, test.values), test.out)
	}
}

// TestReverse verifies the reversal of slices.
func TestReverse(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5},
			out:    []int{5, 4, 3, 2, 1},
		}, {
			descr:  "Single value slice",
			values: []int{0},
			out:    []int{0},
		}, {
			descr:  "Empty slice",
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Reverse(test.values), test.out)
	}
}

// EOF
