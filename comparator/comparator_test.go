// @program:     tiny-stl
// @file:        comparator_test.go.go
// @author:      edte
// @create:      2022-01-14 01:11
// @description:
package comparator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefault_Operator(t *testing.T) {
	type fields struct {
		cmp func(a interface{}, b interface{}) bool
	}
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultCmp{}
			if got := d.Operator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Operator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreater_Operator(t *testing.T) {
	fmt.Println(Reserve(NewGreater()).Operator(1, 2))
	fmt.Println((NewGreater()).Operator(1, 2))
}

func TestLess_Operator(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Less{}
			if got := g.Operator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Operator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		cmp func(a interface{}, b interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want *DefaultCmp
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.cmp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGreater(t *testing.T) {
	type args struct {
		equal []bool
	}
	tests := []struct {
		name string
		args args
		want *Greater
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGreater(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGreater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLess(t *testing.T) {
	tests := []struct {
		name string
		want *Less
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLess(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equalGreater_Operator(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EGreater{}
			if got := e.Operator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Operator() = %v, want %v", got, tt.want)
			}
		})
	}
}
