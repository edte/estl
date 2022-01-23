// @program:     tiny-stl
// @file:        pair.go
// @author:      edte
// @create:      2022-01-23 03:29
// @description:
package algorithm

import "stl/containers/pair"

func MakePair(first, second interface{}) *pair.Pair {
	return pair.New(first, second)
}
