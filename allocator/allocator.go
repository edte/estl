// @program:     tiny-stl
// @file:        allocator.go
// @author:      edte
// @create:      2022-01-23 23:07
// @description:
package allocator

// 使用 cgo 调用 c 的 calloc 和 free 函数分配内存
// https://github.com/dgraph-io/ristretto/tree/master/contrib/memtest
// https://blog.csdn.net/Rong_Toa/article/details/110095720

type Allocator interface {
	// Construct 默认构造函数
	Construct(ptr uintptr)

	// Destroy 析构函数
	Destroy(ptr uintptr)

	// Allocate 配置空间，足以存储n个T对象。第二个参数是个提示。实现上可能会利用它来增进区域性(locality)，或完全忽略之
	Allocate()

	// Deallocate 释放先前配置的空间
	Deallocate(ptr uintptr)

	// Address 返回某个对象的地址，a.address(x)等同于&x
	Address() uintptr

	// Maxsize 返回可成功配置的最大量
	Maxsize()
}
