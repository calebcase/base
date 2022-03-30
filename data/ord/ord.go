package ord

import (
	"testing"

	"github.com/calebcase/base/data/eq"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

type Class[A any] interface {
	eq.Class[A]

	Compare(A, A) Ordering

	LT(A, A) bool
	LTE(A, A) bool
	GT(A, A) bool
	GTE(A, A) bool

	Max(A, A) A
	Min(A, A) A
}

type Type[A any] struct {
	eq.Type[A]

	compareFn CompareFn[A]
}

// Ensure Type implements Class.
var _ Class[int] = Type[int]{}

type CompareFn[A any] func(A, A) Ordering

// NewType derives a type implementing Class.
func NewType[A any](compareFn CompareFn[A]) Type[A] {
	return Type[A]{
		Type: eq.NewType[A](func(x, y A) bool {
			return compareFn(x, y) == EQ{}
		}),

		compareFn: compareFn,
	}
}

func (t Type[A]) Compare(x, y A) Ordering {
	return t.compareFn(x, y)
}

func (t Type[A]) LT(x, y A) bool {
	return t.Compare(x, y) == LT{}
}

func (t Type[A]) LTE(x, y A) bool {
	o := t.Compare(x, y)

	return o == LT{} || o == EQ{}
}

func (t Type[A]) GT(x, y A) bool {
	return t.Compare(x, y) == GT{}
}

func (t Type[A]) GTE(x, y A) bool {
	o := t.Compare(x, y)

	return o == GT{} || o == EQ{}
}

func (t Type[A]) Max(x, y A) A {
	o := t.Compare(x, y)

	if (o == GT{} || o == EQ{}) {
		return x
	}

	return y
}

func (t Type[A]) Min(x, y A) A {
	o := t.Compare(x, y)

	if (o == LT{} || o == EQ{}) {
		return x
	}

	return y
}

type Ordering interface {
	isOrdering()
}

type LT struct{}

func (_ LT) isOrdering() {}

type EQ struct{}

func (_ EQ) isOrdering() {}

type GT struct{}

func (_ GT) isOrdering() {}

// Ordered implements CompareFn for natively ordered types.
func Ordered[A constraints.Ordered](x, y A) Ordering {
	if x < y {
		return LT{}
	}

	if x == y {
		return EQ{}
	}

	return GT{}
}

// Ensure Ordering can be used with NewType.
var _ Type[int] = NewType[int](Ordered[int])

type LTEFn[A any] func(A, A) bool

// FromLTE derives a CompareFn using the provide LTEFn.
func FromLTE[A any](lteFn LTEFn[A]) CompareFn[A] {
	return func(x, y A) Ordering {
		lte := lteFn(x, y)
		gte := lteFn(y, x)

		if lte && gte {
			return EQ{}
		}

		if lte {
			return LT{}
		}

		return GT{}
	}
}

// Ensure FromLTE can be used with NewType.
var _ Type[int] = NewType[int](FromLTE[int](func(x, y int) bool { return x <= y }))

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("eq.Conform", func(t *testing.T) {
			eq.Conform[A](c)(t, x, y, z)
		})

		t.Run("comparability", func(t *testing.T) {
			require.True(t, c.LTE(x, y) || c.LTE(y, x))
		})

		t.Run("transitivity", func(t *testing.T) {
			if c.LTE(x, y) && c.LTE(y, z) {
				require.True(t, c.LTE(x, z))
			}
		})

		t.Run("reflexivity", func(t *testing.T) {
			require.True(t, c.LTE(x, x))
		})

		t.Run("antisymmetry", func(t *testing.T) {
			if c.LTE(x, y) && c.LTE(y, x) {
				require.True(t, c.Equal(x, y))
			}
		})

		require.True(t, c.GTE(x, y) == c.LTE(y, x))
		require.True(t, c.LT(x, y) == (c.LTE(x, y) && c.NE(x, y)))
		require.True(t, c.GT(x, y) == c.LT(y, x))
		require.True(t, c.LT(x, y) == (c.Compare(x, y) == LT{}))
		require.True(t, c.GT(x, y) == (c.Compare(x, y) == GT{}))
		require.True(t, c.Equal(x, y) == (c.Compare(x, y) == EQ{}))

		if c.LTE(x, y) {
			require.Equal(t, c.Min(x, y), x)
		} else {
			require.Equal(t, c.Min(x, y), y)
		}

		if c.GTE(x, y) {
			require.Equal(t, c.Max(x, y), x)
		} else {
			require.Equal(t, c.Max(x, y), y)
		}
	}
}
