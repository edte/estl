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

func Swap(it1, it2 iterator.ForwardIterator) {
	t := it1.Value()
	it1.SetValue(it2.Value())
	it2.SetValue(t)
}

func Merge() {

}

// Reverse range
func Reverse(first, last iterator.BidirectionalIterator) {
	i := first.Clone().(iterator.BidirectionalIterator)
	j := last.Clone().(iterator.BidirectionalIterator).Pre()

	for i.IsFront(j) {
		Swap(i, j)
		i.Next()
		j.Pre()
	}
}

// ReverseCopy copy range reversed
func ReverseCopy(first, last iterator.BidirectionalIterator, result iterator.OutputIterator) {
	i := last.Clone().(iterator.BidirectionalIterator)
	r := result.Clone().(iterator.OutputIterator)

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
	ForEachW(first, last, func(i iterator.ForwardIterator) {
		i.SetValue(val)
	})
}

// FillN fill sequence with value
// Assigns val to the first n elements of the sequence pointed by first.
func FillN(first iterator.OutputIterator, n int, val int) iterator.OutputIterator {
	i := CloneOutput(first)

	for n > 0 {
		i.SetValue(val)
		Advance(i, 1)
		n--
	}

	return first
}

func Rotate() {

}

// Shuffle randomly rearrange elements in range using generator
func Shuffle(first, last iterator.RandomAccessIterator) {
	for i := first.Clone().(iterator.RandomAccessIterator); !i.Equal(last); i.Next() {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		j := first.Position() + i.Position() + (int)(n.Int64())%(first.Distance(last)-i.Position()-first.Position())
		Swap(i, i.Clone().(iterator.RandomAccessIterator).IteratorAt(j))
	}
}

func RandomShuffle(first, last iterator.RandomAccessIterator) {

}

func Transform() {

}

func Move() {

}
