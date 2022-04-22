package list

import (
	"testing"

	"github.com/calebcase/base/data"
	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/monoid"
)

type Class[A any] interface {
	monoid.Class[List[A]]
}

type Type[
	A any,
] struct {
	monoid.Type[List[A]]
}

// Ensure Type implements Class.
var _ Class[int] = Type[int]{}

func NewType[
	A any,
]() Type[A] {
	return Type[A]{
		Type: monoid.NewType(
			func(x, y List[A]) List[A] {
				r := make(List[A], 0, len(x)+len(y))
				r = append(List[A]{}, x...)
				r = append(r, y...)

				return r
			},
			func() List[A] {
				return List[A]{}
			},
		),
	}
}

type List[A any] []A

// Ensure List implements data.Data.
var _ data.Data[int] = List[int]{}

func (l List[A]) DEmpty() bool {
	return len(l) == 0
}

func (l List[A]) DValue() A {
	return l[0]
}

func (l List[A]) DRest() data.Data[A] {
	if len(l) > 1 {
		return l[1:]
	}

	return nil
}

// NewEqualFn returns a list equality checking function given the eq.Class for
// the type A.
func NewEqualFn[A any, LA ~[]A](e eq.Class[A]) func(x, y LA) bool {
	return func(x, y LA) bool {
		if len(x) != len(y) {
			return false
		}

		for i := 0; i < len(x); i++ {
			if e.NE(x[i], y[i]) {
				return false
			}
		}

		return true
	}
}

// List transformations

func Map[A, B any, LA ~[]A](fn func(A) B, xs LA) []B {
	ys := make([]B, 0, len(xs))

	for _, x := range xs {
		ys = append(ys, fn(x))
	}

	return ys
}

func Reverse[A any, LA ~[]A](xs LA) []A {
	ys := make([]A, len(xs))

	for i, x := range xs {
		ys[len(xs)-1-i] = x
	}

	return ys
}

func Intersperse[A any, LA ~[]A](v A, xs LA) []A {
	ys := make([]A, 0, len(xs)+len(xs)/2)

	for i, x := range xs {
		if len(xs)-1 == i {
			ys = append(ys, x)
		} else {
			ys = append(ys, x, v)
		}
	}

	return ys
}

func Intercalate[A any, LA ~[]A, LLA ~[]LA](xs LA, xss LLA) []A {
	return Concat(Intersperse(xs, xss))
}

func Transpose[A any, LA ~[]A, LLA ~[]LA](xss LLA) [][]A {
	result := [][]A{}

	for _, row := range xss {
		for j, col := range row {
			if len(result)-1 < j {
				result = append(result, []A{})
			}

			result[j] = append(result[j], col)
		}
	}

	return result
}

func FoldR[A, B any, LA ~[]A](f func(A, B) B, z B, xs LA) B {
	if len(xs) == 0 {
		return z
	}

	if len(xs) == 1 {
		return f(xs[0], z)
	}

	return f(xs[0], FoldR(f, z, xs[1:]))
}

func NonEmptySubsequences[A any, LA ~[]A](la LA) [][]A {
	if len(la) == 0 {
		return [][]A{}
	}

	x := la[0]
	xs := la[1:]

	f := func(ys []A, r [][]A) [][]A {
		m := append([]A{x}, ys...)

		return append([][]A{ys, m}, r...)
	}

	return append([][]A{{x}}, FoldR(f, [][]A{}, NonEmptySubsequences(xs))...)
}

func Subsequences[A any, LA ~[]A](xs LA) [][]A {
	return append([][]A{{}}, NonEmptySubsequences(xs)...)
}

type T[A any] interface{}

func Concat[A any, LA ~[]A, LLA ~[]LA](xss LLA) []A {
	result := LA{}

	if len(xss) == 0 {
		return result
	}

	for _, xs := range xss {
		result = append(result, xs...)
	}

	return result
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z List[A]) {
	return func(t *testing.T, x, y, z List[A]) {
		t.Run("monoid.Conform", func(t *testing.T) {
			monoid.Conform[List[A]](c)(t, x, y, z)
		})
	}
}
