package endo

import (
	"github.com/calebcase/base/data/function"
	"github.com/calebcase/base/data/monoid"
)

type Type[
	A any,
	F Endo[A],
] struct {
	monoid.Type[F]
}

func NewType[
	A any,
	F Endo[A],
]() Type[A, F] {
	return Type[A, F]{
		Type: monoid.NewType(
			func(x F, y F) F {
				return function.Compose(x, y)
			},
			func() F {
				return function.Id[A]
			},
		),
	}
}

func (t Type[A, F]) Apply(f F, x A) A {
	return f(x)
}

type Endo[A any] func(A) A
