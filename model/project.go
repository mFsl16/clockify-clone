package model

type Project struct {
	Id       int    `json: "id"`
	Name     string `json: "name"`
	Category string `json: "category"`
	Tracked  int64  `json: "tracked"`
	Progress int    `json: "progress"`
	Access   string `json: "access"`
}
