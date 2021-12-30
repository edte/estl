// @program:     tiny-stl
// @file:        comparator.go
// @author:      edte
// @create:      2021-12-30 10:03
// @description:
package comparator

type Comparator interface {
	Operator(a interface{}, b interface{}) bool
}

type Default struct {
	cmp func(a interface{}, b interface{}) bool
}

func New(cmp func(a interface{}, b interface{}) bool) *Default {
	return &Default{cmp: cmp}
}

func (d *Default) Operator(a interface{}, b interface{}) bool {
	return d.cmp(a, b)
}

type Greater struct {
}

func NewGreater() *Greater {
	return &Greater{}
}

func (l *Greater) Operator(a interface{}, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) > b.(int)
	case int64:
		return a.(int64) > b.(int64)
	case int32:
		return a.(int32) > b.(int32)
	case int16:
		return a.(int16) > b.(int16)
	case int8:
		return a.(int8) > b.(int8)
	case float64:
		return a.(float64) > b.(float64)
	case float32:
		return a.(float32) > b.(float32)
	case bool:
		return a.(bool) != b.(bool)
	case string:
		return a.(string) != b.(string)
	default:
		return false
	}
}

type Less struct {
}

func NewLess() *Less {
	return &Less{}
}

func (g *Less) Operator(a interface{}, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) < b.(int)
	case int64:
		return a.(int64) < b.(int64)
	case int32:
		return a.(int32) < b.(int32)
	case int16:
		return a.(int16) < b.(int16)
	case int8:
		return a.(int8) < b.(int8)
	case float64:
		return a.(float64) < b.(float64)
	case float32:
		return a.(float32) < b.(float32)
	case bool:
		return a.(bool) != b.(bool)
	case string:
		return a.(string) != b.(string)
	default:
		return false
	}
}
