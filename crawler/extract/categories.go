package extract

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	City string `json:"city"`
}
