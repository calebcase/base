package list_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"testing"

	"github.com/calebcase/base/data/list"
)

func i2b(is []int32) (bs []byte) {
	buf := &bytes.Buffer{}

	for _, i := range is {
		err := binary.Write(buf, binary.BigEndian, i)
		if err != nil {
			panic(err)
		}
	}

	return buf.Bytes()
}

func b2i(bs []byte) (is []int32) {
	is = []int32{}
	buf := bytes.NewBuffer(bs)

	for {
		var i int32
		err := binary.Read(buf, binary.BigEndian, &i)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			if errors.Is(err, io.ErrUnexpectedEOF) {
				break
			}

			panic(err)
		}

		is = append(is, i)
	}

	return
}

func FuzzConformInt32(f *testing.F) {
	l := list.NewType[int32]()

	type TC struct {
		x []byte
		y []byte
		z []byte
	}

	tcs := []TC{
		{i2b([]int32{1}), i2b([]int32{2}), i2b([]int32{3})},
		{i2b([]int32{1, 2, 3}), i2b([]int32{1, 2, 3}), i2b([]int32{1, 2, 3})},
		{i2b([]int32{0, 0, 0}), i2b([]int32{0, 0, 0}), i2b([]int32{0, 0, 0})},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(func(t *testing.T, x, y, z []byte) {
		list.Conform[int32](l)(t, b2i(x), b2i(y), b2i(z))
	})
}

func FuzzConformByte(f *testing.F) {
	l := list.NewType[byte]()

	type TC struct {
		x []byte
		y []byte
		z []byte
	}

	tcs := []TC{
		{list.List[byte]{0, 0, 0}, list.List[byte]{1, 1, 1}, list.List[byte]{2, 2, 2}},
		{i2b([]int32{1}), i2b([]int32{2}), i2b([]int32{3})},
		{i2b([]int32{1, 2, 3}), i2b([]int32{1, 2, 3}), i2b([]int32{1, 2, 3})},
		{i2b([]int32{0, 0, 0}), i2b([]int32{0, 0, 0}), i2b([]int32{0, 0, 0})},
	}

	for _, tc := range tcs {
		f.Add(tc.x, tc.y, tc.z)
	}

	f.Fuzz(func(t *testing.T, x, y, z []byte) {
		list.Conform[byte](l)(t, x, y, z)
	})
}
