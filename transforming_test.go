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

	first := []int{1, 2, 3}
	second := []int{4, 5, 6}
	third := []int{7, 8, 9}
	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	none := []int{}

	assert.True(slices.IsEqual(slices.Append(first, second, third), all))
	assert.True(slices.IsEqual(slices.Append(first), first))
	assert.True(slices.IsEqual(slices.Append[int](), []int{}))
	assert.True(slices.IsEqual(slices.Append(none, none, none), none))
}

// TestDelete verifies the deleting of a first matching value of a slice.
func TestDelete(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedOne := []int{2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedFive := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	deletedNine := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	none := []int{}

	assert.True(slices.IsEqual(slices.Delete(1, all), deletedOne))
	assert.True(slices.IsEqual(slices.Delete(5, all), deletedFive))
	assert.True(slices.IsEqual(slices.Delete(9, all), deletedNine))
	assert.True(slices.IsEqual(slices.Delete(10, all), all))
	assert.True(slices.IsEqual(slices.Delete(5, none), none))
}

// TestDropWhile verifies the dropping of the slice elements as long
// as a test returns true.
func TestDropWhile(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	none := []int{}
	allGtMinusOne := slices.DropWhile(func(v int) bool { return v <= -1 }, none)
	cmpGtMinusOne := []int{}
	allGtZero := slices.DropWhile(func(v int) bool { return v <= 0 }, all)
	cmpGtZero := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	allGtOne := slices.DropWhile(func(v int) bool { return v <= 1 }, all)
	cmpGtOne := []int{2, 3, 4, 5, 5, 6, 7, 8, 9}
	allGtFive := slices.DropWhile(func(v int) bool { return v <= 5 }, all)
	cmpGtFive := []int{6, 7, 8, 9}
	allGtSix := slices.DropWhile(func(v int) bool { return v <= 6 }, all)
	cmpGtSix := []int{7, 8, 9}
	allGtNine := slices.DropWhile(func(v int) bool { return v <= 9 }, all)
	cmpGtNine := []int{}
	allGtTen := slices.DropWhile(func(v int) bool { return v <= 10 }, all)
	cmpGtTen := []int{}

	assert.True(slices.IsEqual(allGtMinusOne, cmpGtMinusOne))
	assert.True(slices.IsEqual(allGtZero, cmpGtZero))
	assert.True(slices.IsEqual(allGtOne, cmpGtOne))
	assert.True(slices.IsEqual(allGtFive, cmpGtFive))
	assert.True(slices.IsEqual(allGtSix, cmpGtSix))
	assert.True(slices.IsEqual(allGtNine, cmpGtNine))
	assert.True(slices.IsEqual(allGtTen, cmpGtTen))
}

// TestFilter verifies the filtering of slice values.
func TestFilter(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	even := []int{2, 4, 6, 8}
	none := []int{}

	assert.True(slices.IsEqual(slices.Filter(func(v int) bool { return v%2 == 0 }, all), even))
	assert.True(slices.IsEqual(slices.Filter(fevenol { return v > 100 }, all), none))
}

// TestFilterMap verifies the filtering and mapping of slice values.
func TestFilterMap(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	even := []string{"2", "4", "6", "8"}
	none := []bool{}
	evenMap := func(v int) (string, bool) {
		if v%2 == 0 {
			return fmt.Sprintf("%v", v), true
		}
		return "", false
	}
	noneMap := func(v int) (bool, bool) {
		return false, false
	}

	assert.True(slices.IsEqual(slices.FilterMap(evenMap, all), even))
	assert.True(slices.IsEqual(slices.FilterMap(noneMap, all), none))
}

// EOF
