package stock

type StockBase struct {
	Code     string `json:"code"`
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Area     string `json:"area"`
	Industry string `json:"industry"`
	Market   string `json:"market"`
}

type StockInfo struct {
	StockBase
	MinDate string `json:"minDate"`
	MaxDate string `json:"maxDate"`
}
