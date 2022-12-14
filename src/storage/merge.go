package storage

import (
	"sort"
	"strconv"
	"time"

	"github.com/samber/lo"
)

// 将交易订单合并为一分钟k线
func merge(ticks [][]float64) [][]float64 {
	if len(ticks) <= 0 {
		return ticks
	}

	table := lo.GroupBy(ticks, func(f []float64) int {
		return int(f[0]) / 60
	})

	var result [][]float64
	for i, item := range table {
		current := unixToTime(i * 60).Add(time.Minute)
		result = append(result, []float64{
			float64(current.Unix()),
			item[0][1],
			item[len(item)-1][1],
			lo.Max(lo.Map(item, func(it []float64, _ int) float64 {
				return it[1]
			})),
			lo.Min(lo.Map(item, func(it []float64, _ int) float64 {
				return it[1]
			})),
			lo.Sum(lo.Map(item, func(it []float64, _ int) float64 {
				return it[2]
			})),
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

// 将一分钟k线合并为15分钟
func mergeToMinute(kline [][]float64, minute int) [][]float64 {
	if len(kline) <= 0 {
		return kline
	}

	table := lo.GroupBy(kline, func(f []float64) int {
		return int(f[0]) / (minute * 60)
	})

	var result [][]float64
	for i, item := range table {
		current := unixToTime(i * minute * 60).Add(time.Minute * time.Duration(minute))
		result = append(result, []float64{
			float64(current.Unix()),
			item[0][1],
			item[len(item)-1][2],
			lo.Max(lo.Map(item, func(it []float64, _ int) float64 {
				return it[3]
			})),
			lo.Min(lo.Map(item, func(it []float64, _ int) float64 {
				return it[4]
			})),
			lo.Sum(lo.Map(item, func(it []float64, _ int) float64 {
				return it[5]
			})),
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

// 合并日级别 k线 为周级别
func mergeToWeek(kline [][]float64) [][]float64 {

	if len(kline) <= 0 {
		return kline
	}

	table := lo.GroupBy(kline, func(f []float64) int {
		year, week := unixToTime(int(f[0])).ISOWeek()
		return year*100 + week
	})

	var result [][]float64
	for _, item := range table {
		current := endDayOfWeek(
			unixToTime(int(item[0][0])),
		)

		result = append(result, []float64{
			float64(current.Unix()),
			item[0][1],
			item[len(item)-1][2],
			lo.Max(lo.Map(item, func(it []float64, _ int) float64 {
				return it[3]
			})),
			lo.Min(lo.Map(item, func(it []float64, _ int) float64 {
				return it[4]
			})),
			lo.Sum(lo.Map(item, func(it []float64, _ int) float64 {
				return it[5]
			})),
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

// 合并日级别 k线 为月级别
func mergeToMonth(kline [][]float64) [][]float64 {
	if len(kline) <= 0 {
		return kline
	}

	table := lo.GroupBy(kline, func(f []float64) int {
		result, _ := strconv.Atoi(
			unixToTime(int(f[0])).Format("200601"),
		)
		return result
	})

	var result [][]float64
	for k, item := range table {
		current, _ := time.ParseInLocation(
			"200601",
			strconv.Itoa(k),
			time.Local,
		)
		result = append(result, []float64{
			float64(current.AddDate(0, 1, -1).Unix()),
			item[0][1],
			item[len(item)-1][2],
			lo.Max(lo.Map(item, func(it []float64, _ int) float64 {
				return it[3]
			})),
			lo.Min(lo.Map(item, func(it []float64, _ int) float64 {
				return it[4]
			})),
			lo.Sum(lo.Map(item, func(it []float64, _ int) float64 {
				return it[5]
			})),
		})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
