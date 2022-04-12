package foldable_test

import (
	"fmt"

	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/foldable"
	"github.com/calebcase/base/data/function"
	"github.com/calebcase/base/data/list"
	"github.com/calebcase/base/data/maybe"
	"github.com/calebcase/base/data/ord"
)

func ExampleType_Fold() {
	t := foldable.NewType[list.List[int], list.List[int]]()

	ll := list.List[list.List[int]]{
		{1, 2, 3},
		{4, 5},
		{6},
		{},
	}

	lt := list.NewType[int]()

	fmt.Println(t.Fold(lt, ll))

	// Output:
	// [1 2 3 4 5 6]
}

func ExampleType_FoldMap() {
	t := foldable.NewType[list.List[int], list.List[int]]()

	ll := list.List[list.List[int]]{
		{1, 2, 3},
		{4, 5},
		{6},
		{},
	}

	lt := list.NewType[int]()

	fmt.Println(t.FoldMap(lt, function.Id[list.List[int]], ll))

	// Output:
	// [1 2 3 4 5 6]
}

func ExampleType_FoldL() {
	t := foldable.NewType[int, string]()

	l := list.List[int]{1, 2, 3}

	comma := func(acc string, x int) string {
		if acc == "" {
			return fmt.Sprint(x)
		}

		return acc + ", " + fmt.Sprint(x)
	}

	fmt.Println(t.FoldL(comma, "", l))

	// Output:
	// 1, 2, 3
}

func ExampleType_FoldR() {
	t := foldable.NewType[int, int]()

	l := list.List[int]{1, 2, 3}

	sum := func(acc, x int) int {
		return acc + x
	}

	fmt.Println(t.FoldR(sum, 0, l))

	// Output:
	// 6
}

func ExampleType_ToList() {
	t := foldable.NewType[int, int]()

	fmt.Println(t.ToList(maybe.Just[int]{Value: 2}))

	// Output:
	// [2]
}

func ExampleType_Null_maybe() {
	t := foldable.NewType[int, int]()

	fmt.Println(t.Null(maybe.Just[int]{Value: 2}))
	fmt.Println(t.Null(maybe.Nothing[int]{}))

	// Output:
	// false
	// true
}

func ExampleType_Null_list() {
	t := foldable.NewType[int, int]()

	fmt.Println(t.Null(list.List[int]{1, 2, 3}))
	fmt.Println(t.Null(list.List[int]{}))

	// Output:
	// false
	// true
}

func ExampleType_Length_maybe() {
	t := foldable.NewType[int, int]()

	fmt.Println(t.Length(maybe.Just[int]{Value: 2}))
	fmt.Println(t.Length(maybe.Nothing[int]{}))

	// Output:
	// 1
	// 0
}

func ExampleType_Length_list() {
	t := foldable.NewType[int, int]()

	fmt.Println(t.Length(list.List[int]{1, 2, 3}))
	fmt.Println(t.Length(list.List[int]{}))

	// Output:
	// 3
	// 0
}

func ExampleType_Elem_int() {
	t := foldable.NewType[int, int]()

	l := list.List[int]{1, 2, 3}
	e := eq.NewType(eq.Comparable[int])

	fmt.Println(t.Elem(e, 2, l))
	fmt.Println(t.Elem(e, 10, l))

	// Output:
	// true
	// false
}

func ExampleType_Elem_string() {
	t := foldable.NewType[string, string]()

	l := list.List[string]{
		"a",
		"needle",
		"in",
		"a",
		"haystack",
	}
	e := eq.NewType(eq.Comparable[string])

	fmt.Println(t.Elem(e, "needle", l))
	fmt.Println(t.Elem(e, "gold", l))

	// Output:
	// true
	// false
}

func ExampleType_Maximum_int() {
	t := foldable.NewType[int, int]()

	l := list.List[int]{1, 2, 3}
	o := ord.NewType(ord.Ordered[int])

	fmt.Println(t.Maximum(o, l))

	// Output:
	// 3
}

func ExampleType_Maximum_string() {
	t := foldable.NewType[string, string]()

	l := list.List[string]{"a", "b", "c"}
	o := ord.NewType(ord.Ordered[string])

	fmt.Println(t.Maximum(o, l))

	// Output:
	// c
}

func ExampleType_Minimum_int() {
	t := foldable.NewType[int, int]()

	l := list.List[int]{1, 2, 3}
	o := ord.NewType(ord.Ordered[int])

	fmt.Println(t.Minimum(o, l))

	// Output:
	// 1
}

func ExampleType_Minimum_string() {
	t := foldable.NewType[string, string]()

	l := list.List[string]{"a", "b", "c"}
	o := ord.NewType(ord.Ordered[string])

	fmt.Println(t.Minimum(o, l))

	// Output:
	// a
}
