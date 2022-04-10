package model

import (
	"time"
)

type Task struct {
	Id        int       `json: "id"`
	Title     string    `json: "title"`
	Project   string    `json: "project"`
	Tags      []string  `json: "tags"`
	Billable  bool      `json: "billable"`
	StartTime int64     `json: "startTime"`
	EndTime   int64     `json: "endTime"`
	Date      time.Time `json: "date"`
	Duration  int64     `json: "duration"`
}
