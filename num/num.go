package num

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

type Class[A any] interface {
	Add(A, A) A
	Sub(A, A) A
	Mul(A, A) A

	Negate(A) A
	Abs(A) A
	SigNum(A) A

	// FIXME: FromInteger(integer.Integer) A
}

type Numeric interface {
	constraints.Integer | constraints.Float
}

type Type[A Numeric] struct{}

// Ensure Type implements Class.
var _ Class[int] = Type[int]{}

func NewType[A Numeric]() Type[A] {
	return Type[A]{}
}

func (t Type[A]) Add(x, y A) A {
	return x + y
}

func (t Type[A]) Sub(x, y A) A {
	return x - y
}

func (t Type[A]) Mul(x, y A) A {
	return x * y
}

func (t Type[A]) Negate(x A) A {
	return -x
}

func (t Type[A]) Abs(x A) A {
	var z A

	if x < z {
		return t.Negate(x)
	}

	return x
}

func (t Type[A]) SigNum(x A) A {
	var z A

	switch {
	case x < z:
		return z - 1
	case z < x:
		return z + 1
	}

	return z
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A comparable, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("add", func(t *testing.T) {
			t.Run("associativity", func(t *testing.T) {
				require.True(t, c.Add(c.Add(x, y), z) == c.Add(x, c.Add(y, z)))
			})

			t.Run("commutativity", func(t *testing.T) {
				require.True(t, c.Add(x, y) == c.Add(y, x))
			})

			t.Run("additive identity", func(t *testing.T) {
				var z A

				require.True(t, c.Add(x, z) == x)
			})

			t.Run("additive inverse", func(t *testing.T) {
				var z A

				require.True(t, c.Add(x, c.Negate(x)) == z)
			})
		})

		t.Run("mul", func(t *testing.T) {
			t.Run("associativity", func(t *testing.T) {
				require.True(t, c.Add(c.Add(x, y), z) == c.Add(x, c.Add(y, z)))
			})

			t.Run("multiplicative identity", func(t *testing.T) {
				var z A

				require.True(t, c.Add(x, z) == x)
			})
		})

		t.Run("distributivity", func(t *testing.T) {
			require.True(t, c.Mul(x, c.Add(y, z)) == c.Add(c.Mul(x, y), c.Mul(x, z)))
			require.True(t, c.Mul(c.Add(y, z), x) == c.Add(c.Mul(y, x), c.Mul(z, x)))
		})
	}
}
