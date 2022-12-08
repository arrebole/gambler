package storage

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/arrebole/gambler/src/stock"
)

const indexFile = "data/index.json"

var dirCache = make(map[string]int)

type Storage struct {
}

// List 查询所有的股票列表
func (p Storage) List() ([]stock.StockBase, error) {
	buffer, err := os.ReadFile(indexFile)
	if err != nil {
		return nil, err
	}

	var result []stock.StockBase
	if err := json.Unmarshal(buffer, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListFiles 查询指定股票的数据文件
func (p Storage) ListFiles(code string) []fs.DirEntry {
	dir := fmt.Sprintf("data/%s", code[0:6])
	files, err := os.ReadDir(dir)
	if err != nil {
		return make([]fs.DirEntry, 0)
	}
	return files
}

// GetFilesRange 查询指定股票的最大和最小时间
func (p Storage) GetFilesRange(code string) (string, string) {
	files := p.ListFiles(code)
	sort.Slice(files, func(i, j int) bool {
		d1, _ := strconv.Atoi(fileDate(files[i].Name()))
		d2, _ := strconv.Atoi(fileDate(files[j].Name()))
		return d1 < d2
	})
	if len(files) == 0 {
		return "", ""
	}
	return fileDate(files[0].Name()), fileDate(files[len(files)-1].Name())
}

// Exists 判断数据是否存在
func (p Storage) Exists(date, code string) bool {
	file, _ := fileName(date, code)
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

// Write 写入数据
func (p Storage) Write(date, code string, data []byte) error {
	file, dir := fileName(date, code)

	if _, ok := dirCache[dir]; !ok {
		os.Mkdir(dir, os.ModePerm)
		dirCache[dir] = 1
	}
	return os.WriteFile(file, data, os.ModePerm)
}

// GetFileTicks 查询指定股票的某一天的交易订单
func (p Storage) GetFileTicks(code string, date string) ([][]float64, error) {
	// 读取文件
	buffer, err := os.ReadFile(
		fmt.Sprintf("data/%s/%s.json", code, date),
	)
	if err != nil {
		return nil, err
	}

	// 序列化
	dailyTicks := &stock.DailyTicks{}
	if err = json.Unmarshal(buffer, &dailyTicks); err != nil {
		return nil, err
	}

	// 写入结果
	result := make([][]float64, 0)
	for i := range dailyTicks.Deal.Time {
		dateString := fmt.Sprintf("%s%06d", date, dailyTicks.Deal.Time[i])
		dateTime, _ := time.ParseInLocation(
			"20060102150405",
			dateString,
			time.Local,
		)
		result = append(result, []float64{
			float64(dateTime.Unix()),
			float64(dailyTicks.Deal.Price[i]),
			float64(dailyTicks.Deal.Vol[i]),
			float64(dailyTicks.Deal.Amount[i]),
			float64(dailyTicks.Deal.Flag[i]),
		})
	}
	return result, nil
}

// GetFilesTicks 查询指定股票的某段时间的交易订单
func (p Storage) GetFilesTicks(code string, begin, end int) ([][]float64, error) {
	var (
		result [][]float64
		files  = p.ListFiles(safeCode(code))
	)
	for _, file := range files {
		var (
			date       = fileDate(file.Name())
			dateInt, _ = strconv.Atoi(date)
		)

		if dateInt > end || dateInt < begin {
			continue
		}

		items, err := p.GetFileTicks(safeCode(code), date)
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
