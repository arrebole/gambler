package stock

type Quote struct {
	Open   float64 `json:"open"`
	Price  float64 `json:"price"`
	Yclose float64 `json:"yclose"`
}
