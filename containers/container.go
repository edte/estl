// @program:     tiny-stl
// @file:        container.go
// @author:      edte
// @create:      2021-12-30 10:12
// @description:
package containers

import (
	"stl/iterator"
)

type Container interface {
	// 返回容器第一个元素的迭代器指针；
	Begin() iterator.Iterator

	// 返回容器最后一个元素后面一位的迭代器指针
	End() iterator.Iterator

	// 删除容器中的所有的元素
	Clear()

	Size() int

	Empty() bool
}
