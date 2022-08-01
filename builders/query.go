package builders

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

type scanner interface {
	Scan(dest ...any) error
}

type Q interface {
	Where(fieldName, operand string, value interface{}) *whereClause
}

type query[T tbl] struct {
	db        *sql.DB
	tableName string

	builder *strings.Builder

	placeHolderIndex      int
	placeHolderIndexMutex sync.Mutex

	values []interface{}
}

func newQuery[T tbl](db *sql.DB, fn func(builder Q)) *query[T] {
	var t T
	q := &query[T]{builder: &strings.Builder{}, placeHolderIndex: 0, db: db, tableName: t.TableName()}

	if fn != nil {
		fn(q)
	}

	return q
}

func (q *query[T]) addValue(val interface{}) {
	q.values = append(q.values, val)
}

func (q *query[T]) Where(fieldName, operand string, value interface{}) *whereClause {
	q.values = append(q.values, value)
	return newWhereClause(q.builder, q.newPlaceHolder, q.addValue, fieldName, operand, value)
}

func (q *query[T]) FindOne() (data *T, err error) {
	fields, err := q.fields()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(fields, ","), q.tableName, q.builder.String())

	r := q.db.QueryRow(query, q.values...)
	if r.Err() != nil {
		return nil, r.Err()
	}

	return q.decodeResult(fields, r)
}

func (q *query[T]) Find() (data []T, err error) {
	fields, err := q.fields()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ","), q.tableName)

	if q.builder != nil && q.builder.String() != "" {
		query += fmt.Sprintf(" WHERE %s", q.builder.String())
	}

	cur, err := q.db.Query(query, q.values...)
	if err != nil {
		return nil, err
	}
	defer cur.Close()

	for cur.Next() {
		if d, err := q.decodeResult(fields, cur); err != nil {
			continue
		} else {
			data = append(data, *d)
		}
	}

	return data, cur.Err()
}

func (q *query[T]) decodeResult(fields []string, sc scanner) (*T, error) {
	values := make([]interface{}, len(fields))

	for i := range fields {
		values[i] = &values[i]
	}

	err := sc.Scan(values...)
	if err != nil {
		return nil, err
	}

	d := map[string]interface{}{}
	for i, value := range values {
		d[fields[i]] = value
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	var data *T

	err = json.Unmarshal(j, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (q *query[T]) fields() (fields []string, err error) {
	var t T

	m, err := dataToMap(t)
	if err != nil {
		return nil, err
	}

	for key := range m {
		fields = append(fields, key)
	}

	return fields, nil
}

func (q *query[T]) newPlaceHolder() int {
	q.placeHolderIndexMutex.Lock()
	defer q.placeHolderIndexMutex.Unlock()

	q.placeHolderIndex++

	return q.placeHolderIndex
}
