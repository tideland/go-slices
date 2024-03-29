// Tideland Go Testing - Unit Tests
//
// Copyright (C) 2022-2023 Frank Mueller / Tideland / Oldenburg / Germany
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
			out:    nil,
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

// TestContainsAll verifies the testing of all slice values.
func TestContainsAll(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	contains := func(v int) bool { return v < 10 }
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
		assert.Equal(slices.ContainsAll(test.values, contains), test.out)
	}
}

// TestContainsAny verifies the testing of all slice values finding
// at least one value.
func TestContainsAny(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	contains := func(v int) bool { return v < 3 }
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
		assert.Equal(slices.ContainsAny(test.values, contains), test.out)
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
			descr:  "Delete only value",
			value:  1,
			values: []int{1},
			out:    []int{},
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

// TestDeleteAll verifies the deleting of all matching values of a slice.
func TestDeleteAll(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	value := 5
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Delete single value",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 6, 7, 8, 9},
		}, {
			descr:  "Delete multiple values",
			values: []int{1, 2, 5, 3, 4, 5, 5, 5, 5, 6, 7, 5, 8, 9},
			out:    []int{1, 2, 3, 4, 6, 7, 8, 9},
		}, {
			descr:  "Delete uncontained values",
			values: []int{1, 2, 3, 4, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 6, 7, 8, 9},
		}, {
			descr:  "Delete all values",
			values: []int{5, 5, 5, 5, 5},
			out:    []int{},
		}, {
			descr:  "Delete in empty input",
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Delete in nil input",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.DeleteAll(value, test.values), test.out)
	}
}

// TestDeleteAllWith verifies the deleting of slice values.
func TestDeleteAllWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	deleteWith := func(v int) bool { return v%2 == 0 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Many values, some deleted",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{1, 3, 5, 7, 9},
		}, {
			descr:  "Many values, none deleted",
			values: []int{1, 3, 5, 7, 9, 11, 13, 15},
			out:    []int{1, 3, 5, 7, 9, 11, 13, 15},
		}, {
			descr:  "Many values, all deleted",
			values: []int{2, 4, 6, 8, 10, 12, 14, 16},
			out:    []int{},
		}, {
			descr:  "One value, not deleted",
			values: []int{1},
			out:    []int{1},
		}, {
			descr:  "One value, deleted",
			values: []int{2},
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
		assert.Equal(slices.DeleteAllWith(test.values, deleteWith), test.out)
	}
}

// TestDeleteFirst verifies the deleting of the first slice value.
func TestDeleteFirst(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5},
			out:    []int{2, 3, 4, 5},
		}, {
			descr:  "Single value slice",
			values: []int{1},
			out:    nil,
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
		assert.Equal(slices.DeleteFirst(test.values), test.out)
	}
}

// TestDeleteLast verifies the deleting of the last slice value.
func TestDeleteLast(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice",
			values: []int{1, 2, 3, 4, 5},
			out:    []int{1, 2, 3, 4},
		}, {
			descr:  "Single value slice",
			values: []int{1},
			out:    nil,
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
		assert.Equal(slices.DeleteLast(test.values), test.out)
	}
}

// TestDeleteWhile verifies the deleting of the slice elements as long
// as a test returns true.
func TestDeleteWhile(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	shallDrop := func(v int) bool { return v <= 5 }
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
		assert.Equal(slices.DeleteWhile(test.values, shallDrop), test.out)
	}
}

// TestDeleteWith verifies the deleting of a first matching value of a
// slice where pred returns true.
func TestDeleteWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	shallDelete := func(v int) bool { return v == 10 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Delete first",
			values: []int{10, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete last",
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete single value",
			values: []int{1, 2, 3, 4, 5, 10, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete one of double",
			values: []int{1, 2, 3, 4, 10, 10, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 10, 6, 7, 8, 9},
		}, {
			descr:  "Delete uncontained value",
			values: []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9},
		}, {
			descr:  "Delete only value",
			values: []int{10},
			out:    []int{},
		}, {
			descr:  "Delete in empty input",
			values: []int{},
			out:    []int{},
		}, {
			descr:  "Delete in nil input",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.DeleteWith(test.values, shallDelete), test.out)
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
		assert.Equal(slices.Filter(test.values, filter), test.out)
	}
}

