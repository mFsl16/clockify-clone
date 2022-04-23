package request

type ProjectRq struct {
	Name     string `json: "name"`
	Category string `json: "category"`
	Tracked  int64  `json: "tracked"`
	Progress int    `json: "progress"`
	Access   string `json: "access"`
}
