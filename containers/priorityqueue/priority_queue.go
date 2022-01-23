// @program:     tiny-stl
// @file:        priorityqueue.go
// @author:      edte
// @create:      2022-01-15 15:59
// @description:
package priorityqueue

import (
	"stl/algorithm"
	"stl/allocator"
	"stl/comparator"
	"stl/containers"
	"stl/containers/vector"
	"stl/locker"
)

// state
const (
	r = iota
	w
	n
)

type Option func(filter *PriorityQueue)

type PriorityQueue struct {
	container containers.SequentialContainer
	cmp       comparator.Comparator
	locker    locker.Locker

	safe  bool
	state int
}

func New(opts ...Option) *PriorityQueue {
	h := &PriorityQueue{
		container: vector.New(),
		cmp:       comparator.NewGreater(),
		locker:    locker.NewMutex(),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithData(data []interface{}) Option {
	return func(v *PriorityQueue) {
		v.container = vector.New(vector.WithData(data))
	}
}

func WithCap(c int) Option {
	return func(v *PriorityQueue) {
		v.container = vector.New(vector.WithCap(c))
	}
}

func WithOperator(c comparator.Comparator) Option {
	return func(v *PriorityQueue) {
		v.cmp = c
	}
}

func WithCmp(cmp func(a interface{}, b interface{}) bool) Option {
	return func(v *PriorityQueue) {
		v.cmp = comparator.New(cmp)
	}
}

func WithContainer(c containers.SequentialContainer) Option {
	return func(v *PriorityQueue) {
		v.container = c
	}
}

func WithGoroutineSafe() Option {
	return func(h *PriorityQueue) {
		h.safe = true
	}
}

func WithLocker(l locker.Locker) Option {
	return func(h *PriorityQueue) {
		h.locker = l
	}
}

func (h *PriorityQueue) Empty() bool {
	return h.Size() == 0
}

func (h *PriorityQueue) Size() int {
	return h.container.Size()
}

func (h *PriorityQueue) Top() interface{} {
	if !h.safe {
		return h.container.At(0)
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	return h.container.At(0)
}

func (h *PriorityQueue) Push(v interface{}) {
	if !h.safe {
		h.state = w
		h.container.PushBack(v)
		algorithm.PushHeap(h.container.Begin(), h.container.End(), h.cmp)
		h.state = n
		return
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	h.state = w
	h.container.PushBack(v)
	algorithm.PushHeap(h.container.Begin(), h.container.End(), h.cmp)
	h.state = n
}

func (h *PriorityQueue) Emplace(v interface{}) {
	h.container.PushBack(v)
	algorithm.PushHeap(h.container.Begin(), h.container.End(), h.cmp)
}

func (h *PriorityQueue) Pop() {
	if !h.safe {
		h.state = w
		algorithm.PopHeap(h.container.Begin(), h.container.End(), h.cmp)
		h.container.PopBack()
		h.state = n
		return
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	h.state = w
	algorithm.PopHeap(h.container.Begin(), h.container.End(), h.cmp)
	h.container.PopBack()
	h.state = n
}

func (h *PriorityQueue) String() string {
	return algorithm.ShowHeap(h.container.Begin(), h.container.End())
}

func (h *PriorityQueue) Swap(t *containers.Container) {
	h.container.Swap(t)
}

func (h *PriorityQueue) Clone() containers.Container {
	return h.container.Clone()
}

func (h *PriorityQueue) Clear() {
	h.Clear()
}

func (h *PriorityQueue) MaxSize() int {
	return h.container.MaxSize()
}

func (h *PriorityQueue) Capacity() int {
	return h.Capacity()
}

func (h *PriorityQueue) GetAllocator() allocator.Allocator {
	return h.container.GetAllocator()
}

func (h *PriorityQueue) Construct() {
	h.container.Construct()
}

func (h *PriorityQueue) Destroy() {
	h.container.Destroy()
}
