package semigroup

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type Class[A any] interface {
	SAssoc(A, A) A
}

type SAssocFn[A any] func(A, A) A

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("associativity", func(t *testing.T) {
			require.Equal(t, c.SAssoc(x, c.SAssoc(y, z)), c.SAssoc(c.SAssoc(x, y), z))
		})
	}
}
