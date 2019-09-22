package model

type Merchants struct {
	ID         int64   `json:"ID"`
	Name       string  `json:"Name"`
	Address    string  `json:"Address"`
	Rating     float64 `json:"Rating"`
	FavePayCnt int     `json:"FavePayCnt"`
	City       string  `json:"City"`
	Category   string  `json:"Category"`
	Logo       string  `json:"Logo"`
}
