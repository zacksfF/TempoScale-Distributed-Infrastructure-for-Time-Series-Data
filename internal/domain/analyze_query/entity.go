package analyzequery

import "time"

const (
	ObservationAnalyzerRequestInsertType = 1
	ObservationAnalyzerRequestDeleteType = 2
)

type Observation struct {
	EntityID  uint64    `json:"entity_id"` // ex: xxxx-xxx-xxxx-xxxx
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     float64   `json:"value"`     // ex: 123.456
}

type ObservationAnalyzerRequest struct {
	EntityID    uint64         `json:"entity_id"`   // ex: 1234
	UUID        string         `json:"uuid"`        // ex: xxxx-xxx-xxxx-xxxx
	Timestamp   time.Time      `json:"timestamp"`   // ex: 2022-10-09 15:57:52 -0400 -0400
	Type        int8           `json:"type"`        // ex: 1
	Observation *Observation `json:"observation"` // ex: { ... }
}

type ObservationAnalyzerRequestFilter struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}
