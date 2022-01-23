// @program:     tiny-stl
// @file:        priority_queue_test.go.go
// @author:      edte
// @create:      2022-01-16 00:14
// @description:
package priorityqueue

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"stl/algorithm"
	"stl/comparator"
	"stl/containers/vector"
	"stl/locker"
)

func TestHeap_Emplace(t *testing.T) {
	h := New()

	for i := 0; i < 100; i++ {
		if i%3 == 0 || i%7 == 0 {
			func() {
				h.Push(rand.Int() % 100)
			}()
		}
	}

	fmt.Println(h)

	fmt.Println(algorithm.IsHeap(h.container.Begin(), h.container.End()))
}

func TestHeap_Empty(t *testing.T) {
}

func TestHeap_Pop(t *testing.T) {
	h := New()
	h.Push(1)
	h.Push(8)
	h.Push(3)
	h.Push(2)
	h.Push(-1)
	fmt.Println(h)
	h.Pop()
	fmt.Println(h)
}

func TestHeap_Push(t *testing.T) {

}

func TestHeap_Size(t *testing.T) {
	type fields struct {
		data   *vector.Vector
		cmp    comparator.Comparator
		locker locker.Locker
		safe   bool
		state  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PriorityQueue{
				container: tt.fields.data,
				cmp:       tt.fields.cmp,
				locker:    tt.fields.locker,
				safe:      tt.fields.safe,
				state:     tt.fields.state,
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeap_String(t *testing.T) {
	h := New()
	h.Push(1)
	h.Push(8)
	h.Push(3)
	h.Push(2)
	h.Push(-1)
	fmt.Println(h)
}

func TestHeap_Swap(t *testing.T) {

}

func TestHeap_Top(t *testing.T) {
	type fields struct {
		data   *vector.Vector
		cmp    comparator.Comparator
		locker locker.Locker
		safe   bool
		state  int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &PriorityQueue{
				container: tt.fields.data,
				cmp:       tt.fields.cmp,
				locker:    tt.fields.locker,
				safe:      tt.fields.safe,
				state:     tt.fields.state,
			}
			if got := h.Top(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Top() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *PriorityQueue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCap(t *testing.T) {
	type args struct {
		c int
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCap(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithCmp(t *testing.T) {
	type args struct {
		cmp func(a interface{}, b interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCmp(tt.args.cmp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithData(t *testing.T) {
	type args struct {
		data []interface{}
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithGoroutineSafe(t *testing.T) {
	tests := []struct {
		name string
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithGoroutineSafe(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithGoroutineSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithLocker(t *testing.T) {
	type args struct {
		l locker.Locker
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithLocker(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLocker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithOperator(t *testing.T) {
	type args struct {
		c comparator.Comparator
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithOperator(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithOperator() = %v, want %v", got, tt.want)
			}
		})
	}
}
