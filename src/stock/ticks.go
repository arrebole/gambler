package stock

type DailyTicks struct {
	Date int   `json:"date"`
	Day  Quote `json:"day"`
	Deal Deal  `json:"deal"`
}
