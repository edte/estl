// @program:     tiny-stl
// @file:        BestElement.go
// @author:      edte
// @create:      2021-12-30 16:11
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/functional"
	"stl/iterator"
)

// Max return the largest
func Max(data ...interface{}) interface{} {
	return data[0]
}

// Min return the smallest
func Min(data ...interface{}) interface{} {
	return data[0]
}

// MaxMin return smallest and largest elements
func MaxMin(data ...interface{}) (interface{}, interface{}) {
	return Max(data...), Min(data...)
}

// BestElement 取最值
func BestElement(first, last iterator.RandomAccessIterator, cmp comparator.Comparator) (res iterator.RandomAccessIterator) {
	res = Clone(first)
	for l := Clone(first); !l.Equal(last); l.Next() {
		if cmp.Operator(l.Value(), res.Value()) {
			res = Clone(l)
		}
	}
	return res
}

// MinElement return smallest element in range
func MinElement(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) (res iterator.RandomAccessIterator) {
	var c comparator.Comparator = comparator.NewLess()
	if len(cmp) != 0 {
		c = cmp[0]
	}
	return BestElement(first, last, c)
}

// MaxElement return largest element in range
func MaxElement(first, last iterator.RandomAccessIterator, cmp ...comparator.Comparator) (res iterator.RandomAccessIterator) {
	c := comparator.Default(cmp...)
	return BestElement(first, last, c)
}

// MaxMinElement return smallest and largest elements in range
func MaxMinElement(first, last iterator.RandomAccessIterator) (min, max iterator.RandomAccessIterator) {
	return MinElement(first, last), MaxElement(first, last)
}

// Accumulate values in range
func Accumulate(first, last iterator.RandomAccessIterator, ops ...functional.BinaryOp) interface{} {
	op := functional.Add

	if len(ops) != 0 {
		op = ops[0]
	}

	res := 0

	for i := Clone(first); !i.Equal(last); i.Next() {
		res = op(res, i.Value().(int))
	}

	return res
}

// AdjacentDifference compute adjacent difference of range
func AdjacentDifference(first, last iterator.InputIterator, result iterator.OutputIterator, ops ...functional.BinaryOp) iterator.OutputIterator {
	op := functional.Minus

	if len(ops) != 0 {
		op = ops[0]
	}
	r := CloneOutput(result)
	r.SetValue(first.Value())

	for i := CloneInput(first); !CloneInput(i).Next().Equal(last); i.Next() {
		r.Next()
		r.SetValue(op(CloneInput(i).Next().(iterator.InputIterator).Value().(int), i.Value().(int)))
	}

	return result
}

// InnerProduct compute cumulative inner product of range
func InnerProduct(first1, last1, first2 iterator.InputIterator, init int, ops ...functional.BinaryOp) int {
	op1 := functional.Add
	op2 := functional.Multiplies

	if len(ops) != 0 {
		op1 = ops[0]
		op2 = ops[1]
	}

	j := first2.Clone().(iterator.InputIterator)
	for i := first1.Clone().(iterator.InputIterator); !i.Equal(last1); i.Next() {
		init = op1(init, op2(i.Value().(int), j.Value().(int)))
		j.Next()
	}

	return init
}

// PartialSum compute partial sums of range
func PartialSum(first, last iterator.InputIterator, result iterator.OutputIterator, ops ...functional.BinaryOp) iterator.OutputIterator {
	op := functional.Add

	if len(ops) != 0 {
		op = ops[0]
	}

	result.SetValue(first.Value())
	j := result.Clone().(iterator.OutputIterator)

	t := first.Value().(int)

	for i := first.Clone().Next().(iterator.InputIterator); !i.Equal(last); i.Next() {
		j.Next()
		t = op(t, i.Value().(int))
		j.SetValue(t)
	}

	return result
}

// Itoa store increasing sequence
// Assigns to every element in the range [first,last)
// successive values of val, as if incremented with ++val
// after each element is written.
func Itoa(first, last iterator.RandomAccessIterator, val int) {
	for i := first.Clone().(iterator.RandomAccessIterator); !i.Equal(last); i.Next() {
		i.SetValue(val)
		val++
	}
}
