package maybe_test

import (
	"fmt"

	"github.com/calebcase/base/data/maybe"
)

func ExampleFmap() {
	flip := func(a bool) bool {
		return !a
	}

	flop := maybe.Fmap[bool, bool](flip, maybe.Just[bool]{true})
	fmt.Println(flop)

	stringify := func(a bool) string {
		return fmt.Sprint(a)
	}

	str := maybe.Fmap[bool, string](stringify, maybe.Just[bool]{true})
	fmt.Println(str)

	fmt.Println(maybe.Fmap[bool, string](stringify, maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {true}
	// {}
}

func ExampleFreplace() {
	fmt.Println(maybe.Freplace[bool, bool](false, maybe.Just[bool]{true}))
	fmt.Println(maybe.Freplace[string, bool]("done", maybe.Just[bool]{true}))
	fmt.Println(maybe.Freplace[string, bool]("done", maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {done}
	// {}
}
