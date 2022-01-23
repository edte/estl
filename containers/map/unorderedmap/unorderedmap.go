// @program:     tiny-stl
// @file:        unorderedmap.go
// @author:      edte
// @create:      2022-01-02 22:25
// @description:
package unorderedmap

import (
	"stl/algorithm/hash"
	"stl/allocator"
	"stl/containers"
	"stl/iterator"
)

type UnorderedMap struct {
}

func (u UnorderedMap) Swap(t *containers.Container) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Clone() containers.Container {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Clear() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Size() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) MaxSize() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Capacity() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Empty() bool {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) GetAllocator() allocator.Allocator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Construct() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Destroy() {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) String() string {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Begin() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) End() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) CBegin() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) CEnd() iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Find(key interface{}) iterator.ForwardIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Count(key interface{}) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) EqualRange(key interface{}) (iterator.ForwardIterator, iterator.ForwardIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) At(key interface{}) interface{} {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) SetAt(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Insert(val interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) InsertKey(key interface{}, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) InsertPosition(position iterator.InputIterator, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) InsertRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Erase(key interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) ErasePosition(position iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) EraseRange(first, last iterator.InputIterator) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Emplace(value interface{}) (iterator.BidirectionalIterator, bool) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) EmplaceHint(position iterator.InputIterator, value interface{}) iterator.BidirectionalIterator {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) BucketCount() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) MaxBucketCount() int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) BucketSize(n int) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Bucket(k int) int {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) LoadFactor() float64 {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) MaxLoadFactor() float64 {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Rehash(n int) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) Reserve(n int) {
	//TODO implement me
	panic("implement me")
}

func (u UnorderedMap) HashFunction() hash.Hash {
	//TODO implement me
	panic("implement me")
}
