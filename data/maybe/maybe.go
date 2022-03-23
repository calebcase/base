package maybe

import "errors"

// Type is the sum type for maybe.
type Type[T any] interface {
	isType(T)
}

// Just contains a value.
type Just[T any] struct {
	Value T
}

func (j Just[T]) isType(_ T) {}

// Nothing indicates no value is present.
type Nothing[T any] struct{}

func (n Nothing[T]) isType(_ T) {}

// Maybe returns the default value `dflt` if `v` is Nothing. Otherwise it
// returns the result of calling `f` on `v`.
func Maybe[A, B any](dflt B, f func(a A) B, v Type[A]) B {
	switch v.(type) {
	case Just[A]:
		return f(v.(Just[A]).Value)
	case Nothing[A]:
		return dflt
	}

	panic(errors.New("impossible"))
}

func IsJust[A any](v Type[A]) bool {
	switch v.(type) {
	case Just[A]:
		return true
	case Nothing[A]:
		return false
	}

	panic(errors.New("impossible"))
}

func IsNothing[A any](v Type[A]) bool {
	switch v.(type) {
	case Just[A]:
		return false
	case Nothing[A]:
		return true
	}

	panic(errors.New("impossible"))
}

func FromJust[A any](v Type[A]) A {
	return v.(Just[A]).Value
}

func FromMaybe[A any](dflt A, v Type[A]) A {
	switch v.(type) {
	case Just[A]:
		return v.(Just[A]).Value
	case Nothing[A]:
		return dflt
	}

	panic(errors.New("impossible"))
}

func ListToMaybe[A any](vs []A) Type[A] {
	if len(vs) == 0 {
		return Nothing[A]{}
	}

	return Just[A]{vs[0]}
}

func MaybeToList[A any](v Type[A]) []A {
	switch v.(type) {
	case Just[A]:
		return []A{v.(Just[A]).Value}
	case Nothing[A]:
		return []A{}
	}

	panic(errors.New("impossible"))
}

func CatMaybes[A any](vs []Type[A]) []A {
	rs := make([]A, 0, len(vs))

	for _, v := range vs {
		if j, ok := v.(Just[A]); ok {
			rs = append(rs, j.Value)
		}
	}

	return rs
}

func MapMaybes[A, B any](f func(A) Type[B], vs []A) (rs []B) {
	for _, v := range vs {
		r := f(v)

		if j, ok := r.(Just[B]); ok {
			rs = append(rs, j.Value)
		}
	}

	return rs
}
