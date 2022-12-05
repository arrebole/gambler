package stock

type Deal struct {
	Name   string    `json:"name"`
	Symbol string    `json:"symbol"`
	Decnum int       `json:"decnum"`
	Update int       `json:"update"`
	Amount []int     `json:"amount"`
	Flag   []int     `json:"flag"`
	Price  []float64 `json:"price"`
	Time   []int     `json:"time"`
	Vol    []int     `json:"vol"`
}
