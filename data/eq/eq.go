package eq

import (
	"testing"

	"github.com/calebcase/base/data/function"
	"github.com/stretchr/testify/require"
)

type Class[A any] interface {
	Equal(A, A) bool
	NE(A, A) bool
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("reflexivity", func(t *testing.T) {
			require.True(t, c.Equal(x, x))
		})

		t.Run("symmetry", func(t *testing.T) {
			require.True(t, c.Equal(x, y) == c.Equal(y, x))
		})

		t.Run("transitivity", func(t *testing.T) {
			if c.Equal(x, y) && c.Equal(y, z) {
				require.True(t, c.Equal(x, z))
			}
		})

		t.Run("extensionality", func(t *testing.T) {
			if c.Equal(x, y) {
				require.True(t, c.Equal(function.Id(x), function.Id(y)))
			}
		})

		t.Run("negation", func(t *testing.T) {
			require.True(t, c.NE(x, y) == !c.Equal(x, y))
		})
	}
}
