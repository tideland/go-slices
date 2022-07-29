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

	assert.Equal(slices.Append(first, second, third), all)
	assert.Equal(slices.Append(first), first)
	assert.Equal(slices.Append[int](), []int{})
	assert.Equal(slices.Append(none, none, none), none)
}

// TestDelete verifies the deleting of a first matching value of a slice.
func TestDelete(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedOne := []int{2, 3, 4, 5, 5, 6, 7, 8, 9}
	deletedFive := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	deletedNine := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}
	none := []int{}

	assert.Equal(slices.Delete(1, all), deletedOne)
	assert.Equal(slices.Delete(5, all), deletedFive)
	assert.Equal(slices.Delete(9, all), deletedNine)
	assert.Equal(slices.Delete(10, all), all)
	assert.Equal(slices.Delete(5, none), none)
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

	assert.Equal(allGtMinusOne, cmpGtMinusOne)
	assert.Equal(allGtZero, cmpGtZero)
	assert.Equal(allGtOne, cmpGtOne)
	assert.Equal(allGtFive, cmpGtFive)
	assert.Equal(allGtSix, cmpGtSix)
	assert.Equal(allGtNine, cmpGtNine)
	assert.Equal(allGtTen, cmpGtTen)
}

// TestFilter verifies the filtering of slice values.
func TestFilter(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

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

	all := []int{1, 2, 3, 4, 5}
	one := []int{1}
	none := []int{}
	sep := 0
	allJoined := []int{1, 0, 2, 0, 3, 0, 4, 0, 5}
	oneJoined := []int{1}
	noneJoined := []int{}

	assert.Equal(slices.Join(sep, all), allJoined)
	assert.Equal(slices.Join(sep, one), oneJoined)
	assert.Equal(slices.Join(sep, none), noneJoined)
}

// TestMap verifies the mapping of slice values.
func TestMap(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5}
	one := []int{1}
	none := []int{}
	mapper := func(v int) string {
	}

}

// EOF
