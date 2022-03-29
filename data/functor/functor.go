package functor

type F[T any] interface{}

type Class[A, B any, FA F[A], FB F[B]] interface {
	FMap(func(A) B, FA) FB
	FReplace(A, FB) FA
}
