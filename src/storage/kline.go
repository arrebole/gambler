package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/arrebole/gambler/src/constants"
	"github.com/arrebole/gambler/src/stock"
	"github.com/samber/lo"
)

// geDayKlineData 查询指定股票的某一天的k线数据
// date 格式为 YYYYMMDD
func (p Storage) geDayKlineData(code string, date int) ([]float64, error) {
	// 读取文件
	buffer, err := os.ReadFile(
		fmt.Sprintf("data/%s/%d.json", code, date),
	)
	if err != nil {
		return nil, err
	}

	// 序列化
	dailyTicks := &stock.DailyTicks{}
	if err = json.Unmarshal(buffer, &dailyTicks); err != nil {
		return nil, err
	}

	dateTime, _ := time.ParseInLocation(
		"20060102",
		fmt.Sprint(dailyTicks.Date),
		time.Local,
	)

	result := []float64{
		float64(dateTime.Unix()),
		dailyTicks.Day.Open,
		dailyTicks.Deal.Price[len(dailyTicks.Deal.Price)-1],
		lo.Max(dailyTicks.Deal.Price),
		lo.Min(dailyTicks.Deal.Price),
		float64(lo.Sum(dailyTicks.Deal.Vol)),
	}

	return result, nil
}

// getDaysKlinesData 查询指定股票的某一天的k线数据
// date 格式为 YYYYMMDD
func (p Storage) getDaysKlinesData(code string, begin, latest int) ([][]float64, error) {
	var (
		result [][]float64
		files  = p.ListFiles(safeCode(code))
	)
	for _, file := range files {
		var (
			date       = fileDate(file.Name())
			dateInt, _ = strconv.Atoi(date)
		)

		if dateInt > latest || dateInt < begin {
			continue
		}

		items, err := p.geDayKlineData(safeCode(code), dateInt)
		if err != nil {
			return nil, err
		}

		result = append(
			result,
			items,
		)
	}

	return result, nil
}

// GetFilesKline 查询指定股票的某段时间的K线
func (p Storage) GetKlines(code, level string, begin, latest int) ([][]float64, error) {
	var (
		beginTime           = unixToTime(begin)
		latestTime          = unixToTime(latest)
		beginDateFormat, _  = timeDateFormat(beginTime)
		latestDateFormat, _ = timeDateFormat(latestTime)
	)
	fmt.Println(beginDateFormat, latestDateFormat)

	switch level {

	case
		constants.LEVEL_1MIN,
		constants.LEVEL_15MIN,
		constants.LEVEL_5MIN,
		constants.LEVEL_30MIN:

		// 获取交易订单
		ticks, err := p.getFilesTicks(code, beginDateFormat, latestDateFormat)
		if err != nil {
			return nil, err
		}
		ticks = lo.Filter(ticks, func(v []float64, _ int) bool {
			return int(v[0]) >= begin && int(v[0]) <= latest
		})

		//逐笔合并为其他级别
		switch level {
		case constants.LEVEL_1MIN:
			return merge(ticks), nil
		case constants.LEVEL_5MIN:
			return mergeToMinute(merge((ticks)), 5), nil
		case constants.LEVEL_15MIN:
			return mergeToMinute(merge(ticks), 15), nil
		case constants.LEVEL_30MIN:
			return mergeToMinute(merge(ticks), 30), nil
		}

	case
		constants.LEVEL_1DAY,
		constants.LEVEL_1WEEK,
		constants.LEVEL_1MONTH:

		klines, err := p.getDaysKlinesData(code, beginDateFormat, latestDateFormat)
		if err != nil {
			return nil, err
		}

		// 日级合并为其他级别
		switch level {
		case constants.LEVEL_1DAY:
			return klines, nil
		case constants.LEVEL_1WEEK:
			return mergeToWeek(klines), nil
		case constants.LEVEL_1MONTH:
			return mergeToMonth(klines), nil
		}

	default:
		return nil, errors.New("无效的级别")
	}

	return [][]float64{}, nil
}
