// @program:     tiny-stl
// @file:        queue.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package queue

import (
	"stl/allocator"
	"stl/containers"
	"stl/containers/deque"
	"stl/locker"
)

type Option func(*Queue)

type Queue struct {
	container containers.SequentialContainer
	locker    locker.Locker
}

func New(opts ...Option) *Queue {
	q := &Queue{
		container: deque.New(),
		locker:    nil,
	}

	for _, opt := range opts {
		opt(q)
	}

	return q
}

func WithGoroutineSafe(l locker.Locker) Option {
	return func(q *Queue) {
		q.locker = l
	}
}

func (q *Queue) Swap(t *containers.Container) {
	q.Swap(t)
}

func (q *Queue) Clone() containers.Container {
	return q.container.Clone()
}

func (q *Queue) Clear() {
	q.container.Clear()
}

func (q *Queue) Size() int {
	return q.container.Size()
}

func (q *Queue) MaxSize() int {
	return q.container.MaxSize()
}

func (q *Queue) Capacity() int {
	return q.container.Capacity()
}

func (q *Queue) Empty() bool {
	return q.container.Empty()
}

func (q *Queue) GetAllocator() allocator.Allocator {
	return q.container.GetAllocator()
}

func (q *Queue) Construct() {
	q.container.Construct()
}

func (q *Queue) Destroy() {
	q.container.Destroy()
}

func (q *Queue) String() string {
	return q.container.String()
}

func (q *Queue) Push(v interface{}) {
	q.container.PushBack(v)
}

func (q *Queue) Emplace(v interface{}) {
	q.container.EmplaceBack(v)
}

func (q *Queue) Pop() {
	q.container.PopBack()
}
