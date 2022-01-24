// @program:     tiny-stl
// @file:        pair.go
// @author:      edte
// @create:      2022-01-15 17:40
// @description:
package pair

type Pair struct {
	first  interface{}
	second interface{}
}

func New(first, second interface{}) *Pair {
	return &Pair{
		first:  first,
		second: second,
	}
}

func (p *Pair) First() interface{} {
	return p.first
}

func (p *Pair) Second() interface{} {
	return p.second
}
