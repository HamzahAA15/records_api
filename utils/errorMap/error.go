package errorMap

import "errors"

var (
	RecordsNotFound error = errors.New("no records found")
)
