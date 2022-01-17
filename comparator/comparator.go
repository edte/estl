// @program:     tiny-stl
// @file:        comparator.go
// @author:      edte
// @create:      2021-12-30 10:03
// @description:
package comparator

type Operator func(a interface{}, b interface{}) bool

type Cmp func(a interface{}, b interface{}) int

type Comparator interface {
	Operator(a interface{}, b interface{}) bool
	Cmp(a interface{}, b interface{}) int
}

func Reserve(c Comparator) Comparator {

	return New(func(a interface{}, b interface{}) bool {
		return !c.Operator(a, b)
	}, WithCmp(func(a interface{}, b interface{}) int {
		return -c.Cmp(a, b)
	}))
}

type Option func(d *Default)

type Default struct {
	op  Operator
	cmp Cmp
}

func New(op Operator, opts ...Option) *Default {
	d := &Default{
		op: op,
	}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithOperator(op Operator) Option {
	return func(d *Default) {
		d.op = op
	}
}

func WithCmp(cmp Cmp) Option {
	return func(d *Default) {
		d.cmp = cmp
	}
}

func (d *Default) Operator(a interface{}, b interface{}) bool {
	return d.op(a, b)
}

func (d *Default) Cmp(a interface{}, b interface{}) int {
	return d.cmp(a, b)
}

type Greater struct {
}

func NewGreater() *Greater {
	return &Greater{}
}

// Operator 默认相等时不交换，这样许多排序默认就是稳定的
// 升序
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

func (l *Greater) Cmp(a interface{}, b interface{}) int {
	if a == b {
		return 0
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return -1
		}
	case int64:
		if a.(int64) > b.(int64) {
			return 1
		} else if a.(int64) == b.(int64) {
			return 0
		} else {
			return -1
		}
	case int32:
		if a.(int32) > b.(int32) {
			return 1
		} else if a.(int32) == b.(int32) {
			return 0
		} else {
			return -1
		}
	case int16:
		if a.(int16) > b.(int16) {
			return 1
		} else if a.(int16) == b.(int16) {
			return 0
		} else {
			return -1
		}
	case int8:
		if a.(int8) > b.(int8) {
			return 1
		} else if a.(int8) == b.(int8) {
			return 0
		} else {
			return -1
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) == b.(float64) {
			return 0
		} else {
			return -1
		}
	case float32:
		if a.(float32) > b.(float32) {
			return 1
		} else if a.(float32) == b.(float32) {
			return 0
		} else {
			return -1
		}
	case bool:
		if a.(bool) == b.(bool) {
			return 0
		} else {
			return -1
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) == b.(string) {
			return 0
		} else {
			return -1
		}
	default:
		return 0
	}
}

type EGreater struct {
}

func NewEGreater() *EGreater {
	return &EGreater{}
}

func (e *EGreater) Cmp(a interface{}, b interface{}) int {
	if a == b {
		return 0
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) == b.(int) {
			return 0
		} else {
			return -1
		}
	case int64:
		if a.(int64) > b.(int64) {
			return 1
		} else if a.(int64) == b.(int64) {
			return 0
		} else {
			return -1
		}
	case int32:
		if a.(int32) > b.(int32) {
			return 1
		} else if a.(int32) == b.(int32) {
			return 0
		} else {
			return -1
		}
	case int16:
		if a.(int16) > b.(int16) {
			return 1
		} else if a.(int16) == b.(int16) {
			return 0
		} else {
			return -1
		}
	case int8:
		if a.(int8) > b.(int8) {
			return 1
		} else if a.(int8) == b.(int8) {
			return 0
		} else {
			return -1
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) == b.(float64) {
			return 0
		} else {
			return -1
		}
	case float32:
		if a.(float32) > b.(float32) {
			return 1
		} else if a.(float32) == b.(float32) {
			return 0
		} else {
			return -1
		}
	case bool:
		if a.(bool) == b.(bool) {
			return 0
		} else {
			return -1
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) == b.(string) {
			return 0
		} else {
			return -1
		}
	default:
		return 0
	}
}

func (e *EGreater) Operator(a interface{}, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) >= b.(int)
	case int64:
		return a.(int64) >= b.(int64)
	case int32:
		return a.(int32) >= b.(int32)
	case int16:
		return a.(int16) >= b.(int16)
	case int8:
		return a.(int8) >= b.(int8)
	case float64:
		return a.(float64) >= b.(float64)
	case float32:
		return a.(float32) >= b.(float32)
	case bool:
		return a.(bool) != b.(bool)
	case string:
		return a.(string) != b.(string)
	default:
		return false
	}
}

// Less 降序
type Less struct {
}

func (g *Less) Cmp(a interface{}, b interface{}) int {
	panic("implement me")
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

type ELess struct {
}

func (g *ELess) Cmp(a interface{}, b interface{}) int {
	panic("implement me")
}

func NewELess() *ELess {
	return &ELess{}
}

func (g *ELess) Operator(a interface{}, b interface{}) bool {
	switch a.(type) {
	case int:
		return a.(int) <= b.(int)
	case int64:
		return a.(int64) <= b.(int64)
	case int32:
		return a.(int32) <= b.(int32)
	case int16:
		return a.(int16) <= b.(int16)
	case int8:
		return a.(int8) <= b.(int8)
	case float64:
		return a.(float64) <= b.(float64)
	case float32:
		return a.(float32) <= b.(float32)
	case bool:
		return a.(bool) != b.(bool)
	case string:
		return a.(string) != b.(string)
	default:
		return false
	}
}
