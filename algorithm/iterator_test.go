// @program:     tiny-stl
// @file:        iterator_test.go.go
// @author:      edte
// @create:      2022-01-16 22:11
// @description:
package algorithm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"

	"stl/containers/vector"
)

func TestAdvance(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	iter := v.Begin()
	fmt.Println(iter)
	Advance(iter, 3)
	fmt.Println(iter)
}

func TestDistance(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	assert.Equal(t, Distance(v.Begin(), v.End()), 5)
}

func TestNext(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	it := v.Begin()
	i := Next(v.Begin())
	i.Next()
	i.Next()
	assert.Equal(t, "1", fmt.Sprint(it))
}

func TestNextN(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	//it := v.Begin()
	//assert.Equal(t, "23", fmt.Sprint(NextN(it, 3)))
	//assert.Equal(t, "23", fmt.Sprint(NextN(it, 3)))
	//assert.Equal(t, "23", fmt.Sprint(NextN(it, 3)))
	//assert.Equal(t, "1", fmt.Sprint(it))

	it := v.Begin()
	i := NextN(v.Begin(), 3)
	i.Next()
	i.Next()
	assert.Equal(t, "1", fmt.Sprint(it))
}

func TestPre(t *testing.T) {

}

func TestPreN(t *testing.T) {
	v := vector.New(vector.WithData([]interface{}{1, 2, 34, 23, 8}))
	it := v.End().Pre()
	assert.Equal(t, "2", fmt.Sprint(PreN(it, 3)))
	assert.Equal(t, "8", fmt.Sprint(it))
}
