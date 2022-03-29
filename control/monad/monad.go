package monad

import "github.com/calebcase/base/control/applicative"

type M[T any] interface{}

type Class[A, B, C any, MF M[func(A) B], MA M[A], MB M[B], MC M[C]] interface {
	applicative.Class[A, B, C, MF, MA, MB, MC]

	Bind(MA, func(A) MB) MB
	Then(MA, MB) MB
	Return(A) MA
}
