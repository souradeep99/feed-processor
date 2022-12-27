package integrators

import (
	"feed-processor/feedback"
	"time"
)

// Integrator represents an integrator for a feedback source.
type Integrator interface {
	FetchData(startTime time.Time, endTime time.Time) (interface{}, error)
	ProcessData(interface{}) ([]*feedback.Feedback, error)
	StoreData([]*feedback.Feedback, *database.DB) error
}
