// @program:     tiny-stl
// @file:        sort_test.go.go
// @author:      edte
// @create:      2021-12-30 17:19
// @description:
package algorithm

import (
	"fmt"
	"testing"

	"stl/comparator"
	"stl/containers/vector"
)

func TestSort(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	Sort(v.Begin(), v.End())
	fmt.Println(v)
}

func TestSortCmp(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{9, 3, 6, 3, 1, 5, 2, 99, -1}))
	Sort(v.Begin(), v.End(), comparator.NewGreater())
	fmt.Println(v)
}

func TestSortUser(t *testing.T) {
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

	Sort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		if a.(user).age != b.(user).age {
			return a.(user).age > b.(user).age
		}
		return a.(user).name > b.(user).name
	}))

	fmt.Println(v)

	Sort(v.Begin(), v.End(), comparator.New(func(a interface{}, b interface{}) bool {
		if a.(user).age != b.(user).age {
			return a.(user).age > b.(user).age
		}
		return a.(user).name < b.(user).name
	}))

	fmt.Println(v)
}
