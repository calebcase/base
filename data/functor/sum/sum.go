package sum

import "github.com/calebcase/base/data"

type Class interface{}

type Type struct{}

// Ensure Type implements Class.
var _ Class = Type{}

func NewType() Type {
	return Type{}
}

// Sum is the sum type for Sum.
type Sum[F, G, A any] interface {
	isSum(F, G, A)

	data.Data[Sum[F, G, A]]
}

// InL is the left side summation.
type InL[F, G, A any] struct {
	Fn    F
	Value A
}

func (il InL[F, G, A]) isSum(_ F, _ G, _ A) {}

func (il InL[F, G, A]) DValue() Sum[F, G, A] {
	return il
}

func (il InL[F, G, A]) DRest() data.Data[Sum[F, G, A]] {
	return nil
}

// InR is the right side summation.
type InR[F, G, A any] struct {
	Fn    G
	Value A
}

func (ir InR[F, G, A]) isSum(_ F, _ G, _ A) {}

func (ir InR[F, G, A]) DValue() Sum[F, G, A] {
	return ir
}

func (ir InR[F, G, A]) DRest() data.Data[Sum[F, G, A]] {
	return nil
}
