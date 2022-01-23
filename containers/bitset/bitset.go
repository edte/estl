// @program:     tiny-stl
// @file:        bitset.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package bitset

import (
	"fmt"
)

//Bitset
//A bitset stores bits (elements with only two possible values: 0 or 1, true or false, ...).
//
//The class emulates an array of bool elements, but optimized for space allocation: generally,
//each element occupies only one bit (which, on most systems, is eight times less than the smallest elemental type: char).
//
//Each bit position can be accessed individually: for example,
//for a given bitset named foo, the expression foo[3] accesses its fourth bit,
//just like a regular array accesses its elements. But because no elemental type is a single bit in most C++ environments,
//the individual elements are accessed as special references type (see bitset::reference).
//
//Bitsets have the feature of being able to be constructed from and converted to both integer values and binary strings (see its constructor and members to_ulong and to_string).
//They can also be directly inserted and extracted from streams in binary format (see applicable operators).
//
//The size of a bitset is fixed at compile-time (determined by its template parameter).
//For a class that also optimizes for space allocation and allows for dynamic resizing, see the bool specialization of vector (vector<bool>).

type BitSet struct {
	size uint
	data []byte
}

type Option func(*BitSet)

func New(opts ...Option) *BitSet {
	s := &BitSet{
		size: 32,
		data: make([]byte, 1),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithSize(size uint) Option {
	return func(set *BitSet) {
		set.size = size
		// size>>3+size&7 获取容量，即 size/8+%8
		// 对 2^n 取余 = & 2^n-1
		var d uint = 0
		if (size & 7) != 0 {
			d = 1
		}
		set.data = make([]byte, size>>3+d)
	}
}

func WithData(data []byte) Option {
	return func(set *BitSet) {
		set.data = data
		set.size = uint(len(data) * 8)
	}
}

// Count bits set
// Returns the number of bits in the bitset that are set (i.e., that have a value of one).
func (b *BitSet) Count() (res uint) {
	for _, d := range b.data {
		for i := 0; i < 8; i++ {
			res += (uint)((d >> i) & 1)
		}
	}
	return
}

// Size return size
// Returns the number of bits in the bitset.
func (b *BitSet) Size() uint {
	return b.size
}

// Test return bit value
// Returns whether the bit at position pos is set (i.e., whether it is one).
// true if the bit at position pos is set, and false if it is not set.
func (b *BitSet) Test(u uint) bool {
	_ = b.data[u>>3]
	return ((b.data[u>>3] >> (u & 7)) & 1) == 1
}

// Any rest if any bit is set
// Returns whether any of the bits is set (i.e., whether at least one bit in the bitset is set to one).
// This is the opposite of None.
// true if any of the bits in the bitset is set (to one), and false otherwise.
func (b *BitSet) Any() bool {
	return b.Count() > 0
}

// None test if no bit is set
// Returns whether none of the bits is set (i.e., whether all bits in the bitset have a value of zero).
// This is the opposite of Any.
// true if none of the bits in the bitset is set (to one), and false otherwise.
func (b *BitSet) None() bool {
	return b.Count() == 0
}

// All Test if all bits are set
// Returns whether all the bits in the bitset are set (to one).
// For this function to return true, all bits up to the bitset size shall be set.
// true if all the bits in the bitset are set (to one), and false otherwise.
func (b *BitSet) All() bool {
	return b.Count() == b.Size()
}

// Set bits
func (b *BitSet) Set(pos uint) {
	// 边界检查
	_ = b.data[pos>>3]
	b.data[pos>>3] |= 1 << (pos & 7)
}

// UnSet bits
func (b *BitSet) UnSet(pos uint) {
	_ = b.data[pos>>3]
	b.data[pos>>3] &= ^(1 << (pos & 7))
}

// Clear all bits
func (b *BitSet) Clear() {
	b.data = make([]byte, uint(len(b.data)))
}

// Flip bits
func (b *BitSet) Flip(pos uint) {
	_ = b.data[pos>>3]
	b.data[pos>>3] ^= 1 << (pos & 7)
}

// ToUint64 Convert to uint64
func (b *BitSet) ToUint64() uint64 {
	_ = b.data[7]
	return uint64(b.data[7]) | uint64(b.data[6])<<8 | uint64(b.data[5])<<16 | uint64(b.data[4])<<24 |
		uint64(b.data[3])<<32 | uint64(b.data[2])<<40 | uint64(b.data[1])<<48 | uint64(b.data[0])<<56
}

// ToUint32 Convert to uint32
func (b *BitSet) ToUint32() uint32 {
	_ = b.data[3]
	return uint32(b.data[3]) | uint32(b.data[2])<<8 | uint32(b.data[1])<<16 | uint32(b.data[0])<<24
}

// ToString Convert to string
func (b *BitSet) ToString() (s string) {
	size := len(b.data)
	if (b.size & 7) != 0 {
		size--
	}

	for i := 0; i < size; i++ {
		for j := 0; j < 8; j++ {
			s += fmt.Sprint((b.data[i] >> j) & 1)
		}
	}

	for i := 0; i < (int)(b.size&7); i++ {
		s += fmt.Sprint((b.data[len(b.data)-1] >> i) & 1)
	}

	return s
}

func (b *BitSet) String() string {
	return b.ToString()
}
