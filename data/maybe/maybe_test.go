package maybe_test

import (
	"fmt"
	"strconv"

	"github.com/calebcase/base/data/maybe"
)

func ExampleMaybe() {
	// Odd returns true if the integer is odd.
	odd := func(v int) bool {
		return v%2 != 0
	}

	// NOTE: The additional type hinting `maybe.Apply[int](...)` is
	// currently necessary because of a limitation in Go's type
	// inferencing. The hinting may eventually be unnecessary when/if the
	// type inferencing improves for generics. Alternately the types can be
	// explicitly set on the Maybe function instead.
	fmt.Println(maybe.Apply(false, odd, maybe.Maybe[int](maybe.Just[int]{3})))
	fmt.Println(maybe.Apply[int, bool](false, odd, maybe.Just[int]{3}))

	fmt.Println(maybe.Apply(false, odd, maybe.Maybe[int](maybe.Nothing[int]{})))
	fmt.Println(maybe.Apply[int, bool](false, odd, maybe.Nothing[int]{}))

	// These all produce the desired compile time error (because the types
	// are mismatched):
	//
	//fmt.Println(maybe.Apply(false, odd, maybe.Maybe[float32](maybe.Nothing[int]{})))
	//fmt.Println(maybe.Apply(false, odd, maybe.Maybe[float32](maybe.Nothing[float32]{})))
	//fmt.Println(maybe.Apply(false, odd, maybe.Maybe[int](maybe.Just[float32]{3})))
	//fmt.Println(maybe.Apply[int, bool](false, odd, maybe.Just[float32]{3}))

	// str returns the string even or odd for the value.
	str := func(v int) string {
		if v%2 == 0 {
			return "even"
		}

		return "odd"
	}

	fmt.Println(maybe.Apply("unknown", str, maybe.Maybe[int](maybe.Just[int]{3})))
	fmt.Println(maybe.Apply("unknown", str, maybe.Maybe[int](maybe.Just[int]{4})))
	fmt.Println(maybe.Apply("unknown", str, maybe.Maybe[int](maybe.Nothing[int]{})))

	// Output:
	// true
	// true
	// false
	// false
	// odd
	// even
	// unknown
}

func ExampleCatMaybes() {
	values := []maybe.Maybe[int]{
		maybe.Just[int]{1},
		maybe.Nothing[int]{},
		maybe.Just[int]{3},
	}

	fmt.Println(maybe.CatMaybes(values))

	// Output:
	// [1 3]
}

func ExampleMapMaybes() {
	values := []string{
		"1",
		"foo",
		"3",
	}

	maybeInt := func(s string) maybe.Maybe[int] {
		i, err := strconv.Atoi(s)
		if err != nil {
			return maybe.Nothing[int]{}
		}

		return maybe.Just[int]{i}
	}

	fmt.Println(maybe.MapMaybes(maybeInt, values))

	// Output:
	// [1 3]
}