// TestFilterMap verifies the filtering and mapping of slice values.
func TestFilterMap(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	filterMap := func(v int) (int, bool) { return v * 10, v%2 == 0 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Many values, some filtered, rest mapped",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{20, 40, 60, 80},
		}, {
			descr:  "Many values, none filtered, all mapped",
			values: []int{2, 4, 6, 8, 10},
			out:    []int{20, 40, 60, 80, 100},
		}, {
			descr:  "Many values, all filtered",
			values: []int{1, 3, 5, 7, 11},
			out:    []int{},
		}, {
			descr:  "One value, not filtered but mapped",
			values: []int{2},
			out:    []int{20},
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
		assert.Equal(slices.FilterMap(test.values, filterMap), test.out)
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
		assert.Equal(slices.Map(test.values, mapper), test.out)
	}
}

// TestMerge verifies the sorted merging of slices.
func TestMerge(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr   string
		valuesA []int
		valuesB []int
		out     []int
	}{
		{
			descr:   "Overlapping slices",
			valuesA: []int{5, 4, 3, 2, 1},
			valuesB: []int{3, 4, 5, 6, 7},
			out:     []int{1, 2, 3, 3, 4, 4, 5, 5, 6, 7},
		}, {
			descr:   "Individual slices",
			valuesA: []int{1, 2, 3, 4, 5},
			valuesB: []int{9, 8, 7, 6},
			out:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:   "Empty first slice",
			valuesA: []int{},
			valuesB: []int{3, 2, 3, 1},
			out:     []int{1, 2, 3, 3},
		}, {
			descr:   "Empty second slice",
			valuesA: []int{3, 2, 3, 1},
			valuesB: []int{},
			out:     []int{1, 2, 3, 3},
		}, {
			descr:   "Both slices empty",
			valuesA: []int{},
			valuesB: []int{},
			out:     nil,
		}, {
			descr:   "First slice nil",
			valuesA: nil,
			valuesB: []int{3, 2, 3, 1},
			out:     []int{1, 2, 3, 3},
		}, {
			descr:   "Second slice nil",
			valuesA: []int{3, 2, 3, 1},
			valuesB: nil,
			out:     []int{1, 2, 3, 3},
		}, {
			descr:   "Nil slices",
			valuesA: nil,
			valuesB: nil,
			out:     nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Merge(test.valuesA, test.valuesB), test.out)
	}
}

// TestMergeWith verifies the sorted merging of slices.
func TestMergeWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	type V struct {
		k int
		v string
	}

	key := func(v V) int {
		return v.k
	}
	tests := []struct {
		descr   string
		valuesA []V
		valuesB []V
		out     []V
	}{
		{
			descr:   "Overlapping slices",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}, {2, "two"}, {1, "one"}},
			valuesB: []V{{7, "seven"}, {6, "six"}, {5, "five"}, {4, "four"}},
			out: []V{{1, "one"}, {2, "two"}, {3, "three"}, {4, "four"}, {4, "four"},
				{5, "five"}, {5, "five"}, {6, "six"}, {7, "seven"}},
		}, {
			descr:   "Individual slices",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {3, "three"}, {4, "four"}, {5, "five"}, {6, "six"}},
		}, {
			descr:   "Empty first slice",
			valuesA: []V{},
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {6, "six"}},
		}, {
			descr:   "Empty second slice",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: []V{},
			out:     []V{{3, "three"}, {4, "four"}, {5, "five"}},
		}, {
			descr:   "Both slices empty",
			valuesA: []V{},
			valuesB: []V{},
			out:     nil,
		}, {
			descr:   "First slice nil",
			valuesA: nil,
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {6, "six"}},
		}, {
			descr:   "Second slice nil",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: nil,
			out:     []V{{3, "three"}, {4, "four"}, {5, "five"}},
		}, {
			descr:   "Nil slices",
			valuesA: nil,
			valuesB: nil,
			out:     nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.MergeWith(test.valuesA, test.valuesB, key), test.out)
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

// TestSplit veriefies the splitting of slices into two parts.
func TestSplit(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		n      int
		values []int
		lout   []int
		rout   []int
	}{
		{
			descr:  "Split slice in the middle",
			n:      2,
			values: []int{1, 2, 3, 4, 5},
			lout:   []int{1, 2, 3},
			rout:   []int{4, 5},
		}, {
			descr:  "Split slice in the beginning",
			n:      0,
			values: []int{1, 2, 3, 4, 5},
			lout:   []int{1},
			rout:   []int{2, 3, 4, 5},
		}, {
			descr:  "Split slice in the end",
			n:      4,
			values: []int{1, 2, 3, 4, 5},
			lout:   []int{1, 2, 3, 4, 5},
			rout:   nil,
		}, {
			descr:  "Split a single element slice",
			n:      0,
			values: []int{1},
			lout:   []int{1},
			rout:   nil,
		}, {
			descr:  "Split an empty slice",
			n:      0,
			values: []int{},
			lout:   []int{},
			rout:   nil,
		}, {
			descr:  "Split a nil slice",
			n:      0,
			values: nil,
			lout:   nil,
			rout:   nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		lout, rout := slices.Split(test.n, test.values)
		assert.Equal(lout, test.lout)
		assert.Equal(rout, test.rout)
	}
}

