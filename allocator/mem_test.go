// @program:     tiny-stl
// @file:        mem_test.go.go
// @author:      edte
// @create:      2022-01-24 00:48
// @description:
package allocator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_malloc(t *testing.T) {
	str := "hello world"
	str1 := "what   fuck"

	p := malloc(len(str))

	bytes := *(*[]byte)(p)

	for i := range bytes {
		bytes[i] = str[i]
	}

	assert.Equal(t, "hello world", string(bytes))

	for i := range bytes {
		bytes[i] = str1[i]
	}

	assert.Equal(t, "what   fuck", string(bytes))
}

func Test_calloc(t *testing.T) {
	p := calloc(64, 1)

	b := *(*int64)(p)

	b = 1234

	assert.Equal(t, int64(1234), b)
}

func Test_free(t *testing.T) {
	str := "hello world"

	p := malloc(len(str))

	bytes := *(*[]byte)(p)

	for i := range bytes {
		bytes[i] = str[i]
	}

	assert.Equal(t, "hello world", string(bytes))

	free(p)

	assert.NotEqual(t, "hello world", string(bytes))
}

func Test_realloc(t *testing.T) {
	p := malloc(10)

	assert.Equal(t, 10, len(*(*[]byte)(p)))

	ppp := realloc(p, 20)

	assert.Equal(t, 20, len(*(*[]byte)(ppp)))

	assert.False(t, ppp == p)
}
