// @program:     tiny-stl
// @file:        queue_test.go.go
// @author:      edte
// @create:      2022-01-15 18:58
// @description:
package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	q.Pop()
	fmt.Println(q.Front())
	q.Pop()

	fmt.Println(q.Empty())
}
