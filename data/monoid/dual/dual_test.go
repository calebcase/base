package dual_test

import (
	"fmt"

	"github.com/calebcase/base/data/monoid"
	"github.com/calebcase/base/data/monoid/dual"
)

func ExampleNewType() {
	t := dual.NewType[string](monoid.NewType(
		func(x string, y string) string {
			return x + y
		},
		func() string {
			return ""
		},
	))

	d := t.MAppend(
		dual.Dual[string]{Value: "Hello"},
		dual.Dual[string]{Value: "World"},
	)

	fmt.Println(d.Value)

	// Output:
	// WorldHello
}
