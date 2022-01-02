// @program:     tiny-stl
// @file:        boom_test.go.go
// @author:      edte
// @create:      2022-01-02 16:32
// @description:
package boom

import (
	"fmt"
	"math/rand"
	"stl/algorithm/hash"
	"testing"
	"time"
)

func TestBloomFilter_Clear(t *testing.T) {
	b := New(500000)
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
	b.Clear()
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
}

func TestBloomFilter_Contains(t *testing.T) {
	b := New(500000)
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
}

func TestBloomFilter_Empty(t *testing.T) {
	b := New(500000)
	fmt.Println(b.Empty())
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
	fmt.Println(b.Empty())
}

func TestBloomFilter_Insert(t *testing.T) {
	b := New(500000)
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
}

func TestNew(t *testing.T) {
}

func TestWithHash(t *testing.T) {
	b := New(5000000, WithHash(hash.BKDRHash64))
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
}

func TestWithSeeds(t *testing.T) {
	b := New(5000000, WithHash(hash.BKDRHash64), WithSeeds([]uint64{31, 131, 1313, 13131, 131313, 1313131, 13131313, 131313131, 1313131313, 13131313131, 131313131313, 1313131313131, 13131313131313, 131313131313131}))
	b.Insert([]byte("hello"))
	fmt.Println(b.Contains([]byte("hello")))
	fmt.Println(b.Contains([]byte("word")))
}

func GetRandomString(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func BenchmarkBoom(b *testing.B) {
	n := New(1000000000)

	rand.Seed(time.Now().UnixNano())

	var words []string
	for i := 0; i < 1000000; i++ {
		s := GetRandomString(32)
		words = append(words, s)
		n.Insert([]byte(s))
	}

	var rids []string
	for i := 0; i < 10000; i++ {
		rids = append(rids, GetRandomString(32))
	}

	for i := 0; i < 1000000; i++ {
		if i%10000 == 0 {
			fmt.Println(n.Contains([]byte(rids[i/10000])))
		}
		if i%333 == 0 {
			n.Insert([]byte(words[i]))
		}
	}
}
