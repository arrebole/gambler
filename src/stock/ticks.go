package stock

import (
	"fmt"
	"time"

	"github.com/samber/lo"
)

type DailyTicks struct {
	Date   int   `json:"date"`
	Day    Quote `json:"day"`
	Deal   Deal  `json:"deal"`
	Minute Deal  `json:"minute"`
}

// 获取当前交易日的日级别k线
func (p DailyTicks) GetDayKline() []float64 {
	if len(p.Deal.Price) == 0 {
		dateTime, _ := time.ParseInLocation(
			"200601021504",
			fmt.Sprintf("%d%04d", p.Date, p.Minute.Time[len(p.Minute.Time)-1]),
			time.Local,
		)
		return []float64{
			float64(dateTime.Unix()),
			p.Day.Open,
			p.Minute.Price[len(p.Minute.Price)-1],
			lo.Max(p.Minute.Price),
			lo.Min(p.Minute.Price),
			float64(lo.Sum(p.Minute.Vol)),
		}
	}

	dateTime, _ := time.ParseInLocation(
		"20060102150405",
		fmt.Sprintf("%d%06d", p.Date, p.Deal.Time[len(p.Deal.Time)-1]),
		time.Local,
	)
	return []float64{
		float64(dateTime.Unix()),
		p.Day.Open,
		p.Deal.Price[len(p.Deal.Price)-1],
		lo.Max(p.Deal.Price),
		lo.Min(p.Deal.Price),
		float64(lo.Sum(p.Deal.Vol)),
	}
}

func (p DailyTicks) GetTicks() [][]float64 {
	result := make([][]float64, 0)
	if len(p.Deal.Price) == 0 {
		for i := range p.Minute.Time {
			dateString := fmt.Sprintf("%d%04d", p.Date, p.Minute.Time[i])
			dateTime, _ := time.ParseInLocation(
				"200601021504",
				dateString,
				time.Local,
			)
			result = append(result, []float64{
				float64(dateTime.Unix()),
				float64(p.Minute.Price[i]),
				float64(p.Minute.Vol[i]),
				float64(p.Minute.Amount[i]),
				float64(p.Minute.Flag[i]),
			})
		}
		return result
	}

	for i := range p.Deal.Time {
		dateString := fmt.Sprintf("%d%06d", p.Date, p.Deal.Time[i])
		dateTime, _ := time.ParseInLocation(
			"20060102150405",
			dateString,
			time.Local,
		)
		result = append(result, []float64{
			float64(dateTime.Unix()),
			float64(p.Deal.Price[i]),
			float64(p.Deal.Vol[i]),
			float64(p.Deal.Amount[i]),
			float64(p.Deal.Flag[i]),
		})
	}
	return result
}
