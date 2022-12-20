package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/arrebole/gambler/src/stock"
)

// GetFileTicks 查询指定股票的某一天的交易订单
// date 格式为 YYYYMMDD
func (p Storage) getFileTicks(code string, date int) ([][]float64, error) {
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

	return dailyTicks.GetTicks(), nil
}

// GetFilesTicks 查询指定股票的某段时间的交易订单
// begin latest 格式为 YYYYMMDD
func (p Storage) getFilesTicks(code string, begin, latest int) ([][]float64, error) {
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

		items, err := p.getFileTicks(safeCode(code), dateInt)
		if err != nil {
			return nil, err
		}

		result = append(
			result,
			items...,
		)
	}

	return result, nil
}
