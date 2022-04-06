package function_test

import (
	"fmt"

	"github.com/calebcase/base/data/function"
)

func ExampleId() {
	fmt.Println(function.Id(3))
	fmt.Println(function.Id("hello"))

	// Output:
	// 3
	// hello
}

func ExampleConst() {
	fmt.Println(function.Const(1, true))

	// Output:
	// 1
}

func ExampleCompose() {
	i2s := func(x int) string {
		return fmt.Sprint(x)
	}
	f2i := func(y float32) int {
		return int(y)
	}

	v := function.Compose(i2s, f2i)(3.14)
	fmt.Printf("%v %T", v, v)

	// Output:
	// 3 string
}

func ExampleFlip() {
	cat := func(a, b string) string {
		return a + "\n" + b
	}

	v := function.Flip(cat)("world", "hello")
	fmt.Printf("%v\n%T", v, v)

	// Output:
	// hello
	// world
	// string
}

func ExampleApply() {
	increment := func(x int) int {
		return x + 1
	}

	v := function.Apply(increment, 2)
	fmt.Printf("%v %T", v, v)

	// Output:
	// 3 int
}

func ExampleOn() {
	concat := func(a, b string) string {
		return a + b
	}

	i2s := func(x int) string {
		return fmt.Sprint(x)
	}

	v := function.On(concat, i2s)(4, 2)
	fmt.Printf("%v %T", v, v)

	// Output:
	// 42 string
}
