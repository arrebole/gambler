package storage

import (
	"fmt"
)

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
