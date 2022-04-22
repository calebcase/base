package foldable

import (
	"github.com/calebcase/base/data"
	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/monoid"
	"github.com/calebcase/base/data/ord"
)

type T[T any] interface{}

type Class[
	A any,
	B any,

	MA monoid.Class[A],
	MB monoid.Class[B],

	EA eq.Class[A],
	OA ord.Class[A],

	TA T[A],
] interface {
	Fold(MA, TA) A
	FoldMap(MB, func(A) B, TA) B

	FoldR(func(A, B) B, B, TA) B
	FoldL(func(B, A) B, B, TA) B

	ToList(TA) []A
	Null(TA) bool
	Length(TA) int

	Elem(EA, A, TA) bool

	Maximum(OA, TA) A
	Minimum(OA, TA) A

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
	eq.Class[int],
	ord.Class[int],
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

	if ta.DEmpty() {
		return zero
	}

	value := ta.DValue()
	result := ma.MAppend(zero, value)

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = ma.MAppend(result, rest.DValue())
	}

	return result
}

func (t Type[A, B]) FoldMap(mb monoid.Class[B], f func(A) B, ta data.Data[A]) B {
	zero := mb.MEmpty()

	if ta.DEmpty() {
		return zero
	}

	value := ta.DValue()
	result := mb.MAppend(zero, f(value))

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = mb.MAppend(result, f(rest.DValue()))
	}

	return result
}

func (t Type[A, B]) FoldR(f func(A, B) B, z B, ta data.Data[A]) B {
	if ta.DEmpty() {
		return z
	}

	if ta.DRest() == nil {
		return f(ta.DValue(), z)
	}

	return f(ta.DValue(), t.FoldR(f, z, ta.DRest()))
}

func (t Type[A, B]) FoldL(f func(B, A) B, z B, ta data.Data[A]) B {
	if ta.DEmpty() {
		return z
	}

	value := ta.DValue()
	result := f(z, value)

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		result = f(result, rest.DValue())
	}

	return result
}

func (t Type[A, B]) ToList(ta data.Data[A]) (result []A) {
	return NewType[A, []A]().FoldL(func(l []A, x A) []A {
		return append(l, x)
	}, []A{}, ta)
}

func (t Type[A, B]) Null(ta data.Data[A]) bool {
	return ta.DEmpty()
}

func (t Type[A, B]) Length(ta data.Data[A]) int {
	return NewType[A, int]().FoldL(func(total int, x A) int {
		return total + 1
	}, 0, ta)
}

func (t Type[A, B]) Elem(ea eq.Class[A], a A, ta data.Data[A]) bool {
	if ta.DEmpty() {
		return false
	}

	if ea.Equal(a, ta.DValue()) {
		return true
	}

	for rest := ta.DRest(); rest != nil; rest = rest.DRest() {
		if ea.Equal(a, rest.DValue()) {
			return true
		}
	}

	return false
}

func (t Type[A, B]) Maximum(oa ord.Class[A], ta data.Data[A]) A {
	return NewType[A, A]().FoldL(func(max A, x A) A {
		if oa.GT(x, max) {
			return x
		}

		return max
	}, ta.DValue(), ta.DRest())
}

func (t Type[A, B]) Minimum(oa ord.Class[A], ta data.Data[A]) A {
	return NewType[A, A]().FoldL(func(min A, x A) A {
		if oa.LT(x, min) {
			return x
		}

		return min
	}, ta.DValue(), ta.DRest())
}

// Special biased folds

//TODO: foldrM
//TODO: foldlM

// Folding actions

//TODO: traverse_
//TODO: for_
//TODO: sequenceA_
//TODO: asum

// Monadic actions

//TODO: mapM_
//TODO: forM_
//TODO: sequence_
//TODO: msum

// Specialized folds

func Concat[
	A any,
	FA Class[
		A,
		A,
		monoid.Class[A],
		monoid.Class[A],
		eq.Class[A],
		ord.Class[A],
		data.Data[A],
	],
	MA monoid.Class[A],
	DA data.Data[A],
](fa FA, ma MA, ta DA) A {
	return fa.Fold(ma, ta)
}
