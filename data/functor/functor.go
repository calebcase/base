package functor

import (
	"testing"

	"github.com/calebcase/base/data/function"
	"github.com/calebcase/curry"
	"github.com/stretchr/testify/require"
)

type F[T any] interface{}

type Class[
	A any,
	B any,

	FA F[A],
	FB F[B],
] interface {
	FMap(func(A) B, FA) FB
	FReplace(A, FB) FA
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[
	A any,

	FA F[A],

	CA Class[A, A, FA, FA],
](c CA) func(t *testing.T, x FA) {
	return func(t *testing.T, x FA) {
		t.Run("identity", func(t *testing.T) {
			// fmap id == id

			left := curry.A2R1(c.FMap)(function.Id[A])
			right := function.Id[FA]

			require.Equal(t, left(x), right(x))
		})

		t.Run("compose", func(t *testing.T) {
			// fmap (f . g) == fmap f . fmap g

			f := function.Id[A]
			g := function.Id[A]

			left := curry.A2R1(c.FMap)(function.Compose(f, g))
			right := function.Compose(
				curry.A2R1(c.FMap)(f),
				curry.A2R1(c.FMap)(g),
			)

			require.Equal(t, left(x), right(x))
		})
	}
}
