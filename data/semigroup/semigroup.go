package semigroup

type Class[A any] interface {
	SAssoc(A, A) A
}
