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
	"tideland.dev/go/audit/generators"

	"tideland.dev/go/slices"
)

//--------------------
// TESTS
//--------------------

// TestSort verifies the standard sorting of slices.
func TestSort(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    []int
	}{
		{
			descr:  "Simple unordered slice",
			values: []int{5, 7, 1, 3, 4, 2, 8, 6, 9},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Unordered double value slice",
			values: []int{9, 5, 7, 3, 1, 3, 5, 4, 2, 8, 5, 6, 9},
			out:    []int{1, 2, 3, 3, 4, 5, 5, 5, 6, 7, 8, 9, 9},
		}, {
			descr:  "Already ordered slice",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Reverse ordered slice",
			values: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			out:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, {
			descr:  "Single value slice",
			values: []int{1, 1, 1, 1, 1},
			out:    []int{1, 1, 1, 1, 1},
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
		assert.Equal(slices.Sort(test.values), test.out)
	}
}

// TestSortWith verifies the sorting of slices with a less function.
func TestSortWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	less := func(vs []string, i, j int) bool { return len(vs[i]) < len(vs[j]) }
	tests := []struct {
		descr  string
		values []string
		out    []string
	}{
		{
			descr:  "Simple unordered slice",
			values: []string{"alpha", "beta", "phi", "epsilon", "lambda", "pi"},
			out:    []string{"pi", "phi", "beta", "alpha", "lambda", "epsilon"},
		}, {
			descr:  "Unordered double value slice",
			values: []string{"phi", "alpha", "beta", "phi", "epsilon", "beta", "lambda", "pi"},
			out:    []string{"pi", "phi", "phi", "beta", "beta", "alpha", "lambda", "epsilon"},
		}, {
			descr:  "Already ordered slice",
			values: []string{"pi", "phi", "beta", "alpha", "lambda", "epsilon"},
			out:    []string{"pi", "phi", "beta", "alpha", "lambda", "epsilon"},
		}, {
			descr:  "Reverse ordered slice",
			values: []string{"epsilon", "lambda", "alpha", "beta", "phi", "pi"},
			out:    []string{"pi", "phi", "beta", "alpha", "lambda", "epsilon"},
		}, {
			descr:  "Single value slice",
			values: []string{"alpha", "alpha", "alpha", "alpha", "alpha"},
			out:    []string{"alpha", "alpha", "alpha", "alpha", "alpha"},
		}, {
			descr:  "Empty slice",
			values: []string{},
			out:    []string{},
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    nil,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.SortWith(less, test.values), test.out)
	}
}

// TestIsSorted verifies the check of sorted slices.
func TestIsSorted(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	tests := []struct {
		descr  string
		values []int
		out    bool
	}{
		{
			descr:  "Unordered slice",
			values: []int{5, 7, 1, 3, 4, 2, 8, 6, 9},
			out:    false,
		}, {
			descr:  "Ordered slice",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			out:    true,
		}, {
			descr:  "Reverse ordered slice",
			values: []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			out:    false,
		}, {
			descr:  "Single value slice",
			values: []int{1, 1, 1, 1, 1},
			out:    true,
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
		assert.Equal(slices.IsSorted(test.values), test.out)
	}
}

// TestIsSortedWith verifies the check of sorted slices.
func TestIsSortedWith(t *testing.T) {
	assert := asserts.NewTesting(t, asserts.FailStop)

	less := func(a, b string) bool { return len(a) < len(b) }
	tests := []struct {
		descr  string
		values []string
		out    bool
	}{
		{
			descr:  "Unordered slice",
			values: []string{"alpha", "beta", "phi", "epsilon", "lambda", "pi"},
			out:    false,
		}, {
			descr:  "Ordered slice",
			values: []string{"pi", "phi", "beta", "alpha", "lambda", "epsilon"},
			out:    true,
		}, {
			descr:  "Reverse ordered slice",
			values: []string{"epsilon", "lambda", "alpha", "beta", "phi", "pi"},
			out:    false,
		}, {
			descr:  "Single value slice",
			values: []string{"alpha", "alpha", "alpha", "alpha", "alpha"},
			out:    true,
		}, {
			descr:  "Empty slice",
			values: []string{},
			out:    true,
		}, {
			descr:  "Nil slice",
			values: nil,
			out:    true,
		},
	}

	for _, test := range tests {
		assert.Logf(test.descr)
		assert.Equal(slices.IsSortedWith(less, test.values), test.out)
	}
}

//--------------------
// BENCHMARKS AND FUZZ TESTS
//--------------------

// BenchmarkSort runs a performance test on standard sorting.
func BenchmarkSort(b *testing.B) {
	gen := generators.New(generators.FixedRand())
	vs := gen.Ints(0, 1000, 10000)

	slices.Sort(vs)
}

// BenchmarkSortWith runs a performance test on sorting with comparator.
func BenchmarkSortWith(b *testing.B) {
	gen := generators.New(generators.FixedRand())
	vs := gen.Words(10000)
	less := func(vs []string, i, j int) bool { return len(vs[i]) < len(vs[j]) }

	slices.SortWith(less, vs)
}

// FuzzSort runs a fuzz test on the standard sorting.
func FuzzSort(f *testing.F) {
	gen := generators.New(generators.FixedRand())

	f.Add(5)
	f.Fuzz(func(t *testing.T, i int) {
		vs := gen.Ints(0, 1000, 10000)

		slices.Sort(vs)
	})
}

// EOF
