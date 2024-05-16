package logger

import "time"

type Event struct {
	Time    time.Time
	ID      int
	Client    string
	TableNum int
}
type Table struct {
	Client string
	Start time.Time
	Duration time.Duration
	Income int
}

type Club struct {
	Tables     []Table
	Clients    map[string]bool
	Queue      []string
	StartTime  time.Time
	EndTime    time.Time
	HourlyRate int
	Events     []Event
	Result []Event
}