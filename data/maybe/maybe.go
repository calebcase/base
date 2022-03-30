package maybe

import (
	"github.com/calebcase/base/control/monad"
)

type Class[A, B, C any] interface {
	monad.Class[A, B, C, Maybe[func(A) B], Maybe[A], Maybe[B], Maybe[C]]
}

type Type[A, B, C any] struct{}

func NewType[A, B, C any]() Type[A, B, C] {
	return Type[A, B, C]{}
}

var _ Class[int, int, int] = Type[int, int, int]{}

func (t Type[A, B, C]) FMap(f func(A) B, v Maybe[A]) Maybe[B] {
	if j, ok := v.(Just[A]); ok {
		return Just[B]{f(j.Value)}
	}

	return Nothing[B]{}
}

func (t Type[A, B, C]) FReplace(a A, v Maybe[B]) Maybe[A] {
	if _, ok := v.(Just[B]); ok {
		return Just[A]{a}
	}

	return Nothing[A]{}
}

func (t Type[A, B, C]) Pure(x A) Maybe[A] {
	return Just[A]{x}
}

func (t Type[A, B, C]) Apply(f Maybe[func(A) B], m Maybe[A]) Maybe[B] {
	if jf, ok := f.(Just[func(A) B]); ok {
		return t.FMap(jf.Value, m)
	}

	return Nothing[B]{}
}

func (t Type[A, B, C]) LiftA2(f func(A, B) C, x Maybe[A], y Maybe[B]) Maybe[C] {
	jx, ok := x.(Just[A])
	if !ok {
		return Nothing[C]{}
	}

	jy, ok := y.(Just[B])
	if !ok {
		return Nothing[C]{}
	}

	return Just[C]{f(jx.Value, jy.Value)}
}

func (t Type[A, B, C]) ApplyR(x Maybe[A], y Maybe[B]) Maybe[B] {
	return y
}

func (t Type[A, B, C]) ApplyL(x Maybe[A], y Maybe[B]) Maybe[A] {
	return x
}

func (t Type[A, B, C]) Bind(x Maybe[A], k func(A) Maybe[B]) Maybe[B] {
	if jx, ok := x.(Just[A]); ok {
		return k(jx.Value)
	}

	return Nothing[B]{}
}

func (t Type[A, B, C]) Then(x Maybe[A], y Maybe[B]) Maybe[B] {
	return t.ApplyR(x, y)
}

func (t Type[A, B, C]) Return(x A) Maybe[A] {
	return t.Pure(x)
}

// Maybe is the sum type for maybe.
type Maybe[T any] interface {
	isMaybe(T)
}

// Just contains a value.
type Just[T any] struct {
	Value T
}

func (j Just[T]) isMaybe(_ T) {}

// Nothing indicates no value is present.
type Nothing[T any] struct{}

func (n Nothing[T]) isMaybe(_ T) {}

// Apply returns the default value `dflt` if `v` is Nothing. Otherwise it
// returns the result of calling `f` on `v`.
func Apply[A, B any](dflt B, f func(a A) B, v Maybe[A]) B {
	if j, ok := v.(Just[A]); ok {
		return f(j.Value)
	}

	return dflt
}

func IsJust[A any](v Maybe[A]) bool {
	_, ok := v.(Just[A])

	return ok
}

func IsNothing[A any](v Maybe[A]) bool {
	_, ok := v.(Nothing[A])

	return ok
}

func FromJust[A any](v Maybe[A]) A {
	return v.(Just[A]).Value
}

func FromMaybe[A any](dflt A, v Maybe[A]) A {
	if j, ok := v.(Just[A]); ok {
		return j.Value
	}

	return dflt
}

func ListToMaybe[A any](vs []A) Maybe[A] {
	if len(vs) == 0 {
		return Nothing[A]{}
	}

	return Just[A]{vs[0]}
}

func MaybeToList[A any](v Maybe[A]) []A {
	if j, ok := v.(Just[A]); ok {
		return []A{j.Value}
	}

	return []A{}
}

func CatMaybes[A any](vs []Maybe[A]) []A {
	rs := make([]A, 0, len(vs))

	for _, v := range vs {
		if j, ok := v.(Just[A]); ok {
			rs = append(rs, j.Value)
		}
	}

	return rs
}

func MapMaybes[A, B any](f func(A) Maybe[B], vs []A) (rs []B) {
	for _, v := range vs {
		r := f(v)

		if j, ok := r.(Just[B]); ok {
			rs = append(rs, j.Value)
		}
	}

	return rs
}
