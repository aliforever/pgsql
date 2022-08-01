package builders

import (
	"encoding/json"
	"fmt"
	"strings"
)

func dataToMap(data any) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(j, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func traverse(w *where) string {
	var ands []string

	str := w.keyOperandValue()
	if str != "" {
		ands = append(ands, str)
	}

	if w.previous != nil {
		q := traverse(w.previous)
		if q != "" {
			ands = append(ands, q)
		}
	}

	if w.group != nil {
		q := traverse(w.group)
		if q != "" {
			ands = append(ands, fmt.Sprintf("(%s)", q))
		}
	}

	reverse(ands)

	return strings.Join(ands, fmt.Sprintf(" %s ", w.operator()))
}
