// @program:     tiny-stl
// @file:        modify.go
// @author:      edte
// @create:      2022-01-02 22:42
// @description:
package algorithm

import (
	"crypto/rand"
	"math/big"

	"stl/iterator"
)

func Swap(it1, it2 iterator.RandomAccessIterator) {
	t := it1.Value()
	it1.SetValue(it2.Value())
	it2.SetValue(t)
}

func Merge() {

}

// Reverse range
func Reverse(first, last iterator.RandomAccessIterator) {
	i := first.Clone()
	j := last.Clone().Pre()

	for i.IsFront(j) {
		Swap(i, j)
		i.Next()
		j.Pre()
	}
}

// ReverseCopy copy range reversed
func ReverseCopy(first, last iterator.BidirectionalIterator, result iterator.OutputIterator) {
	i := last.Clone()
	r := result.Clone()

	for !i.Equal(first) {
		i.Pre()
		r.SetValue(i.Value())
		r.Next()
	}
}

func Replace() {

}

func Generate() {

}

// Fill range with value
// Assigns val to all the elements in the range [first,last).
func Fill(first, last iterator.ForwardIterator, val int) {
	for i := first.Clone(); !i.Equal(last); i.Next() {
		i.SetValue(val)
	}
}

func Rotate() {

}

// Shuffle randomly rearrange elements in range using generator
func Shuffle(first, last iterator.RandomAccessIterator) {
	for i := first.Clone(); !i.Equal(last); i.Next() {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		j := first.Position() + i.Position() + (int)(n.Int64())%(first.Distance(last)-i.Position()-first.Position())
		Swap(i, i.Clone().IteratorAt(j))
	}
}

func RandomShuffle(first, last iterator.RandomAccessIterator) {

}

func transform() {

}

func Move() {

}
