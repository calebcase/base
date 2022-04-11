package applicative

import (
	"testing"

	"github.com/calebcase/base/data/functor"
)

type F[T any] interface{}

type Class[
	A any,
	B any,
	C any,

	FF F[func(A) B],

	FA F[A],
	FB F[B],
	FC F[C],
] interface {
	functor.Class[A, B, FA, FB]

	Pure(A) FA

	// <*>
	Apply(FF, FA) FB

	LiftA2(func(A, B) C, FA, FB) FC

	// *>
	ApplyR(FA, FB) FB

	// <*
	ApplyL(FA, FB) FA
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[
	A any,

	FF F[func(A) A],
	FA F[A],

	CA Class[
		A,
		A,
		A,

		FF,

		FA,
		FA,
		FA,
	],
](c CA) func(t *testing.T, x A) {
	return func(t *testing.T, x A) {
		t.Run("functor.Conform", func(t *testing.T) {
			functor.Conform[A, FA](c)(t, c.Pure(x))
		})

		// FIXME: Add the applicative specific laws.
	}
}
