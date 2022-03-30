package num_test

import (
	"fmt"
	"testing"

	"github.com/calebcase/base/num"
)

func ExampleNewType() {
	t := num.NewType[int]()

	fmt.Println(t.Add(5, 6))

	// Output:
	// 11
}

func FuzzConformInt(f *testing.F) {
	n := num.NewType[int]()

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

	f.Fuzz(num.Conform[int](n))
}
