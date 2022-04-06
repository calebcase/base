package foldable_test

import (
	"fmt"

	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/foldable"
	"github.com/calebcase/base/data/function"
	"github.com/calebcase/base/data/list"
)

func ExampleNewType() {
	{
		t := foldable.NewType[int, int]()

		l := list.List[int]{1, 2, 3}

		sum := func(x, y int) int {
			return x + y
		}

		fmt.Println(t.FoldL(sum, 0, l))
		fmt.Println(t.FoldR(sum, 0, l))
	}

	{
		t := foldable.NewType[int, string]()

		l := list.List[int]{1, 2, 3}

		comma := func(x string, y int) string {
			if x == "" {
				return fmt.Sprint(y)
			}

			return x + ", " + fmt.Sprint(y)
		}

		fmt.Println(t.FoldL(comma, "", l))
	}

	{
		t := foldable.NewType[list.List[int], list.List[int]]()

		ll := list.List[list.List[int]]{{1, 2, 3}, {4, 5}, {6}, {}}

		lt := list.NewType[int](eq.NewType(eq.Comparable[int]))

		fmt.Println(t.Fold(lt, ll))
		fmt.Println(t.FoldMap(lt, function.Id[list.List[int]], ll))
	}

	// Output:
	// 6
	// 6
	// 1, 2, 3
	// [1 2 3 4 5 6]
	// [1 2 3 4 5 6]
}
