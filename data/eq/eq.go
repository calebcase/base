package eq

import (
	"testing"

	"github.com/calebcase/base/data/function"
	"github.com/stretchr/testify/require"
)

type Class[A any] interface {
	Equal(A, A) bool
	NE(A, A) bool
}

type Type[A any] struct {
	equalFn EqualFn[A]
}

// Ensure Type implements Class.
var _ Class[int] = Type[int]{}

type EqualFn[A any] func(A, A) bool

func NewType[A any](equalFn EqualFn[A]) Type[A] {
	return Type[A]{
		equalFn: equalFn,
	}
}

func (t Type[A]) Equal(x, y A) bool {
	return t.equalFn(x, y)
}

func (t Type[A]) NE(x, y A) bool {
	return !t.Equal(x, y)
}

// Comparable implements EqualFn for natively comparable types.
func Comparable[A comparable](x, y A) bool {
	return x == y
}

// Ensure Comparable can be used with NewType.
var _ Type[int] = NewType[int](Comparable[int])

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("reflexivity", func(t *testing.T) {
			require.True(t, c.Equal(x, x))
		})

		t.Run("symmetry", func(t *testing.T) {
			require.True(t, c.Equal(x, y) == c.Equal(y, x))
		})

		t.Run("transitivity", func(t *testing.T) {
			if c.Equal(x, y) && c.Equal(y, z) {
				require.True(t, c.Equal(x, z))
			}
		})

		t.Run("extensionality", func(t *testing.T) {
			if c.Equal(x, y) {
				require.True(t, c.Equal(function.Id(x), function.Id(y)))
			}
		})

		t.Run("negation", func(t *testing.T) {
			require.True(t, c.NE(x, y) == !c.Equal(x, y))
		})
	}
}
