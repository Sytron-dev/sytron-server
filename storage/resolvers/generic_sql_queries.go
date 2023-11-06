package resolvers

import "sytron-server/storage/conn"

type SqlResolver[T any] struct {
	tableName         string
	model             T
	atomicRows        string // id or label-value pairs
	minimalRows       string // preferred for multi-selection
	comprehensizeRows string // preferred for single-selection
}

var db = conn.Supa.DB

// Generic queries

func (r SqlResolver[T]) FindMany() (data []T, err error) {
	err = db.From(r.tableName).Select(r.atomicRows).Execute(&data)
	return
}

func (r SqlResolver[T]) FindOne() (data T, err error) {
	err = db.From(r.tableName).Select(r.comprehensizeRows).Execute(&data)
	return
}
