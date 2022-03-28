package functor

type F[T any] interface{}

type Class[A, B any] interface {
	FMap(func(A) B, F[A]) F[B]
	FReplace(A, F[B]) F[A]
}
