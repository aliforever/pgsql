package builders

import (
	"fmt"
	"testing"
)

func Test_newWhere(t *testing.T) {
	w := newWhere("first_name", "=", "Ali", false, 1).
		Where("last_name", "=", "Dehkharghani").
		OrWhere("age", "<", 20).
		WhereGroup(func(w *where) *where {
			return w.Where("city", "=", "Tabriz").Where("state", "=", "EA")
		}).
		build()
	fmt.Println(w)
}

// func Test_newWhere2(t *testing.T) {
// 	w := newQueryBuilder(func(builder *query) {
// 		builder.Where("first_name", "=", "Ali").
// 			And("last_name", "=", "Dehkharghani").
// 			OrGroup("age", "<", 20, func(qb *whereClause) {
// 				qb.Or("age", ">", 40)
// 			}).
// 			OrGroup("city", "=", "Tabriz", func(qb *whereClause) {
// 				qb.And("state", "=", "EA")
// 			})
// 	}).GroupBy("city")
// 	fmt.Println(w.builder.String())
// }
