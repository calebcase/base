package monad

import (
	"testing"

	"github.com/calebcase/base/control/applicative"
	"github.com/stretchr/testify/require"
)

type M[T any] interface{}

type Class[
	A any,
	B any,
	C any,

	MF M[func(A) B],

	MA M[A],
	MB M[B],
	MC M[C],
] interface {
	applicative.Class[A, B, C, MF, MA, MB, MC]

	// >>=
	Bind(MA, func(A) MB) MB

	// >>
	Then(MA, MB) MB

	Return(A) MA
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[
	A any,

	MF M[func(A) A],
	MA M[A],

	CA Class[
		A,
		A,
		A,

		MF,

		MA,
		MA,
		MA,
	],
](c CA) func(t *testing.T, x A) {
	return func(t *testing.T, x A) {
		t.Run("applicative.Conform", func(t *testing.T) {
			applicative.Conform[A, MF, MA](c)(t, x)
		})

		t.Run("left identity", func(t *testing.T) {
			// return a >>= k = k a

			k := c.Return
			require.Equal(t, c.Bind(c.Return(x), k), k(x))
		})

		t.Run("right identity", func(t *testing.T) {
			// m >>= return = m

			m := c.Return(x)
			require.Equal(t, c.Bind(m, c.Return), m)
		})

		t.Run("associativity", func(t *testing.T) {
			// m >>= (\x -> k x >>= h) = (m >>= k) >>= h

			m := c.Return(x)
			k := c.Return
			h := c.Return

			left := c.Bind(m, func(x A) MA {
				return c.Bind(k(x), h)
			})
			right := c.Bind(c.Bind(m, k), h)

			require.Equal(t, left, right)
		})

		// pure = return
		require.Equal(t, c.Pure(x), c.Return(x))
	}
}
