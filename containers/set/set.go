// @program:     tiny-stl
// @file:        set.go
// @author:      edte
// @create:      2022-01-02 22:25
// @description:
package set

import (
	"stl/allocator"
	"stl/containers"
	"stl/containers/rbtree"
	"stl/iterator"
)

type Set struct {
	tree rbtree.RBtree
}

func (s Set) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (s Set) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (s Set) Clear() {
	//TODO implement me
	panic("implement me")
}

func (s Set) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s Set) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (s Set) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (s Set) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (s Set) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (s Set) Construct() {
	//TODO implement me
	panic("implement me")
}

func (s Set) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (s Set) String() string {
	//TODO implement me
	panic("implement me")
}

func (s Set) Begin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) End() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) RBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) REnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) CBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) CEnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) CRBegin() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) CREnd() iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) Find(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) Count(key interface{}) int {
	//TODO implement me
	panic("implement me")
}

func (s Set) LowerBound(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) UpperBound(key interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (s Set) EqualRange(key interface{}) (iterator.BidirectionalIterator, iterator.BidirectionalIterator) {
	//TODO implement me
	panic("implement me")
}

func (s Set) At(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (s Set) SetAt(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s Set) Insert(val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s Set) InsertKey(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s Set) InsertPosition(position iterator.InputIterator, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s Set) InsertRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (s Set) Erase(key interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s Set) ErasePosition(position iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (s Set) EraseRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (s Set) Emplace(value interface{}) (iterator.BidirectionalIterator, bool) {
	//TODO implement me
	panic("implement me")
}

func (s Set) EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}
