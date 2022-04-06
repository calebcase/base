package either

import (
	"github.com/calebcase/base/data"
	"github.com/calebcase/base/data/list"
)

type Class[A, B any] interface {
	NewLeft(A) Left[A, B]
	NewRight(B) Right[A, B]
}

type Type[A, B any] struct{}

// Ensure Type implements Class.
var _ Class[int, string] = Type[int, string]{}

func NewType[A, B any]() Type[A, B] {
	return Type[A, B]{}
}

func (t Type[A, B]) NewLeft(v A) Left[A, B] {
	return Left[A, B]{v}
}

func (t Type[A, B]) NewRight(v B) Right[A, B] {
	return Right[A, B]{v}
}

// Either is the sum type for Either.
type Either[A, B any] interface {
	isEither(A, B)
}

// Left contains left value A.
type Left[A, B any] struct {
	Value A
}

// Ensure Left implements data.Data.
var _ data.Data[int] = Left[int, string]{}

func (l Left[A, B]) isEither(_ A, _ B) {}

func (l Left[A, B]) DValue() A {
	return l.Value
}

func (l Left[A, B]) DRest() data.Data[A] {
	return nil
}

// Right contains the right value.
type Right[A, B any] struct {
	Value B
}

// Ensure Right implements data.Data
var _ data.Data[string] = Right[int, string]{}

func (r Right[A, B]) isEither(_ A, _ B) {}

func (r Right[A, B]) DValue() B {
	return r.Value
}

func (r Right[A, B]) DRest() data.Data[B] {
	return nil
}

// Apply returns the default value `dflt` if `v` is Nothing. Otherwise it
// returns the result of calling `f` on `v`.
func Apply[A, B, C any](fL func(a A) C, fR func(b B) C, v Either[A, B]) C {
	if l, ok := v.(Left[A, B]); ok {
		return fL(l.Value)
	}

	if r, ok := v.(Right[A, B]); ok {
		return fR(r.Value)
	}

	panic("impossible")
}

func Lefts[A, B any](es list.List[Either[A, B]]) (vs list.List[A]) {
	vs = list.List[A]{}

	if es == nil || len(es) == 0 {
		return vs
	}

	if l, ok := es.DValue().(Left[A, B]); ok {
		vs = append(vs, l.Value)
	}

	for rest := es.DRest(); rest != nil; rest = rest.DRest() {
		if l, ok := rest.DValue().(Left[A, B]); ok {
			vs = append(vs, l.Value)
		}
	}

	return vs
}
