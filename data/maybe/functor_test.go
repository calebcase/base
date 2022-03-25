package maybe_test

import (
	"fmt"

	"github.com/calebcase/base/data/maybe"
)

func ExampleFMap() {
	flip := func(a bool) bool {
		return !a
	}

	flop := maybe.FMap[bool, bool](flip, maybe.Just[bool]{true})
	fmt.Println(flop)

	stringify := func(a bool) string {
		return fmt.Sprint(a)
	}

	str := maybe.FMap[bool, string](stringify, maybe.Just[bool]{true})
	fmt.Println(str)

	fmt.Println(maybe.FMap[bool, string](stringify, maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {true}
	// {}
}

func ExampleFReplace() {
	fmt.Println(maybe.FReplace[bool, bool](false, maybe.Just[bool]{true}))
	fmt.Println(maybe.FReplace[string, bool]("done", maybe.Just[bool]{true}))
	fmt.Println(maybe.FReplace[string, bool]("done", maybe.Nothing[bool]{}))

	// Output:
	// {false}
	// {done}
	// {}
}
