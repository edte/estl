// @program:     tiny-stl
// @file:        queue.go
// @author:      edte
// @create:      2021-12-25 00:33
// @description:
package queue

// FIFO queue
// queues are a type of container adaptor, specifically
// designed to operate in a FIFO context (first-in first-out),
// where elements are inserted into one end of the container
// and extracted from the other.
//
// queues are implemented as containers adaptors, which are
// classes that use an encapsulated object of a specific
// container class as its underlying container, providing
// a specific set of member functions to access its elements.
// Elements are pushed into the "back" of the specific
// container and popped from its "front".
//
// The underlying container may be one of the standard
// container class template or some other specifically
// designed container class. This underlying container shall
// support at least the following operations:
// empty
// size
// front
// back
// push_back
// pop_front
//
// The standard container classes deque and list fulfill
// these requirements. By default, if no container class is
// specified for a particular queue class instantiation,
// the standard container deque is used.

type queue interface {
	Empty() bool
	Size() int
	Front() interface{}
	Back() interface{}
	Push(v interface{})
	Emplace(v interface{})
	Pop()
	Swap(t *Queue)
}

type Option func(*Queue)

type Queue struct {
	data []interface{}
}

func New(opts ...Option) *Queue {
	q := &Queue{}

	for _, opt := range opts {
		opt(q)
	}

	return q
}

// Empty test whether container is empty
func (q *Queue) Empty() bool {
	return q.Size() == 0
}

// Size return size
func (q *Queue) Size() int {
	return len(q.data)
}

// Front access next element
func (q *Queue) Front() interface{} {
	return q.data[0]
}

// Back access last element
func (q *Queue) Back() interface{} {
	return q.data[q.Size()-1]
}

// Push insert element
func (q *Queue) Push(v interface{}) {
	q.data = append(q.data, v)
}

// Emplace construct and insert element
func (q *Queue) Emplace(v interface{}) {
	q.Push(v)
}

// Pop remove next element
func (q *Queue) Pop() {
	q.data = q.data[1:]
}

// Swap contents
func (q *Queue) Swap(t *Queue) {
	tmp := t.data
	t.data = q.data
	q.data = tmp
}
