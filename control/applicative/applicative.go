package applicative

import "github.com/calebcase/base/data/functor"

type F[T any] interface{}

type Class[A, B, C any, FF F[func(A) B], FA F[A], FB F[B], FC F[C]] interface {
	functor.Class[A, B, FA, FB]

	Pure(A) FA
	Apply(FF, FA) FB
	LiftA2(func(A, B) C, FA, FB) FC
	ApplyR(FA, FB) FB
	ApplyL(FA, FB) FA
}
