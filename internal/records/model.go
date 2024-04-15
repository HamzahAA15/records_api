package records

import (
	"time"

	"github.com/lib/pq"
)

type RequestPayload struct {
	StartDate string `query:"startDate"`
	EndDate   string `query:"endDate"`
	MinCount  int64  `query:"minCount"`
	MaxCount  int64  `query:"maxCount"`
}

type ResponsePayload struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Records []Record `json:"records,omitempty"`
}

type Record struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	TotalMarks int64     `json:"totalMarks"`
}

type RecordModel struct {
	ID        int64         `json:"id"`
	Name      string        `json:"name"`
	Marks     pq.Int64Array `json:"marks"`
	CreatedAt time.Time     `json:"createdAt"`
}