// TestSplitWith veriefies the splitting of slices into two parts based
// on value testing.
func TestSplitWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		pred   func(int) bool
		values []int
		lout   []int
		rout   []int
	}{
		{
			descr:  "Split slice in the middle",
			pred:   func(v int) bool { return v < 4 },
			values: []int{1, 2, 3, 4, 5},
			lout:   []int{1, 2, 3},
			rout:   []int{4, 5},
		}, {
			descr:  "Split slice in the beginning",
			pred:   func(v int) bool { return v != 1 },
			values: []int{1, 3, 2, 4, 5, 1, 2},
			lout:   nil,
			rout:   []int{1, 3, 2, 4, 5, 1, 2},
		}, {
			descr:  "Split slice in the end",
			pred:   func(v int) bool { return v < 6 },
			values: []int{1, 2, 3, 4, 5},
			lout:   []int{1, 2, 3, 4, 5},
			rout:   nil,
		}, {
			descr:  "Split a single element slice",
			pred:   func(v int) bool { return v == 1 },
			values: []int{1},
			lout:   []int{1},
			rout:   nil,
		}, {
			descr:  "Split an empty slice",
			pred:   func(v int) bool { return true },
			values: []int{},
			lout:   nil,
			rout:   nil,
		}, {
			descr:  "Split a nil slice",
			pred:   func(v int) bool { return true },
			values: nil,
			lout:   nil,
			rout:   nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		lout, rout := slices.SplitWith(test.values, test.pred)
		assert.Equal(lout, test.lout)
		assert.Equal(rout, test.rout)
	}
}

// TestSubslice verifies the convenient retrieval of sections of slices into
// new created slices.
func TestSubsclice(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr      string
		values     []int
		fpos, tpos int
		out        []int
	}{
		{
			descr:  "Copy total slice",
			values: []int{1, 2, 3, 4, 5},
			fpos:   0,
			tpos:   4,
			out:    []int{1, 2, 3, 4, 5},
		}, {
			descr:  "Copy subslice from insert of slice",
			values: []int{1, 2, 3, 4, 5},
			fpos:   1,
			tpos:   3,
			out:    []int{2, 3, 4},
		}, {
			descr:  "Copy from beginning",
			values: []int{1, 2, 3, 4, 5},
			fpos:   0,
			tpos:   2,
			out:    []int{1, 2, 3},
		}, {
			descr:  "Copy till ending",
			values: []int{1, 2, 3, 4, 5},
			fpos:   2,
			tpos:   4,
			out:    []int{3, 4, 5},
		}, {
			descr:  "Copy from before beginning",
			values: []int{1, 2, 3, 4, 5},
			fpos:   -9,
			tpos:   2,
			out:    []int{1, 2, 3},
		}, {
			descr:  "Copy till after ending",
			values: []int{1, 2, 3, 4, 5},
			fpos:   2,
			tpos:   9,
			out:    []int{3, 4, 5},
		}, {
			descr:  "Copy with fpos higher than tpos",
			values: []int{1, 2, 3, 4, 5},
			fpos:   3,
			tpos:   1,
			out:    nil,
		}, {
			descr:  "Copy from before slice",
			values: []int{1, 2, 3, 4, 5},
			fpos:   -9,
			tpos:   -1,
			out:    nil,
		}, {
			descr:  "Copy from behind slice",
			values: []int{1, 2, 3, 4, 5},
			fpos:   11,
			tpos:   13,
			out:    nil,
		}, {
			descr:  "Copy from empty slice",
			values: []int{},
			fpos:   1,
			tpos:   3,
			out:    nil,
		}, {
			descr:  "Copy from nil slice",
			values: nil,
			fpos:   1,
			tpos:   3,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Subslice(test.values, test.fpos, test.tpos), test.out)
	}
}

