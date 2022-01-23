// @program:     tiny-stl
// @file:        stack.go
// @author:      edte
// @create:      2022-01-15 17:37
// @description:
package stack

import (
	"stl/allocator"
	"stl/containers"
	"stl/containers/deque"
	"stl/locker"
)

type Option func(s *Stack)

type Stack struct {
	container containers.SequentialContainer
	locker    locker.Locker
}

func New(opts ...Option) *Stack {
	s := &Stack{
		container: deque.New(),
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithContainer(c containers.SequentialContainer) Option {
	return func(s *Stack) {
		s.container = c
	}
}

func (s *Stack) Swap(t *containers.Container) {
	s.container.Swap(t)
}

func (s *Stack) Clone() containers.Container {
	return s.container.Clone()
}

func (s *Stack) Clear() {
	s.container.Clear()
}

func (s *Stack) MaxSize() int {
	return s.container.MaxSize()
}

func (s *Stack) Capacity() int {
	return s.container.Capacity()
}

func (s *Stack) GetAllocator() allocator.Allocator {
	return s.container.GetAllocator()
}

func (s *Stack) Construct() {
	s.container.Construct()
}

func (s *Stack) Destroy() {
	s.container.Destroy()
}

func (s *Stack) String() string {
	return s.container.String()
}

func (s *Stack) Empty() bool {
	return s.container.Empty()
}

func (s *Stack) Size() int {
	return s.container.Size()
}

func (s *Stack) Top() interface{} {
	return s.container.Back()
}

func (s *Stack) Push(x interface{}) {
	s.container.PushBack(x)
}

func (s *Stack) Emplace(x interface{}) {
	s.container.EmplaceBack(x)
}

func (s *Stack) Pop() {
	s.container.PopBack()
}
