package dual

import (
	"github.com/calebcase/base/data/monoid"
)

type Type[
	A any,
] struct {
	monoid.Type[Dual[A]]
}

func NewType[
	A any,
](m monoid.Class[A]) Type[A] {
	return Type[A]{
		Type: monoid.NewType(
			func(x Dual[A], y Dual[A]) Dual[A] {
				return Dual[A]{m.SAssoc(y.Value, x.Value)}
			},
			func() Dual[A] {
				return Dual[A]{m.MEmpty()}
			},
		),
	}
}

type Dual[A any] struct {
	Value A
}
