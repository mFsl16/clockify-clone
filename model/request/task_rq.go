package request

type TaskRq struct {
	Title     string `json:"title"`
	Project   string `json:"project"`
	Tags      string `json:"tags"`
	Billable  bool   `json:"billable"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Date      string `json:"date"`
	Duration  int64  `json:"duration"`
}
