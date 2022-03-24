package monad

type Functor[A, B any] interface {
	Fmap(func(A) B, Functor[A, B]) Functor[A, B]
	Freplace(A, Functor[A, B]) Functor[A, B]
}
