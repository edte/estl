// @program:     tiny-stl
// @file:        mem.go
// @author:      edte
// @create:      2022-01-23 23:47
// @description:

package allocator

//#include <stdlib.h>
import "C"

import (
	"reflect"
	"unsafe"
)

// Allocates size bytes of uninitialized storage.
// If allocation succeeds, returns a pointer that is suitably
// aligned for any object type with fundamental alignment.
// If size is zero, the behavior of malloc is implementation-defined.
// For example, a null pointer may be returned. Alternatively,
// a non-null pointer may be returned; but such a pointer should
// not be dereferenced, and should be passed to free to avoid memory
// leaks.
//  malloc is thread-safe: it behaves as though only accessing the
// memory locations visible through its argument, and not any static
// storage.
// A previous call to free or realloc that deallocates a region of
// memory synchronizes-with a call to malloc that allocates the same
// or a part of the same region of memory. This synchronization occurs
// after any access to the memory by the deallocating function and
// before any access to the memory by malloc. There is a single total
// order of all allocation and deallocation functions operating on
// each particular region of memory.

// @c:
// void *malloc(size_t size)

// @Parameters:
// size − This is the size of the memory block, in bytes.
//
// @Return Value:
// This function returns a pointer to the allocated memory,
// or NULL if the request fails.
func malloc(size int) unsafe.Pointer {
	if size == 0 {
		return nil
	}

	ptr := C.malloc(C.size_t(size))
	if ptr == nil {
		//panic("malloc error")
		return nil
	}

	r := &reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  size,
		Cap:  size,
	}

	return unsafe.Pointer(r)
}

// Allocates memory for an array of num objects of size and initializes
// all bytes in the allocated storage to zero.
// If allocation succeeds, returns a pointer to the lowest (first)
// byte in the allocated memory block that is suitably aligned for any
// object type.
// If size is zero, the behavior is implementation defined (null
// pointer may be returned, or some non-null pointer may be returned
// that may not be used to access storage)
// calloc is thread-safe: it behaves as though only accessing the
// memory locations visible through its argument, and not any static
// storage.
// A previous call to free or realloc that deallocates a region of
// memory synchronizes-with a call to calloc that allocates the same
// or a part of the same region of memory. This synchronization
// occurs after any access to the memory by the deallocating function
// and before any access to the memory by calloc. There is a single
// total order of all allocation and deallocation functions operating
// on each particular region of memory.

// @c:
// void* calloc( size_t num, size_t size );

// @Parameters:
// num	-	number of objects
// size	-	size of each object

// @Return Value:
// On success, returns the pointer to the beginning of newly
// allocated memory. To avoid a memory leak, the returned pointer
// must be deallocated with free() or realloc().
// On failure, returns a null pointer.
func calloc(num, size int) (v unsafe.Pointer) {
	if num == 0 || size == 0 {
		return nil
	}

	ptr := C.calloc(C.size_t(num), C.size_t(size))
	if ptr == nil {
		//panic("calloc error")
		return nil
	}

	r := &reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  num * size,
		Cap:  num * size,
	}

	return unsafe.Pointer(r)
}

// Reallocates the given area of memory. It must be previously
// allocated by malloc(), calloc() or realloc() and not yet freed
// with a call to free or realloc. Otherwise, the results are
// undefined.
// The reallocation is done by either:
// a) expanding or contracting the existing area pointed to by ptr,
// if possible. The contents of the area remain unchanged up to the
// lesser of the new and old sizes. If the area is expanded, the
// contents of the new part of the array are undefined.
// b) allocating a new memory block of size new_size bytes, copying
// memory area with size equal the lesser of the new and the old sizes, and freeing the old block.
// If there is not enough memory, the old memory block is not freed
// and null pointer is returned.
// If ptr is NULL, the behavior is the same as calling malloc(new_size).
// If new_size is zero, the behavior is implementation defined (null
// pointer may be returned (in which case the old memory block may or
// may not be freed), or some non-null pointer may be returned that
// may not be used to access storage). Such usage is deprecated
// (via DR 400). (since C17)
// (until C23)
// If new_size is zero, the behavior is undefined.
// (since C23)
// realloc is thread-safe: it behaves as though only accessing the
// memory locations visible through its argument, and not any static
// storage.
// A previous call to free or realloc that deallocates a region of
// memory synchronizes-with a call to any allocation function,
// including realloc that allocates the same or a part of the same
// region of memory. This synchronization occurs after any access to
// the memory by the deallocating function and before any access to
// the memory by realloc. There is a single total order of all
// allocation and deallocation functions operating on each particular
// region of memory.

// @c:
// void *realloc( void *ptr, size_t new_size );

// @Parameters:
// ptr	    -	pointer to the memory area to be reallocated
// new_size	-	new size of the array in bytes
//
// @Return Value:
// On success, returns the pointer to the beginning of newly allocated
// memory. To avoid a memory leak, the returned pointer must be
// deallocated with free() or realloc(). The original pointer ptr is
// invalidated and any access to it is undefined behavior (even if
// reallocation was in-place).
// On failure, returns a null pointer. The original pointer ptr
// remains valid and may need to be deallocated with free() or
// realloc().
func realloc(p unsafe.Pointer, newSize int) (v unsafe.Pointer) {
	if newSize == 0 {
		return nil
	}

	bytes := *(*[]byte)(p)
	bytes = bytes[:cap(bytes)]
	p = unsafe.Pointer(&bytes[0])

	ptr := C.realloc(p, C.size_t(newSize))
	if ptr == nil {
		//panic("realloc error")
		return nil
	}

	r := &reflect.SliceHeader{
		Data: uintptr(ptr),
		Len:  newSize,
		Cap:  newSize,
	}

	return unsafe.Pointer(r)
}

// Deallocates the space previously allocated by malloc(), calloc(),
// aligned_alloc(), (since C11) or realloc().
// If ptr is a null pointer, the function does nothing.
// The behavior is undefined if the value of ptr does not equal a
// value returned earlier by malloc(), calloc(), realloc(), or aligned
// _alloc() (since C11).
// The behavior is undefined if the memory area referred to by ptr has
// already been deallocated, that is, free() or realloc() has already
// been called with ptr as the argument and no calls to malloc(),
// calloc() or realloc() resulted in a pointer equal to ptr afterwards.
// The behavior is undefined if after free() returns, an access is made
// through the pointer ptr (unless another allocation function
// happened to result in a pointer value equal to ptr)
// free is thread-safe: it behaves as though only accessing the memory
// locations visible through its argument, and not any static storage.
// A call to free that deallocates a region of memory synchronizes-with
// a call to any subsequent allocation function that allocates the same
// or a part of the same region of memory. This synchronization occurs
// after any access to the memory by the deallocating function and
// before any access to the memory by the allocation function. There
// is a single total order of all allocation and deallocation functions
// operating on each particular region of memory.

// @c:
// void free( void* ptr );

// @Parameters:
// ptr	-	pointer to the memory to deallocate

// @Return Value:
// (none)
func free(ptr unsafe.Pointer) {
	bytes := *(*[]byte)(ptr)
	bytes = bytes[:cap(bytes)]
	ptr = unsafe.Pointer(&bytes[0])

	if len(bytes) == 0 || cap(bytes) == 0 {
		return
	}

	C.free(ptr)
}

func freeBytes(bs []byte) {
	if len(bs) == 0 || cap(bs) == 0 {
		return
	}

	bs = bs[:cap(bs)]
	C.free(unsafe.Pointer(&bs[0]))
}

func freeByte(ptr *byte) {
	C.free(unsafe.Pointer(ptr))
}
