package storage

import (
	"fmt"
	"os"
)

var dirCache = make(map[string]int)

func File(date, code string) (string, string) {
	dir := fmt.Sprintf("data/%s", code[0:6])
	file := fmt.Sprintf("%s/%s.json", dir, date)
	return file, dir
}

func Exists(date, code string) bool {
	file, _ := File(date, code)
	if _, err := os.Stat(file); err == nil {
		return true
	}
	return false
}

func Write(date, code string, data []byte) error {
	file, dir := File(date, code)

	if _, ok := dirCache[dir]; !ok {
		os.Mkdir(dir, os.ModePerm)
		dirCache[dir] = 1
	}

	return os.WriteFile(file, data, os.ModePerm)
}
