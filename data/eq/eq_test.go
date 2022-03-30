package eq_test

import (
	"fmt"
	"testing"

	"github.com/calebcase/base/data/eq"
)

func ExampleNewType() {
	t := eq.NewType[int](eq.Comparable[int])

	fmt.Println(t.Equal(1, 1))
	fmt.Println(t.NE(1, 2))

	// Output:
	// true
	// true
}

func FuzzConformInt(f *testing.F) {
	e := eq.NewType[int](eq.Comparable[int])

	type TC struct {
		x int
		y int
		z int
	}

	tcs := []TC{
		{1, 2, 3},
		{0, 0, 0},
		{3, 2, 1},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(eq.Conform[int](e))
}

func FuzzConformString(f *testing.F) {
	e := eq.NewType[string](eq.Comparable[string])

	type TC struct {
		x string
		y string
		z string
	}

	tcs := []TC{
		{"foo", "bar", "baz"},
		{"foo", "foo", "foo"},
		{"a", "b", "c"},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(eq.Conform[string](e))
}
