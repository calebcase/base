package either_test

import (
	"fmt"

	"github.com/calebcase/base/data/either"
	"github.com/calebcase/base/data/list"
)

func ExampleEither() {
	l := either.Left[int, string]{1}
	fmt.Println(l.Value)

	r := either.Right[int, string]{"one"}
	fmt.Println(r.Value)

	{
		es := list.List[either.Either[string, int]]{
			either.Left[string, int]{"foo"},
			either.Right[string, int]{3},
			either.Left[string, int]{"bar"},
			either.Right[string, int]{7},
			either.Left[string, int]{"baz"},
		}
		fmt.Println(either.Lefts(es))
	}

	{
		t := either.NewType[string, int]()

		es := list.List[either.Either[string, int]]{
			t.NewLeft("foo"),
			t.NewRight(3),
			t.NewLeft("bar"),
			t.NewRight(7),
			t.NewLeft("baz"),
		}
		fmt.Println(either.Lefts(es))
	}

	// Output:
	// 1
	// one
	// [foo bar baz]
	// [foo bar baz]
}
