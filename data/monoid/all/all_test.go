package all_test

import (
	"fmt"

	"github.com/calebcase/base/data/list"
	"github.com/calebcase/base/data/monoid/all"
)

func ExampleNewType() {
	t := all.NewType()

	fmt.Println(
		t.SAssoc(
			true,
			t.SAssoc(
				t.MEmpty(),
				false,
			),
		),
	)

	fmt.Println(
		t.MConcat(
			list.Map(
				func(x int) bool {
					return x%2 == 0
				},
				list.List[int]{2, 4, 6, 7, 8},
			),
		),
	)

	// Output:
	// false
	// false
}
