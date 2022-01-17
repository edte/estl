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
func Advance(it iterator.InputIterator, n int) {
	for n != 0 {
		it = it.Next()
	}
}

// Distance return Distance between iterators
func Distance(first, last iterator.RandomAccessIterator) int {
	dis := 0

	for i := first.Clone(); !i.Equal(last); i.Next() {
		dis++
	}

	return dis
}

// Begin iterator to beginning
// Returns an iterator pointing to the first element in the sequence:
func Begin(container containers.Container) iterator.Iterator {
	return container.Begin()
}

// End iterator to End
// Returns an iterator pointing to the past-the-end element in the sequence:
func End(container containers.Container) iterator.Iterator {
	return container.End()
}

// Prev get iterator to previous element
// Returns an iterator pointing to the element that it would be pointing to if advanced -n positions.
func Prev(it iterator.BidirectionalIterator) iterator.BidirectionalIterator {
	return it.Clone().Pre()
}

func PrevN(it iterator.BidirectionalIterator, n int) iterator.BidirectionalIterator {
	for n > 0 {
		it = it.Clone().Pre()
		n--
	}
	return it
}

// Next get iterator to Next element
// Returns an iterator pointing to the element that it would be pointing to if advanced n positions.
func Next(it iterator.BidirectionalIterator) iterator.BidirectionalIterator {
	return it.Clone().Next()
}

func NextN(it iterator.BidirectionalIterator, n int) iterator.BidirectionalIterator {
	for n > 0 {
		it = it.Clone().Next()
		n--
	}
	return it
}
