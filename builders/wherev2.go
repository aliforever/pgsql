package builders

import (
	"fmt"
	"strings"
	"sync"
)

type queryBuilder struct {
	builder *strings.Builder

	placeHolderIndex      int
	placeHolderIndexMutex sync.Mutex

	values []interface{}
}

func newWhereV2() *queryBuilder {
	return &queryBuilder{builder: &strings.Builder{}, placeHolderIndex: 0}
}

func (w *queryBuilder) newPlaceHolder() int {
	w.placeHolderIndexMutex.Lock()
	defer w.placeHolderIndexMutex.Unlock()

	w.placeHolderIndex++

	return w.placeHolderIndex
}

func (w *queryBuilder) Where(fieldName, operand string, value interface{}) *whereClause {
	w.values = append(w.values, value)
	return newWhereClause(w, fieldName, operand, value)
}

func (w *queryBuilder) GroupBy(column string) *queryBuilder {
	w.builder.WriteString(fmt.Sprintf(" GROUP BY %s", column))

	return w
}

// OrderBy TODO: ASC/DESC per field
func (w *queryBuilder) OrderBy(column string) *queryBuilder {
	w.builder.WriteString(fmt.Sprintf(" GROUP BY %s", column))

	return w
}
