// @program:     tiny-stl
// @file:        main.go
// @author:      edte
// @create:      2021-12-24 19:24
// @description:
package main

import (
	"stl/algorithm"
	"stl/containers/vector"
)

func main() {
	v := vector.New()
	v.PushBack(23)
	v.PushBack(3)
	algorithm.ForEach(v.Begin(), v.End())
}
