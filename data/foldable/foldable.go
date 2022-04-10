package foldable

import (
	"github.com/calebcase/base/data"
	"github.com/calebcase/base/data/monoid"
)

type Class[
	A any,
	B any,

	MA monoid.Class[A],
	MB monoid.Class[B],

	DA data.Data[A],
] interface {
	//Fold(TM) M
	//FoldMap(func(A) M, TA) M
	Fold(MA, DA) A
	FoldMap(MB, func(A) B, DA) B

	//FoldR(func(A, B) B, B, TA) B
	//FoldL(func(B, A) B, B, TA) B
	FoldR(func(A, B) B, B, DA) B
	FoldL(func(B, A) B, B, DA) B

	//ToList(TA) []A
	//Null(TA) bool
	//Length(TA) int

	//Elem(E, TE) bool

	//Maximum(TO) A
	//Minimum(TO) A

	//Sum(TN) A
	//Product(TN) A
}

type Type[
	A any,
	B any,
] struct{}

// Ensure Type implements Class.
var _ Class[
	int,
	int,
	monoid.Class[int],
	monoid.Class[int],
	data.Data[int],
] = Type[int, int]{}

func NewType[
	A any,
	B any,
]() Type[A, B] {
	return Type[A, B]{}
}

func (t Type[A, B]) Fold(ma monoid.Class[A], ta data.Data[A]) A {
	zero := ma.MEmpty()
	value := ta.DValue()
	result := ma.MAppend(zero, value)

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = ma.MAppend(result, rest.DValue())
	}

	return result
}

func (t Type[A, B]) FoldMap(mb monoid.Class[B], f func(A) B, ta data.Data[A]) B {
	zero := mb.MEmpty()
	value := ta.DValue()
	result := mb.MAppend(zero, f(value))

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = mb.MAppend(result, f(rest.DValue()))
	}

	return result
}

func (t Type[A, B]) FoldR(f func(A, B) B, z B, ta data.Data[A]) B {
	if ta.DRest() == nil {
		return f(ta.DValue(), z)
	}

	return f(ta.DValue(), t.FoldR(f, z, ta.DRest()))
}

func (t Type[A, B]) FoldL(f func(B, A) B, z B, ta data.Data[A]) B {
	value := ta.DValue()
	result := f(z, value)

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = f(result, rest.DValue())
	}

	return result
}
