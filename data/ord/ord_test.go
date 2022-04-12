package ord_test

import (
	"fmt"
	"testing"

	"github.com/calebcase/base/data/ord"
)

func ExampleNewType() {
	t := ord.NewType(ord.Ordered[int])

	fmt.Println(t.LT(5, 6))

	// Output:
	// true
}

func FuzzConformInt(f *testing.F) {
	o := ord.NewType(ord.Ordered[int])

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

	f.Fuzz(ord.Conform[int](o))
}

func FuzzConformString(f *testing.F) {
	o := ord.NewType(ord.Ordered[string])

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

	f.Fuzz(ord.Conform[string](o))
}
