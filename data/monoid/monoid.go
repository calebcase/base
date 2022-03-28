package monoid

import "github.com/calebcase/base/data/semigroup"

type Class[A any] interface {
	semigroup.Class[A]

	MEmpty() A
	MAppend(x, y A) A
	MConcat([]A) A
}
