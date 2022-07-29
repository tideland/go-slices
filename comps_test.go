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
	"strings"
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

	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	one := []int{1}
	none := []int{}
	in := "0"
	outAll := "0987654321"
	outOne := "01"
	outNone := "0"
	stringer := func(v int, acc string) string { return fmt.Sprintf("%s%d", acc, v) }

	assert.Equal(slices.FoldR(stringer, in, all), outAll)
	assert.Equal(slices.FoldR(stringer, in, one), outOne)
	assert.Equal(slices.FoldR(stringer, in, none), outNone)
}

// TestFoldRLast verifies the right folding of a slice.
func TestFoldRLast(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	one := []int{1}
	none := []int{}
	outAll := 987654321
	outOne := 1
	outNone := 0
	potentiator := func(v, acc int) int { return acc*10 + v }

	assert.Equal(slices.FoldRLast(potentiator, all), outAll)
	assert.Equal(slices.FoldRLast(potentiator, one), outOne)
	assert.Equal(slices.FoldRLast(potentiator, none), outNone)
}

// TestMapFoldL verifies the combined mapping and folding.
func MapFoldL(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	all := []string{"one", "two", "three"}
	one := []string{"one"}
	none := []string{}
	in := 0
	outAll := []string{"ONE", "TWO", "THREE"}
	outOne := []string{"ONE"}
	outNone := []string{}
	lengthAll := 11
	lengthOne := 3
	lengthNone := 0
	upperLengther := func(v string, acc int) (string, int) {
		return strings.ToUpper(v), acc + len(v)
	}

	mapped, out := slices.MapFoldL(upperLengther, in, all)
	assert.Equal(mapped, outAll)
	assert.Equal(out, lengthAll)
	mapped, out = slices.MapFoldL(upperLengther, in, one)
	assert.Equal(mapped, outOne)
	assert.Equal(out, lengthOne)
	mapped, out = slices.MapFoldL(upperLengther, in, none)
	assert.Equal(mapped, outNone)
	assert.Equal(out, lengthNone)
}

// EOF
