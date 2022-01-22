// @program:     tiny-stl
// @file:        modify.go
// @author:      edte
// @create:      2022-01-02 22:42
// @description:
package algorithm

import (
	"crypto/rand"
	"math/big"
	"stl/functional"
	"stl/iterator"
)

// Copy range of elements
func Copy(first, last iterator.InputIterator, result iterator.OutputIterator) iterator.OutputIterator {
	j := result.Clone().(iterator.OutputIterator)
	for i := first.Clone().(iterator.InputIterator); !i.Equal(last); i.Next() {
		j.SetValue(i.Value())
		j.Next()
	}
	return j
}

// CopyN copy certain elements of range
// Copies the first n elements from the range beginning at first
// into the range beginning at result.
// The function returns an iterator to the end of the destination
// range (which points to one past the last element copied).
// If n is negative, the function does nothing.
// If the ranges overlap, some of the elements in the range pointed
// by result may have undefined but valid values.
func CopyN(first iterator.InputIterator, n int, result iterator.OutputIterator) iterator.OutputIterator {
	if n < 0 {
		return result
	}

	res := CloneOutput(result)
	i := CloneInput(first)

	for n > 0 {
		res.SetValue(i.Value())
		Advance(i, 1)
		Advance(res, 1)
		n--
	}

	return res
}

// CopyIf return number of elements in range satisfying condition
func CopyIf(first, last iterator.InputIterator, result iterator.OutputIterator, pred functional.Pred) iterator.OutputIterator {
	out := result.Clone().(iterator.OutputIterator)

	for i := first.Clone().(iterator.InputIterator); !i.Equal(last); i.Next() {
		if pred(i.Value()) {
			out.SetValue(i.Value())
			out.Next()
		}
	}

	return result
}

// CopyBackward opy range of elements backward
func CopyBackward(first, last, result iterator.BidirectionalIterator) iterator.BidirectionalIterator {
	out := PreB(result)

	for i := PreB(last); i.IsBackEqual(first); i.Pre() {
		out.SetValue(i.Value())
		out.Pre()
	}

	return result
}

// Swap exchange values of objects pointed to by two iterators
func Swap(it1, it2 iterator.ForwardIterator) {
	t := it1.Value()
	it1.SetValue(it2.Value())
	it2.SetValue(t)
}

