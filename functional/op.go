// @program:     tiny-stl
// @file:        op.go.go
// @author:      edte
// @create:      2022-01-16 21:34
// @description:
package functional

import (
	rand "crypto/rand"
	"fmt"
	"math/big"
	"stl/iterator"
)

// Op 一元函数
type Op func(val interface{}) interface{}

var (
	// Print 打印
	Print Op = func(val interface{}) interface{} {
		fmt.Print(val)
		fmt.Print(" ")
		return val
	}

	// Negate 取相反数
	Negate Op = func(val interface{}) interface{} {
		x := val.(int)
		return -x
	}

	// Increase 自增
	Increase Op = func(val interface{}) interface{} {
		i := val.(int)
		i++
		return i
	}
)

// BinaryOp 二元函数
type BinaryOp func(x, y int) int

var (
	// BitAnd 按位与
	BitAnd BinaryOp = func(x int, y int) int {
		return x & y
	}

	// BitOr 按位或
	BitOr BinaryOp = func(x int, y int) int {
		return x | y
	}

	// BitXor 按位异或
	BitXor BinaryOp = func(x int, y int) int {
		return x ^ y
	}

	// Divides 除法
	Divides BinaryOp = func(x int, y int) int {
		return x / y
	}

	// Minus 减法
	Minus BinaryOp = func(x int, y int) int {
		return x - y
	}

	// Modulus 取余
	Modulus BinaryOp = func(x int, y int) int {
		return x % y
	}

	// Multiplies 乘法
	Multiplies BinaryOp = func(x int, y int) int {
		return x * y
	}

	// Add 加法
	Add BinaryOp = func(x int, y int) int {
		return x + y
	}
)

// IterOp 迭代器函数
type IterOp func(i iterator.ForwardIterator)

// Pred 一元谓词函数
type Pred func(i interface{}) bool

var (
	// IsOdd 奇数
	IsOdd Pred = func(i interface{}) bool {
		return i.(int)%2 != 0
	}

	// IsEven 偶数
	IsEven Pred = func(i interface{}) bool {
		return i.(int)%2 == 0
	}

	// IsPositive 正数
	IsPositive Pred = func(i interface{}) bool {
		return i.(int) > 0
	}

	// IsNegative 正数
	IsNegative Pred = func(i interface{}) bool {
		return i.(int) < 0
	}

	// LogicalNot 取反
	LogicalNot Pred = func(i interface{}) bool {
		x, ok := i.(bool)
		if !ok {
			return false
		}
		return !x
	}
)

// BinaryPred 二元谓词函数
type BinaryPred func(x, y interface{}) bool

var (
	// EqualTo 相等
	EqualTo BinaryPred = func(x, y interface{}) bool {
		return x == y
	}

	// NotEqualTo 不相等
	NotEqualTo BinaryPred = func(x, y interface{}) bool {
		return x != y
	}

	// Greater 大于
	Greater BinaryPred = func(x, y interface{}) bool {
		return x.(int) > y.(int)
	}

	// GreaterEqual 大于等于
	GreaterEqual BinaryPred = func(x, y interface{}) bool {
		return x.(int) >= y.(int)
	}

	// Less 小于
	Less BinaryPred = func(x, y interface{}) bool {
		return x.(int) < y.(int)
	}

	// LessEqual 小于等于
	LessEqual BinaryPred = func(x, y interface{}) bool {
		return x.(int) <= y.(int)
	}

	// LogicalAnd 逻辑与
	LogicalAnd BinaryPred = func(x, y interface{}) bool {
		return x.(bool) && y.(bool)
	}

	// LogicalOr 逻辑或
	LogicalOr BinaryPred = func(x, y interface{}) bool {
		return x.(bool) || y.(bool)
	}
)

func DefaultBinaryPared(preds ...BinaryPred) BinaryPred {
	if len(preds) != 0 {
		return preds[0]
	}
	return EqualTo
}

// Not1 return negation of unary function object
func Not1(pred Pred) Pred {
	return func(i interface{}) bool {
		return !pred(i)
	}
}

// Not2 return negation of binary function object
func Not2(pred BinaryPred) BinaryPred {
	return func(x, y interface{}) bool {
		return !pred(x, y)
	}
}

// IGenerator 数据产生器
type IGenerator interface {
	Gen() interface{}
}

type gen struct {
	size int
	Generator
}

func newGen(size int) *gen {
	return &gen{size: size}
}

func (g *gen) Gen() interface{} {
	return g.Generator()
}

type Generator func() interface{}

var (
	RandomNum Generator = func() interface{} {
		n, _ := rand.Int(rand.Reader, big.NewInt(100))
		return n
	}
)
