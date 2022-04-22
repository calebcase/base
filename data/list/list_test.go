package list_test

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/calebcase/base/data/list"
)

func ExampleMap() {
	fmt.Println(
		list.Map(
			func(x int) int {
				return x + 1
			},
			list.List[int]{1, 2, 3},
		),
	)

	// Output:
	// [2 3 4]
}

func ExampleReverse() {
	fmt.Println(list.Reverse(list.List[int]{}))
	fmt.Println(list.Reverse(list.List[int]{42}))
	fmt.Println(list.Reverse(list.List[int]{2, 5, 7}))

	// Output:
	// []
	// [42]
	// [7 5 2]
}

func ExampleIntersperse() {
	fmt.Println(string(list.Intersperse(',', list.List[byte]("abcde"))))

	// Output:
	// a,b,c,d,e
}

func ExampleIntercalate_listbyte() {
	fmt.Println(string(list.Intercalate(
		list.List[byte](", "),
		list.List[list.List[byte]]{
			list.List[byte]("Lorem"),
			list.List[byte]("ipsum"),
			list.List[byte]("dolor"),
		},
	)))

	// Output:
	// Lorem, ipsum, dolor
}

func ExampleIntercalate_arraybyte() {
	fmt.Println(string(list.Intercalate(
		[]byte(", "),
		[][]byte{
			[]byte("Lorem"),
			[]byte("ipsum"),
			[]byte("dolor"),
		},
	)))

	// Output:
	// Lorem, ipsum, dolor
}

func ExampleTranspose_equal() {
	fmt.Println(list.Transpose([][]int{{1, 2, 3}, {4, 5, 6}}))

	// Output:
	// [[1 4] [2 5] [3 6]]
}

func ExampleTranspose_mixed() {
	fmt.Println(list.Transpose([][]int{{10, 11}, {20}, {}, {30, 31, 32}}))

	// Output:
	// [[10 20 30] [11 31] [32]]
}

func ExampleNonEmptySubsequences() {
	fmt.Println(list.NonEmptySubsequences([]int{1, 2, 3}))

	// Output:
	// [[1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}

func ExampleSubsequences() {
	fmt.Println(list.Subsequences([]int{1, 2, 3}))

	// Output:
	// [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}

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
