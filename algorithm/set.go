// @program:     tiny-stl
// @file:        set.go
// @author:      edte
// @create:      2022-01-01 22:04
// @description:
package algorithm

import (
	"stl/comparator"
	"stl/containers/vector"
	"stl/iterator"
)

// 集合算法

// SetUnion Union of two sorted ranges
func SetUnion(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) iterator.OutputIterator {
	return result
}

// SetIntersection Intersection of two sorted ranges
func SetIntersection(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) iterator.OutputIterator {

	return result
}

// SetDifference Difference of two sorted ranges
func SetDifference(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) iterator.OutputIterator {

	return result
}

// SetSymmetricDifference Symmetric difference of two sorted ranges
func SetSymmetricDifference(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) iterator.OutputIterator {

	return result
}

// SetCartesianProduct Cartesian Product
func SetCartesianProduct(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) iterator.OutputIterator {

	return result
}

// Merge sorted ranges
// Combines the elements in the sorted ranges [first1,last1) and
// [first2,last2), into a new range beginning at result with all its
// elements sorted.
//
// The elements are compared using operator< for the first version,
// and comp for the second. The elements in both ranges shall already
// be ordered according to this same criterion (operator< or comp).
// The resulting range is also sorted according to this.
func Merge(first1, last1, first2, last2 iterator.InputIterator, result iterator.OutputIterator, cmp ...comparator.Comparator) {
	c := comparator.Default(cmp...)
	i := CloneInput(first1)
	j := CloneInput(first2)
	out := CloneOutput(result)

	for {
		if i.Equal(last1) {
			Copy(j, last2, out)
			break
		}
		if j.Equal(last2) {
			Copy(i, last1, out)
			break
		}
		if c.Operator(i.Value(), j.Value()) {
			out.SetValue(j.Value())
			j.Next()
		} else {
			out.SetValue(i.Value())
			i.Next()
		}
		out.Next()
	}
}

// InplaceMerge merge consecutive sorted ranges
// Merges two consecutive sorted ranges: [first,middle) and
// [middle,last), putting the result into the combined sorted range
// [first,last).
func InplaceMerge(first, middle, last iterator.BidirectionalIterator, cmp ...comparator.Comparator) {
	d := Distance(first, last)

	tmp := vector.New(vector.WithCap(d))

	Merge(first, middle, middle, last, tmp.Begin(), cmp...)

	Copy(tmp.Begin(), tmp.End(), first)
}

// Includes Test whether sorted range includes another sorted range
// Returns true if the sorted range [first1,last1) contains all
// the elements in the sorted range [first2,last2).
// The elements are compared using operator< for the first version,
// and comp for the second. Two elements, a and b are considered
// equivalent if (!(a<b) && !(b<a)) or if (!comp(a,b) && !comp(b,a)).
// The elements in the range shall already be ordered according to
// this same criterion (operator< or comp).
func Includes(first1, last1, first2, last2 iterator.InputIterator, cmp ...comparator.Comparator) bool {
	c := comparator.Default(cmp...)

	i := CloneInput(first1)
	j := CloneInput(first2)

	for !i.Equal(last1) && !j.Equal(last2) {
		if i.Equal(last1) || c.Operator(j.Value(), i.Value()) {
			return false
		}

		if c.Operator(i.Value(), j.Value()) {
			j.Next()
			continue
		}

		i.Next()
	}

	return true
}
