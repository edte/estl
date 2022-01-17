// @program:     tiny-stl
// @file:        op.go.go
// @author:      edte
// @create:      2022-01-16 21:34
// @description:
package functional

type Op func(val interface{})

type BinaryOp func(x, y int) int

// Pred UnaryPredicate
type Pred func(i interface{}) bool

type BinaryPred func(x, y interface{}) bool
