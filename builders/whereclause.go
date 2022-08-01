package builders

import "fmt"

type whereClause struct {
	builder *queryBuilder

	field   string
	operand string
	val     interface{}
}

func newWhereClause(builder *queryBuilder, field, operand string, val interface{}) *whereClause {
	_, _ = builder.builder.WriteString(fmt.Sprintf("%s %s $%d", field, operand, builder.newPlaceHolder()))
	return &whereClause{builder: builder, field: field, operand: operand, val: val}
}

func (w *whereClause) And(field, operand string, val interface{}) *whereClause {
	_, _ = w.builder.builder.WriteString(fmt.Sprintf(" AND %s %s $%d", field, operand, w.builder.newPlaceHolder()))
	w.builder.values = append(w.builder.values, val)
	return w
}

func (w *whereClause) Or(field, operand string, val interface{}) *whereClause {
	_, _ = w.builder.builder.WriteString(fmt.Sprintf(" OR %s %s $%d", field, operand, w.builder.newPlaceHolder()))
	w.builder.values = append(w.builder.values, val)
	return w
}

func (w *whereClause) OrGroup(field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.group("OR", field, operand, val, fn)
	return w
}

func (w *whereClause) AndGroup(field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.group("AND", field, operand, val, fn)
	return w
}

func (w *whereClause) group(keyword, field, operand string, val interface{}, fn func(qb *whereClause)) *whereClause {
	w.builder.builder.WriteString(fmt.Sprintf(" %s (", keyword))

	wc := newWhereClause(w.builder, field, operand, val)

	fn(wc)

	wc.builder.builder.WriteString(")")

	w.builder.values = append(w.builder.values, val)

	return w
}
