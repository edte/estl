// @program:     tiny-stl
// @file:        heap.go
// @author:      edte
// @create:      2022-01-15 15:59
// @description:
package heap

import (
	"stl/algorithm"
	"stl/comparator"
	"stl/containers/vector"
	"stl/locker"
)

// state
const (
	r = iota
	w
	n
)

type heap interface {
	Empty() bool
	Size() int
	Top() interface{}
	Push(v interface{})
	Emplace(v interface{})
	Pop()
	Swap(t *Heap)
}

type Option func(filter *Heap)

type Heap struct {
	data   *vector.Vector
	cmp    comparator.Comparator
	locker locker.Locker

	safe  bool
	state int
}

func New(opts ...Option) *Heap {
	h := &Heap{
		data:   vector.New(),
		cmp:    comparator.NewGreater(),
		locker: locker.NewMutex(),
	}

	for _, opt := range opts {
		opt(h)
	}

	return h
}

func WithData(data []interface{}) Option {
	return func(v *Heap) {
		v.data = vector.New(vector.WithData(data))
	}
}

func WithCap(c int) Option {
	return func(v *Heap) {
		v.data = vector.New(vector.WithCap(c))
	}
}

func WithOperator(c comparator.Comparator) Option {
	return func(v *Heap) {
		v.cmp = c
	}
}

func WithCmp(cmp func(a interface{}, b interface{}) bool) Option {
	return func(v *Heap) {
		v.cmp = comparator.New(cmp)
	}
}

func WithGoroutineSafe() Option {
	return func(h *Heap) {
		h.safe = true
	}
}

func WithLocker(l locker.Locker) Option {
	return func(h *Heap) {
		h.locker = l
	}
}

func (h *Heap) Empty() bool {
	return h.Size() == 0
}

func (h *Heap) Size() int {
	return h.data.Size()
}

func (h *Heap) Top() interface{} {
	if !h.safe {
		return h.data.At(0)
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	return h.data.At(0)
}

func (h *Heap) Push(v interface{}) {
	if !h.safe {
		h.state = w
		h.data.PushBack(v)
		algorithm.PushHeap(h.data.Begin(), h.data.End(), h.cmp)
		h.state = n
		return
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	h.state = w
	h.data.PushBack(v)
	algorithm.PushHeap(h.data.Begin(), h.data.End(), h.cmp)
	h.state = n
}

func (h *Heap) Emplace(v interface{}) {
	h.data.PushBack(v)
	algorithm.PushHeap(h.data.Begin(), h.data.End(), h.cmp)
}

func (h *Heap) Pop() {
	if !h.safe {
		h.state = w
		algorithm.PopHeap(h.data.Begin(), h.data.End(), h.cmp)
		h.data.PopBack()
		h.state = n
		return
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	h.state = w
	algorithm.PopHeap(h.data.Begin(), h.data.End(), h.cmp)
	h.data.PopBack()
	h.state = n
}

func (h *Heap) Swap(t *Heap) {
	if !h.safe {
		tmp := h.data
		h.data = t.data
		t.data = tmp
	}

	h.locker.Lock()
	defer h.locker.Unlock()

	h.state = w

	tmp := h.data
	h.data = t.data
	t.data = tmp

	h.state = n
}

func (h *Heap) String() string {
	return algorithm.ShowHeap(h.data.Begin(), h.data.End())
}
