package builders

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

type query[T tbl] struct {
	db        *sql.DB
	tableName string

	singles map[string]interface{} // WHERE (name = "ALI") // WHERE (NAME = "ALI" OR NAME = "Anıl")
}

func newQuery[T tbl](db *sql.DB, tbl string) *query[T] {
	return &query[T]{
		db:        db,
		tableName: tbl,
		singles:   map[string]interface{}{},
	}
}

func (q *query[T]) Where(key string, operand string, value interface{}) *query[T] {
	q.singles[key] = value
	return q
}

func (q *query[T]) FindOne() (data *T, err error) {
	fields, err := q.fields()
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ","), q.tableName)

	values := []interface{}{}

	i := 1
	for key, val := range q.singles {
		query += fmt.Sprintf(" WHERE %s = $%d", key, i)
		values = append(values, val)
		i++
	}

	r := q.db.QueryRow(query, values...)
	if r.Err() != nil {
		return nil, r.Err()
	}

	return q.decodeResult(fields, r)
}

func (q *query[T]) decodeResult(fields []string, row *sql.Row) (*T, error) {
	values := make([]interface{}, len(fields))

	for i := range fields {
		values[i] = &values[i]
	}

	err := row.Scan(values...)
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