// SwapRanges exchange values of two ranges
func SwapRanges(first1, last1, first2 iterator.ForwardIterator) iterator.ForwardIterator {
	i := CloneForward(first1)
	j := CloneForward(first2)

	for !i.Equal(last1) {
		Swap(i, j)
		i.Next()
		j.Next()
	}
	return j
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

// Replace value in range
func Replace(first, last iterator.ForwardIterator, oldValue, newValue interface{}) {
	ForEachIter(first, last, func(i iterator.ForwardIterator) {
		if functional.EqualTo(i.Value(), oldValue) {
			i.SetValue(newValue)
		}
	})
}

// ReplaceIf replace value in range
// Assigns new_value to all the elements in the range [first,last)
// for which pred returns true.
func ReplaceIf(first, last iterator.ForwardIterator, pred functional.Pred, newValue interface{}) {
	ForEachIter(first, last, func(i iterator.ForwardIterator) {
		if pred(i.Value()) {
			i.SetValue(newValue)
		}
	})
}

// ReplaceCopy opy range replacing value
func ReplaceCopy(first, last iterator.InputIterator, result iterator.OutputIterator, oldValue, newValue interface{}) iterator.OutputIterator {
	i := CloneInput(first)
	out := CloneOutput(result)

	for !i.Equal(last) {
		if functional.EqualTo(i.Value(), oldValue) {
			out.SetValue(newValue)
		} else {
			out.SetValue(i.Value())
		}
		i.Next()
		out.Next()
	}

	return result
}

// ReplaceCopyIf copy range replacing value
func ReplaceCopyIf(first, last iterator.InputIterator, result iterator.OutputIterator, pred functional.Pred, newValue interface{}) iterator.OutputIterator {
	i := CloneInput(first)
	out := CloneOutput(result)

	for !i.Equal(last) {
		if pred(i.Value()) {
			out.SetValue(newValue)
		} else {
			out.SetValue(i.Value())
		}
		i.Next()
		out.Next()
	}

	return result
}

// Fill range with value
// Assigns val to all the elements in the range [first,last).
func Fill(first, last iterator.ForwardIterator, val int) {
	ForEachIter(first, last, func(i iterator.ForwardIterator) {
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

// Generate values for range with function
func Generate(first, last iterator.ForwardIterator, gen functional.Generator) {
	for i := CloneForward(first); !i.Equal(last); i.Next() {
		i.SetValue(gen())
	}
}

// GenerateN generate values for range with function
func GenerateN(first iterator.OutputIterator, n int, gen functional.Generator) {
	i := CloneOutput(first)

	for n > 0 {
		i.SetValue(gen())
		i.Next()
		n--
	}
}

// Rotate left the elements in range
func Rotate(first, middle, last iterator.ForwardIterator) {
	j := CloneForward(middle)
	i := CloneForward(first)
	k := CloneForward(middle)

	for !j.Equal(i) {
		Swap(i, j)
		i.Next()
		j.Next()

		if j.Equal(last) {
			j = CloneForward(k)
		}

		if i.Equal(k) {
			k = CloneForward(j)
		}
	}
}

// RotateCopy rotate left the elements in range
func RotateCopy(first, middle, last iterator.ForwardIterator, result iterator.OutputIterator) iterator.OutputIterator {
	res := Copy(middle, last, result)
	return Copy(first, middle, res)
}

// Shuffle randomly rearrange elements in range using generator
func Shuffle(first, last iterator.RandomAccessIterator) {
	for i := Clone(first); !i.Equal(last); i.Next() {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		j := first.Position() + i.Position() + (int)(n.Int64())%(first.Distance(last)-i.Position()-first.Position())
		Swap(i, Clone(i).IteratorAt(j))
	}
}

func RandomShuffle(first, last iterator.RandomAccessIterator) {
	Shuffle(first, last)
}

// Transform range
func Transform(first1, last1 iterator.InputIterator, result iterator.OutputIterator, op functional.Op) iterator.OutputIterator {
	i := CloneInput(first1)
	out := CloneOutput(result)

	for !i.Equal(last1) {
		out.SetValue(op(i.Value()))
		i.Next()
		out.Next()
	}

	return result
}

// Move range of elements
func Move(first, last iterator.InputIterator, result iterator.OutputIterator) iterator.OutputIterator {
	i := CloneInput(first)
	j := CloneOutput(result)

	for !i.Equal(last) {
		j.SetValue(i.Value())
		i.Next()
		j.Next()
	}

	return j
}

// Remove value from range
func Remove(first, last iterator.ForwardIterator, val interface{}) iterator.ForwardIterator {
	return nil
}

// RemoveIf remove elements from range
func RemoveIf(first, last iterator.ForwardIterator, pred functional.Pred) iterator.ForwardIterator {
	return nil
}

// RemoveCopy copy range removing value
func RemoveCopy(first, last iterator.InputIterator, result iterator.OutputIterator, val interface{}) iterator.OutputIterator {
	i := CloneInput(first)
	out := CloneOutput(result)

	for !i.Equal(last) {
		if !(i.Value() == val) {
			out.SetValue(i.Value())
			out.Next()
		}
		i.Next()
	}

	return result
}

// RemoveCopyIf copy range removing values
func RemoveCopyIf(first, last iterator.InputIterator, result iterator.OutputIterator, pred functional.Pred) iterator.OutputIterator {
	i := CloneInput(first)
	out := CloneOutput(result)

	for !i.Equal(last) {
		if !(pred(i.Value())) {
			out.SetValue(i.Value())
			out.Next()
		}
		i.Next()
	}

	return result
}

// Unique remove consecutive duplicates in range
func Unique(first, last iterator.ForwardIterator, preds ...functional.BinaryPred) iterator.ForwardIterator {
	pred := functional.NotEqualTo
	if len(preds) != 0 {
		pred = preds[0]
	}
	if first.Equal(last) {
		return last
	}

	res := CloneForward(first)

	for i := NextForward(first); !i.Equal(last); i.Next() {
		if pred(res.Value(), i.Value()) {
			res.Next()
			res.SetValue(i.Value())
		}
	}

	return NextForward(res)
}

// UniqueCopy copy range removing duplicates
func UniqueCopy(first, last iterator.InputIterator, result iterator.OutputIterator, preds ...functional.BinaryPred) iterator.OutputIterator {
	pred := functional.NotEqualTo
	if len(preds) != 0 {
		pred = preds[0]
	}
	if first.Equal(last) {
		return result
	}

	result.SetValue(first.Value())

	res := CloneForward(result)

	for i := NextForward(first); !i.Equal(last); i.Next() {
		if pred(res.Value(), i.Value()) {
			res.Next()
			res.SetValue(i.Value())
		}
	}

	return result
}
