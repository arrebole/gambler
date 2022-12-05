package provider

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/arrebole/gambler/src/storage"
)

type TaskAction struct {
	Code string
	Date string
}

func consume(queue <-chan TaskAction, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range queue {

		// 查询该股票的该天行情
		buffer, err := fetchDailyTick(task.Date, task.Code)
		if err != nil {
			log.Println("ERROR " + err.Error())
			continue
		}

		// 将数据写入磁盘
		if err = storage.Write(task.Date, task.Code, buffer); err != nil {
			log.Println("ERROR " + err.Error())
		}
	}
}

func fetchDailyTick(day, code string) ([]byte, error) {
	resp, err := http.Get(
		fmt.Sprintf(
			"https://opensourcecache.zealink.com/cache/dealday/day/%s/%s.json",
			day,
			code,
		),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"查询交易行情失败 [StatusCode=%d code=%s date=%s]",
			resp.StatusCode,
			code,
			day,
		)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func FetchAndSaveMarketData(code, lo, hi string) error {
	latest, err := time.Parse("20060102", hi)
	if err != nil {
		return err
	}
	point, err := time.Parse("20060102", lo)
	if err != nil {
		return err
	}

	// 创建管道和线程控制组
	var wg sync.WaitGroup
	channel := make(chan TaskAction, 100)

	// 创建 5 个线程执行任务
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go consume(channel, &wg)
	}

	// 下发任务
	for latest.After(point) {
		date := point.Format("20060102")
		point = point.Add(time.Hour * 24)

		channel <- TaskAction{
			Code: code,
			Date: date,
		}
	}

	close(channel)
	wg.Wait()

	return nil
}
