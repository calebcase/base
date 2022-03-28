package list

import (
	"github.com/calebcase/base/data/monoid"
)

type Class[A any] interface {
	monoid.Class[A]
}

type Type[A any, L []A] struct{}

var _ Class[[]int] = Type[int, []int]{}

func (t Type[A, L]) SAssoc(x, y L) L {
	r := make(L, 0, len(x)+len(y))
	r = append(t.MEmpty(), x...)
	r = append(r, y...)

	return r
}

func (t Type[A, L]) MEmpty() L {
	return L{}
}

func (t Type[A, L]) MAppend(x, y L) L {
	return t.SAssoc(x, y)
}

func (t Type[A, L]) MConcat(xs []L) L {
	length := 0

	for _, x := range xs {
		length += len(x)
	}

	r := make(L, 0, length)

	for _, x := range xs {
		r = append(r, x...)
	}

	return r
}
