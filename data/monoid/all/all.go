package all

import (
	"github.com/calebcase/base/data/monoid"
)

type Type struct {
	monoid.Type[All]
}

func NewType() Type {
	return Type{
		Type: monoid.NewType(
			func(x, y All) All {
				return x && y
			},
			func() All {
				return true
			},
		),
	}
}

type All = bool
