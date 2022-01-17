// @program:     tiny-stl
// @file:        bitset_test.go.go
// @author:      edte
// @create:      2021-12-25 13:18
// @description:
package bitset

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	b := New()
	b.Set(2)
	fmt.Println(b)
}

func TestNewWithSize(t *testing.T) {
	b := New(WithSize(3))
	b.Set(2)
	fmt.Println(b)
}

func TestNewWithData(t *testing.T) {
	b := New(WithData([]byte{0b1111}))
	b.UnSet(2)
	fmt.Println(b)
}

func TestBitSet_All(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b.All())
	b1 := New(WithData([]byte{0xff}))
	fmt.Println(b1.All())
}

func TestBitSet_Any(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b.Any())
	b1 := New(WithData([]byte{0}))
	fmt.Println(b1.Any())
}

func TestBitSet_Clear(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b)
	b.Clear()
	fmt.Println(b)
}

func TestBitSet_Count(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b.Count())
}

func TestBitSet_Flip(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b)
	b.Flip(2)
	fmt.Println(b)
}

func TestBitSet_None(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b.None())
	fmt.Println(New().None())
}

func TestBitSet_Set(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b)
	b.Set(1)
	fmt.Println(b)
}

func TestBitSet_UnSet(t *testing.T) {
	b := New(WithData([]byte{0b1110101}))
	fmt.Println(b)
	b.UnSet(2)
	fmt.Println(b)
}

func TestBitSet_Size(t *testing.T) {
	b := New(WithData([]byte{0xff, 0x81}))
	fmt.Println(b.Size())
	b1 := New(WithSize(3))
	fmt.Println(b1.Size())
}

func TestBitSet_String(t *testing.T) {
	b := New(WithData([]byte{0xfa}))
	fmt.Println(b)
}

func TestBitSet_ToString(t *testing.T) {

}

func TestBitSet_ToUint32(t *testing.T) {

}

func TestBitSet_ToUint64(t *testing.T) {

}

func TestBitSet_Test(t *testing.T) {
	b := New()
	b.Set(2)
	b.Set(3)
	b.Set(5)
	fmt.Println(b)
	fmt.Println(b.Test(4))
	fmt.Println(b.Test(3))
	fmt.Println(b.Test(7))
	fmt.Println(b.Test(5))
}
