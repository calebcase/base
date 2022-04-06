package monoid

import (
	"testing"

	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/semigroup"
	"github.com/stretchr/testify/require"
)

type Class[A any] interface {
	semigroup.Class[A]

	MEmpty() A
	MAppend(A, A) A
	MConcat([]A) A
}

type Type[A any] struct {
	eq.Type[A]

	equalFn  eq.EqualFn[A]
	sAssocFn semigroup.SAssocFn[A]
	mEmptyFn MEmptyFn[A]
}

type MEmptyFn[A any] func() A

func NewType[A any](
	equalFn eq.EqualFn[A],
	sAssocFn semigroup.SAssocFn[A],
	mEmptyFn MEmptyFn[A],
) Type[A] {
	return Type[A]{
		Type: eq.NewType(equalFn),

		sAssocFn: sAssocFn,
		mEmptyFn: mEmptyFn,
	}
}

func (t Type[A]) SAssoc(x, y A) A {
	return t.sAssocFn(x, y)
}

func (t Type[A]) MEmpty() A {
	return t.mEmptyFn()
}

func (t Type[A]) MAppend(x, y A) A {
	return t.SAssoc(x, y)
}

func (t Type[A]) MConcat(xs []A) A {
	// FIXME: Use `foldr mappend mempty`? The below effectively uses
	// `foldl'`, but that's the right thing to use for a finite strict
	// structure. This is part of the lazy/strict mismatch...

	r := t.MEmpty()

	for _, x := range xs {
		r = t.MAppend(r, x)
	}

	return r
}

// Conform returns a function testing if the implementation abides by its laws.
func Conform[A any, CA Class[A]](c CA) func(t *testing.T, x, y, z A) {
	return func(t *testing.T, x, y, z A) {
		t.Run("semigroup.Conform", func(t *testing.T) {
			semigroup.Conform[A](c)(t, x, y, z)
		})

		t.Run("right identity", func(t *testing.T) {
			require.True(t, c.Equal(
				c.SAssoc(x, c.MEmpty()),
				x,
			))
		})

		t.Run("left identity", func(t *testing.T) {
			require.True(t, c.Equal(
				c.SAssoc(c.MEmpty(), x),
				x,
			))
		})

		// FIXME: Needs a working version of foldr and/or a resolution
		// to the lazy/strict mismatch.
		/*
			t.Run("concatenation", func(t *testing.T) {
				l := [x, y, z]
				require.True(t, m.Concat(l) == FoldR(c.SAssoc, c.MEmpty, l))
			})
		*/
	}
}
