package endo_test

import (
	"fmt"

	"github.com/calebcase/base/data/monoid/endo"
)

func ExampleNewType() {
	t := endo.NewType[string]()

	prefix := func(x string) string {
		return "Hello, " + x
	}

	postfix := func(x string) string {
		return x + "!"
	}

	computation := t.MAppend(prefix, postfix)

	fmt.Println(t.Apply(computation, "Haskell"))

	// Output:
	// Hello, Haskell!
}
