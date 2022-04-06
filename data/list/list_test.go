package list_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/calebcase/base/data/eq"
	"github.com/calebcase/base/data/list"
)

func i2b(is []int) (bs []byte) {
	buf := bytes.NewBuffer(bs)

	for _, i := range is {
		binary.Write(buf, binary.BigEndian, i)
	}

	return
}

func b2i(bs []byte) (is []int) {
	buf := bytes.NewBuffer(bs)

	for {
		var i int
		err := binary.Read(buf, binary.BigEndian, &i)
		if err != nil {
			break
		}

		is = append(is, i)
	}

	return
}

func FuzzConformInt(f *testing.F) {
	l := list.NewType[int](eq.NewType(eq.Comparable[int]))

	type TC struct {
		x []byte
		y []byte
		z []byte
	}

	tcs := []TC{
		{i2b([]int{1}), i2b([]int{2}), i2b([]int{3})},
		{i2b([]int{1, 2, 3}), i2b([]int{1, 2, 3}), i2b([]int{1, 2, 3})},
		{i2b([]int{0, 0, 0}), i2b([]int{0, 0, 0}), i2b([]int{0, 0, 0})},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(func(t *testing.T, x, y, z []byte) {
		list.Conform[int](l)(t, b2i(x), b2i(y), b2i(z))
	})
}

func FuzzConformByte(f *testing.F) {
	l := list.NewType[byte](eq.NewType(eq.Comparable[byte]))

	type TC struct {
		x []byte
		y []byte
		z []byte
	}

	tcs := []TC{
		{i2b([]int{1}), i2b([]int{2}), i2b([]int{3})},
		{i2b([]int{1, 2, 3}), i2b([]int{1, 2, 3}), i2b([]int{1, 2, 3})},
		{i2b([]int{0, 0, 0}), i2b([]int{0, 0, 0}), i2b([]int{0, 0, 0})},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(func(t *testing.T, x, y, z []byte) {
		list.Conform[byte](l)(t, x, y, z)
	})
}
