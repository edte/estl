// @program:     tiny-stl
// @file:        container.go
// @author:      edte
// @create:      2021-12-30 10:12
// @description:
package containers

type Container interface {
	// 返回容器第一个元素的迭代器指针；
	Begin()
	// 返回容器最后一个元素后面一位的迭代器指针
	End()

	// 删除容器中的所有的元素
	Clear()
	// 删除迭代器指针 it 处元素
	Erase(it int)

	Size() int

	MaxSize() int

	Swap()

	Empty()

	Operator(a interface{}, b interface{}) bool
}
