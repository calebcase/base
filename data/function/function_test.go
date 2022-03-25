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
	fmt.Println(function.Compose(
		func(x int) string {
			return fmt.Sprint(x)
		},
		func(y float32) int {
			return int(y)
		},
		3.14,
	))

	// Output:
	// 3
}

func ExampleFlip() {
	fmt.Println(function.Flip(
		func(a, b string) string {
			return a + " " + b
		},
		"world",
		"hello",
	))

	// Output:
	// hello world
}

func ExampleApply() {
	fmt.Println(function.Apply(
		func(x int) int {
			return x + 1
		},
		2,
	))

	// Output:
	// 3
}

func ExampleOn() {
	v := function.On(
		func(a, b string) string {
			return a + b
		},
		func(x int) string {
			return fmt.Sprint(x)
		},
		4,
		2,
	)
	fmt.Printf("%v %T", v, v)

	// Output:
	// 42 string
}
