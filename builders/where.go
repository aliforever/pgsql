package builders

import (
	"fmt"
)

type where struct {
	key      string
	operand  string
	value    interface{}
	previous *where
	isOr     bool
	group    *where
	ignore   bool

	placeHolderIndex int
}

func newWhere(key string, operand string, value interface{}, isOr bool, placeHolderIndex int) *where {
	return &where{
		key:     key,
		operand: operand,
		value:   value,
		isOr:    isOr,

		placeHolderIndex: placeHolderIndex,
	}
}

func newQueryBuilder(fn func(builder *queryBuilder)) *queryBuilder {
	qb := newWhereV2()
	fn(qb)
	return qb
}

func (w *where) where(key string, operand string, value interface{}, isOr, isGroup bool) *where {
	nw := &where{
		key:              key,
		operand:          operand,
		value:            value,
		isOr:             isOr,
		previous:         w,
		placeHolderIndex: w.placeHolderIndex + 1,
	}

	return nw
}

func (w *where) WhereGroup(closure func(*where) *where) *where {
	nw := newWhere("", "", nil, false, w.placeHolderIndex)
	w.group = closure(nw)

	return w
}

func (w *where) Where(key string, operand string, value interface{}) *where {

	return w.where(key, operand, value, false, false)
}

func (w *where) OrWhere(key string, operand string, value interface{}) *where {
	return w.where(key, operand, value, true, false)
}

func (w *where) keyOperandValue() string {
	if w.key == "" || w.value == nil || w.operand == "" {
		return ""
	}

	return fmt.Sprintf("%s %s $%d", w.key, w.operand, w.placeHolderIndex)
}

func (w *where) operator() string {
	fmt.Println(w)
	if w.isOr {
		return "OR"
	}

	return "AND"
}

func (w *where) build() string {
	return traverse(w)
}
