// @program:     tiny-stl
// @file:        unorderedset.go
// @author:      edte
// @create:      2022-01-02 22:25
// @description:
package unorderedset

import (
	"stl/algorithm/hash"
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
)

type UnorderedSet struct {
}

func (u UnorderedSet) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Clear() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Size() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Construct() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) String() string {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Begin() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) End() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) CBegin() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) CEnd() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Find(key interface{}) iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Count(key interface{}) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) EqualRange(key interface{}) (iterator.ForwardIterator, iterator.ForwardIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) At(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) SetAt(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Insert(val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) InsertKey(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) InsertPosition(position iterator.InputIterator, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) InsertRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Erase(key interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) ErasePosition(position iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) EraseRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Emplace(value interface{}) (iterator.BidirectionalIterator, bool) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) BucketCount() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) MaxBucketCount() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) BucketSize(n int) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Bucket(k int) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) LoadFactor() float64 {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) MaxLoadFactor() float64 {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Rehash(n int) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedSet) HashFunction() hash.Hash {
	//TODO implement me
	panic("implement me")
}
