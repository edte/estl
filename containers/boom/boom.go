// @program:     tiny-stl
// @file:        boom.go
// @author:      edte
// @create:      2021-12-25 00:32
// @description:
package boom

import (
	"stl/algorithm/hash"
	"stl/containers/bitset"
)

type boom interface {
	Empty() bool
	Size() int
	Insert(key []byte)
	Contains(key []byte) bool
	Clear()
}

const (
	defaultHashNum = 14
)

type Option func(filter *BloomFilter)

type BloomFilter struct {
	cap     uint      // 构造的大小
	hashNum uint64    // hash 函数个数，默认为 14
	size    int       // 插入的元素个数
	seeds   []uint64  // 哈希种子序列
	h       hash.Hash // 哈希函数，默认为 Murmur3
	data    *bitset.BitSet
}

// New create a Filter, store is the backed redis, key is the key for the bloom filter,
// bits is how many bits will be used, maps is how many hashes for each addition.
// best practices:
// elements - means how many actual elements
// when maps = 14, formula: 0.7*(bits/maps), bits = 20*elements, the error rate is 0.000067 < 1e-4
// for detailed error rate table, see http://pages.cs.wisc.edu/~cao/papers/summary-cache/node8.html
func New(cap uint, opts ...Option) *BloomFilter {
	b := &BloomFilter{
		cap:     cap,
		hashNum: defaultHashNum, // 默认为 14
		size:    0,
		// 默认为素数序列
		seeds: []uint64{98317, 196613, 393241, 786433, 1572869, 3145739, 6291469, 12582917, 12582917, 50331653, 100663319, 201326611, 402653189, 805306457},
		// 默认为, Murmur3,原因见 https://softwareengineering.stackexchange.com/questions/49550/which-hashing-algorithm-is-best-for-uniqueness-and-speed
		h:    hash.MurmurHash64,
		data: bitset.New(bitset.WithSize(cap)),
	}

	for _, opt := range opts {
		opt(b)
	}

	return b
}

func WithSeeds(seeds []uint64) Option {
	return func(filter *BloomFilter) {
		filter.seeds = seeds
	}
}

func WithHash(h hash.Hash) Option {
	return func(filter *BloomFilter) {
		filter.h = h
	}
}

func (b *BloomFilter) Empty() bool {
	return b.size == 0
}

func (b *BloomFilter) Insert(key []byte) {
	for _, seed := range b.seeds {
		h := b.h(key, seed)
		b.data.Set(uint(h) % b.cap)
	}
	b.size++
}

func (b *BloomFilter) Contains(key []byte) bool {
	for _, seed := range b.seeds {
		h := b.h(key, seed)
		if !b.data.Test(uint(h) % b.cap) {
			return false
		}
	}
	return true
}

func (b *BloomFilter) Clear() {
	b.data.Clear()
	b.size = 0
}

func (b *BloomFilter) Size() int {
	return b.size
}
