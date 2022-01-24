// @program:     tiny-stl
// @file:        iterator.go
// @author:      edte
// @create:      2022-01-16 20:30
// @description:
package algorithm

import (
	"stl/containers"
	"stl/iterator"
)

// Advance advances the iterator it by n element positions.
// 原地前进 n 步
func Advance(it iterator.Iterator, n int) {
	i, ok := it.(iterator.RandomAccessIterator)
	if ok {
		i.NextN(n)
		return
	}

	for n != 0 {
		it = it.Next()
		n--
	}
}

// Distance return Distance between iterators
func Distance(first, last iterator.Iterator) int {
	_, ok := first.(iterator.RandomAccessIterator)
	if ok {
		return last.(iterator.RandomAccessIterator).Position() - first.(iterator.RandomAccessIterator).Position()
	}

	dis := 0

	for i := first.Clone(); !i.Equal(last); i.Next() {
		dis++
	}

	return dis
}

// Begin iterator to beginning
// Returns an iterator pointing to the first element in the sequence:
func Begin(container containers.Container) iterator.Iterator {
	switch container.(type) {
	case containers.Vector:
		return container.(containers.Vector).Begin()
	case containers.List:
		return container.(containers.List).Begin()
	case containers.ForwardList:
		return container.(containers.ForwardList).Begin()
	case containers.Deque:
		return container.(containers.Deque).Begin()
	case containers.String:
		return container.(containers.String).Begin()
	case containers.AssociativeContainer:
		return container.(containers.AssociativeContainer).Begin()
	case containers.UnorderedAssociativeContainer:
		return container.(containers.UnorderedAssociativeContainer).Begin()
	default:
		return nil
	}
}

// End iterator to End
// Returns an iterator pointing to the past-the-end element in the sequence:
func End(container containers.Container) iterator.Iterator {
	switch container.(type) {
	case containers.Vector:
		return container.(containers.Vector).End()
	case containers.List:
		return container.(containers.List).End()
	case containers.ForwardList:
		return container.(containers.ForwardList).End()
	case containers.Deque:
		return container.(containers.Deque).End()
	case containers.String:
		return container.(containers.String).End()
	case containers.AssociativeContainer:
		return container.(containers.AssociativeContainer).End()
	case containers.UnorderedAssociativeContainer:
		return container.(containers.UnorderedAssociativeContainer).End()
	default:
		return nil
	}
}

func Pre(it iterator.Iterator) iterator.RandomAccessIterator {
	return PreB(it).Clone().(iterator.RandomAccessIterator)
}

func PreN(it iterator.Iterator, n int) iterator.RandomAccessIterator {
	return PreBN(it, n).Clone().(iterator.RandomAccessIterator)
}

// PreB get iterator to previous element
// Returns an iterator pointing to the element that it would be pointing to if advanced -n positions.
func PreB(it iterator.Iterator) iterator.BidirectionalIterator {
	return CloneBidirectional(it).Pre()
}

func PreBN(it iterator.Iterator, n int) iterator.BidirectionalIterator {
	for n > 0 {
		it = PreB(it)
		n--
	}
	return it.Clone().(iterator.BidirectionalIterator)
}

func NextInput(it iterator.Iterator) iterator.InputIterator {
	return CloneInput(it).Next().(iterator.InputIterator)
}

func NextOutput(it iterator.Iterator) iterator.OutputIterator {
	return CloneOutput(it).Next().(iterator.OutputIterator)

}

func NextForward(it iterator.Iterator) iterator.ForwardIterator {
	return CloneForward(it).Next().(iterator.ForwardIterator)

}

func NextBidirectional(it iterator.Iterator) iterator.BidirectionalIterator {
	return CloneBidirectional(it).Next().(iterator.BidirectionalIterator)
}

// Next get iterator to Next element
// Returns an iterator pointing to the element that it would be pointing to if advanced n positions.
func Next(it iterator.Iterator) iterator.RandomAccessIterator {
	return Clone(it).Next().(iterator.RandomAccessIterator)
}

func NextN(it iterator.Iterator, n int) iterator.RandomAccessIterator {
	for n > 0 {
		it = Next(it)
		n--
	}
	return it.Clone().(iterator.RandomAccessIterator)
}

func CloneInput(iter iterator.Iterator) iterator.InputIterator {
	return iter.Clone().(iterator.InputIterator)
}

func CloneOutput(iter iterator.Iterator) iterator.OutputIterator {
	return iter.Clone().(iterator.OutputIterator)
}

func CloneForward(iter iterator.Iterator) iterator.ForwardIterator {
	return iter.Clone().(iterator.ForwardIterator)
}

func CloneBidirectional(iter iterator.Iterator) iterator.BidirectionalIterator {
	return iter.Clone().(iterator.BidirectionalIterator)
}

func Clone(iter iterator.Iterator) iterator.RandomAccessIterator {
	return iter.Clone().(iterator.RandomAccessIterator)
}

func Value(iter iterator.Iterator) interface{} {
	return iter.(iterator.InputIterator).Value()
}

func ValueAt(iter iterator.Iterator, n int) interface{} {
	return iter.(iterator.RandomAccessIterator).ValueAt(n)
}

func SetValue(iter iterator.Iterator, val interface{}) {
	iter.(iterator.OutputIterator).SetValue(val)
}

func SetValueAt(iter iterator.Iterator, val interface{}, n int) {
	iter.(iterator.RandomAccessIterator).IteratorAt(n).SetValue(val)
}

func IteratorAt(iter iterator.Iterator, n int) iterator.RandomAccessIterator {
	return iter.(iterator.RandomAccessIterator).IteratorAt(n)
}

func Mid(first, last iterator.RandomAccessIterator) iterator.RandomAccessIterator {
	return NextN(first, first.Distance(last)/2)
}

func Position(iter iterator.Iterator) int {
	return iter.(iterator.RandomAccessIterator).Position()
}
