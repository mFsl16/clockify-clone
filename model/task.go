package model

import "time"

type Task struct {
	Id        int       `json: "id"; gorm: "primaryKey"; gorm: "autoIncrement"`
	Title     string    `json: "title"`
	Project   string    `json: "project"`
	Tags      string    `json: "tags"`
	Billable  bool      `json: "billable"`
	StartTime time.Time `json: "startTime"`
	EndTime   time.Time `json: "endTime"`
	Date      time.Time `json: "date"`
	Duration  int64     `json: "duration"`
}