// TestSubtract verifies the subtracting of values froma a slice.
func TestSubtract(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr    string
		values   []int
		subtract []int
		out      []int
	}{
		{
			descr:    "Subtract of random and multiple values",
			values:   []int{1, 2, 3, 2, 4, 5, 1, 2},
			subtract: []int{2, 1, 2},
			out:      []int{3, 4, 5, 1, 2},
		}, {
			descr:    "Subtract of first values",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			subtract: []int{3, 2, 1},
			out:      []int{4, 5, 6, 7, 8, 9},
		}, {
			descr:    "Subtract of last values",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			subtract: []int{7, 8, 9},
			out:      []int{1, 2, 3, 4, 5, 6},
		}, {
			descr:    "Subtract value more than existing",
			values:   []int{1, 2, 3, 4, 3, 2, 1, 2, 3},
			subtract: []int{1, 1, 1, 1},
			out:      []int{2, 3, 4, 3, 2, 2, 3},
		}, {
			descr:    "Subtract of not existing values",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			subtract: []int{0, 0, 0},
			out:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:    "Subtract of an empty slice",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			subtract: []int{},
			out:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:    "Subtract of a nil slice",
			values:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			subtract: nil,
			out:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:    "Subtract from an empty slice",
			values:   []int{},
			subtract: []int{3, 2, 1},
			out:      []int{},
		}, {
			descr:    "Subtract from a nil slice",
			values:   nil,
			subtract: []int{3, 2, 1},
			out:      nil,
		}, {
			descr:    "Subtract nil from a nil slice",
			values:   nil,
			subtract: nil,
			out:      nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.Subtract(test.values, test.subtract), test.out)
	}
}

// TestTakeWhile verifies the copying of the slice values as long
// as a test returns true.
func TestTakeWhile(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	shallTake := func(v int) bool { return v <= 5 }
	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice with one stopper",
			values: []int{1, 2, 3, 4, 5, 6, 5, 4, 3},
			out:    []int{1, 2, 3, 4, 5},
		}, {
			descr:  "Longer slice without taking",
			values: []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			out:    nil,
		}, {
			descr:  "Single value slice with taking",
			values: []int{0},
			out:    []int{0},
		}, {
			descr:  "Single value slice without taking",
			values: []int{9},
			out:    nil,
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
		assert.Equal(slices.TakeWhile(test.values, shallTake), test.out)
	}
}

// TestUnique verifies the unifying of a slice.
func TestUnique(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Longer slice with one double value",
			values: []int{1, 2, 3, 4, 5, 6, 5, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Longer slice with one multiple time value",
			values: []int{1, 2, 5, 3, 5, 4, 5, 6, 5, 7, 8, 9},
			out:    []int{1, 2, 5, 3, 4, 6, 7, 8, 9},
		}, {
			descr:  "Longer slice with multiple multiple time values",
			values: []int{1, 2, 5, 3, 4, 5, 4, 5, 6, 6, 5, 7, 7, 8, 7, 9},
			out:    []int{1, 2, 5, 3, 4, 6, 7, 8, 9},
		}, {
			descr:  "Longer slice without any double value",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Longer slice only with double values",
			values: []int{1, 1, 1, 1, 1, 1, 1, 1, 1},
			out:    []int{1},
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
		assert.Equal(slices.Unique(test.values), test.out)
	}
}

// TestUniqueMerge verifies the unique sorted merging of slices.
func TestUniqueMerge(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr   string
		valuesA []int
		valuesB []int
		out     []int
	}{
		{
			descr:   "Overlapping slices",
			valuesA: []int{5, 4, 3, 2, 1},
			valuesB: []int{3, 4, 5, 6, 7},
			out:     []int{1, 2, 3, 4, 5, 6, 7},
		}, {
			descr:   "Individual slices",
			valuesA: []int{1, 2, 3, 4, 5},
			valuesB: []int{9, 8, 7, 6},
			out:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:   "Empty first slice",
			valuesA: []int{},
			valuesB: []int{3, 2, 3, 1},
			out:     []int{1, 2, 3},
		}, {
			descr:   "Empty second slice",
			valuesA: []int{3, 2, 3, 1},
			valuesB: []int{},
			out:     []int{1, 2, 3},
		}, {
			descr:   "Both slices empty",
			valuesA: []int{},
			valuesB: []int{},
			out:     nil,
		}, {
			descr:   "First slice nil",
			valuesA: nil,
			valuesB: []int{3, 2, 3, 1},
			out:     []int{1, 2, 3},
		}, {
			descr:   "Second slice nil",
			valuesA: []int{3, 2, 3, 1},
			valuesB: nil,
			out:     []int{1, 2, 3},
		}, {
			descr:   "Nil slices",
			valuesA: nil,
			valuesB: nil,
			out:     nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.UniqueMerge(test.valuesA, test.valuesB), test.out)
	}
}

