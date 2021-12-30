// @program:     tiny-stl
// @file:        best_test.go.go
// @author:      edte
// @create:      2021-12-30 16:22
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestMinElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	fmt.Println(v)

	e := MinElement(v.Begin(), v.End())
	fmt.Println(e.Value())
}

func TestBestElement(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 6, 3, 7, 2, 98, -1, 63}))
	fmt.Println(v)
	j := BestElement(v.Begin(), v.End(), comparator.NewGreater())
	fmt.Println(j.Value())
}

func TestBestElementUser(t *testing.T) {
	type user struct {
		age  int
		name string
	}

	v := vector.New(vector.WithData([]interface{}{
		user{
			age:  17,
			name: "a",
		},
		user{
			age:  20,
			name: "b",
		},
		user{
			age:  20,
			name: "a",
		},
		user{
			age:  16,
			name: "c",
		},
		user{
			age:  17,
			name: "c",
		},
		user{
			age:  16,
			name: "e",
		},
	}))

	e := MaxElement(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		if a.(user).age != b.(user).age {
			return a.(user).age > b.(user).age
		}
		return a.(user).name < b.(user).name
	}))

	fmt.Println(e.Value())
}
