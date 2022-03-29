package maybe_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/calebcase/base/data/function"
	"github.com/calebcase/base/data/maybe"
	"github.com/calebcase/curry"
	"github.com/stretchr/testify/require"
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

func ExampleType_FMap() {
	flip := func(a bool) bool {
		return !a
	}

	flop := maybe.Type[bool, bool, bool]{}.FMap(flip, maybe.Just[bool]{true})
	fmt.Println(flop)

	stringify := func(a bool) string {
		return fmt.Sprint(a)
	}

	str := maybe.Type[bool, string, bool]{}.FMap(stringify, maybe.Just[bool]{true})
	fmt.Println(str)

	fmt.Println(maybe.Type[bool, string, bool]{}.FMap(stringify, maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {true}
	// {}
}

func ExampleType_FReplace() {
	fmt.Println(maybe.Type[bool, bool, bool]{}.FReplace(false, maybe.Just[bool]{true}))
	fmt.Println(maybe.Type[string, bool, bool]{}.FReplace("done", maybe.Just[bool]{true}))
	fmt.Println(maybe.Type[string, bool, bool]{}.FReplace("done", maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {done}
	// {}
}

func Test(t *testing.T) {
	mt := maybe.Type[bool, bool, bool]{}

	t.Run("functor", func(t *testing.T) {
		t.Run("fmap", func(t *testing.T) {
			t.Run("identity", func(t *testing.T) {
				// fmap id == id

				left := curry.A2R1(mt.FMap)(function.Id[bool])
				right := function.Id[maybe.Maybe[bool]]

				require.Equal(t, left(maybe.Just[bool]{true}), right(maybe.Just[bool]{true}))
			})

			t.Run("compose", func(t *testing.T) {
				// fmap (f . g) == fmap f . fmap g

				f := function.Id[bool]
				g := function.Id[bool]

				left := curry.A2R1(mt.FMap)(
					curry.A3R1(function.Compose[bool, bool, bool])(f)(g),
				)
				right := curry.A3R1(
					function.Compose[maybe.Maybe[bool], maybe.Maybe[bool], maybe.Maybe[bool]],
				)(curry.A2R1(mt.FMap)(f))(curry.A2R1(mt.FMap)(g))

				require.Equal(t, left(maybe.Just[bool]{true}), right(maybe.Just[bool]{true}))
			})
		})
	})
}
