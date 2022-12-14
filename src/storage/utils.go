package storage

import (
	"fmt"
	"strconv"
	"time"
)

func endDayOfWeek(date time.Time) time.Time {
	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24)*time.Hour).AddDate(0, 0, 4)
	return result
}

func unixToTime(timestamp int) time.Time {
	return time.Unix(int64(timestamp), 0)
}

func timeDateFormat(t time.Time) (int, error) {
	return strconv.Atoi(t.Format("20060102"))
}

func fileName(date, code string) (string, string) {
	dir := fmt.Sprintf("data/%s", safeCode(code))
	file := fmt.Sprintf("%s/%s.json", dir, date)
	return file, dir
}

func fileDate(code string) string {
	return code[0:8]
}

func safeCode(code string) string {
	return code[0:6]
}