// TestUniqueMergeWith verifies the unique sorted merging of slice.
func TestUniqueMergeWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	type V struct {
		k int
		v string
	}

	key := func(v V) int {
		return v.k
	}
	tests := []struct {
		descr   string
		valuesA []V
		valuesB []V
		out     []V
	}{
		{
			descr:   "Overlapping slices",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}, {2, "two"}, {1, "one"}},
			valuesB: []V{{7, "seven"}, {6, "six"}, {5, "five"}, {4, "four"}},
			out:     []V{{1, "one"}, {2, "two"}, {3, "three"}, {4, "four"}, {5, "five"}, {6, "six"}, {7, "seven"}},
		}, {
			descr:   "Individual slices",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {3, "three"}, {4, "four"}, {5, "five"}, {6, "six"}},
		}, {
			descr:   "Empty first slice",
			valuesA: []V{},
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {6, "six"}},
		}, {
			descr:   "Empty second slice",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: []V{},
			out:     []V{{3, "three"}, {4, "four"}, {5, "five"}},
		}, {
			descr:   "Both slices empty",
			valuesA: []V{},
			valuesB: []V{},
			out:     nil,
		}, {
			descr:   "First slice nil",
			valuesA: nil,
			valuesB: []V{{2, "two"}, {1, "one"}, {6, "six"}},
			out:     []V{{1, "one"}, {2, "two"}, {6, "six"}},
		}, {
			descr:   "Second slice nil",
			valuesA: []V{{5, "five"}, {4, "four"}, {3, "three"}},
			valuesB: nil,
			out:     []V{{3, "three"}, {4, "four"}, {5, "five"}},
		}, {
			descr:   "Nil slices",
			valuesA: nil,
			valuesB: nil,
			out:     nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.UniqueMergeWith(test.valuesA, test.valuesB, key), test.out)
	}
}

// TestUnique verifies the unifying of a slice.
func TestUniqueWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	type V struct {
		k string
		v int
	}

	key := func(v V) string {
		return v.k
	}
	tests := []struct {
		descr  string
		values []V
		out    []V
	}{
		{
			descr:  "Longer slice with one double value",
			values: []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"three", 3}},
			out:    []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}},
		}, {
			descr:  "Longer slice with one multiple time value",
			values: []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"three", 3}, {"three", 3}},
			out:    []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}},
		}, {
			descr: "Longer slice with multiple multiple time values",
			values: []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"three", 3}, {"three", 3},
				{"two", 2}, {"three", 3}, {"four", 4}, {"one", 1}},
			out: []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}},
		}, {
			descr:  "Longer slice without any double values",
			values: []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"five", 5}},
			out:    []V{{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"five", 5}},
		}, {
			descr:  "Longer slice ony with double values/keys",
			values: []V{{"one", 1}, {"one", 2}, {"one", 3}, {"one", 4}, {"one", 5}},
			out:    []V{{"one", 1}},
		}, {
			descr:  "Single value slice",
			values: []V{{"one", 1}},
			out:    []V{{"one", 1}},
		}, {
			descr:  "Empty slice",
			values: []V{},
			out:    []V{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.UniqueWith(test.values, key), test.out)
	}
}

// TestSearch verifies the search in a slice.
func TestSearch(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	pred := func(v int) bool {
		return v == 3
	}
	tests := []struct {
		descr  string
		values []int
		out    int
		ok     bool
	}{
		{
			descr:  "Search in empty slice",
			values: []int{},
			out:    0,
			ok:     false,
		}, {
			descr:  "Search in nil slice",
			values: nil,
			out:    0,
			ok:     false,
		}, {
			descr:  "Search in slice with one value",
			values: []int{3},
			out:    3,
			ok:     true,
		}, {
			descr:  "Search in slice with one value not found",
			values: []int{1},
			out:    0,
			ok:     false,
		}, {
			descr:  "Search in slice with multiple values",
			values: []int{1, 2, 3, 4, 5},
			out:    3,
			ok:     true,
		}, {
			descr:  "Search in slice with multiple values not found",
			values: []int{1, 2, 4, 5},
			out:    0,
			ok:     false,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		v, ok := slices.Search(pred, test.values)
		assert.Equal(v, test.out)
		assert.Equal(ok, test.ok)
	}
}

// EOF
