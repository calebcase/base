package list

import (
	"testing"

	"github.com/calebcase/base/data"
	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/monoid"
)

type Class[A any] interface {
	monoid.Class[List[A]]
}

type Type[
	A any,
] struct {
	monoid.Type[List[A]]
}

// Ensure Type implements Class.
var _ Class[int] = Type[int]{}

func NewType[
	A any,
]() Type[A] {
	return Type[A]{
		Type: monoid.NewType[List[A]](
			func(x, y List[A]) List[A] {
				r := make(List[A], 0, len(x)+len(y))
				r = append(List[A]{}, x...)
				r = append(r, y...)

				return r
			},
			func() List[A] {
				return List[A]{}
			},
		),
	}
}

type List[A any] []A

// Ensure List implements data.Data.
var _ data.Data[int] = List[int]{}

func (l List[A]) DValue() A {
	return l[0]
}

func (l List[A]) DRest() data.Data[A] {
	if len(l) > 1 {
		return l[1:]
	}

	return nil
}

// NewEqualFn returns a list equality checking function given the eq.Class for
// the type A.
func NewEqualFn[A any](e eq.Class[A]) func(x, y List[A]) bool {
	return func(x, y List[A]) bool {
		if len(x) != len(y) {
			return false
		}

		for i := 0; i < len(x); i++ {
			if e.NE(x[i], y[i]) {
				return false
			}
		}

		return true
	}
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z List[A]) {
	return func(t *testing.T, x, y, z List[A]) {
		t.Run("monoid.Conform", func(t *testing.T) {
			monoid.Conform[List[A]](c)(t, x, y, z)
		})
	}
}
