package data

import "errors"

var ErrNoValue = errors.New("data: no value")

// Data provides a general way to get values from a sum type.
type Data[A any] interface {
	// DValue returns the "first" value v from the sum type. If no value is
	// present it should panic with ErrNoValue.
	DValue() (v A)

	// DRest returns a Data interface that is scoped to the remaining
	// values. If no more values remain, then it should return nil.
	DRest() Data[A]
}
