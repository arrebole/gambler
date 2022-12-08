package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/arrebole/gambler/src/constants"
	"github.com/arrebole/gambler/src/storage"
)

func main() {

	// 获取随机的股票
	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		stock := storage.Random(storage.RandomOptions{
			MinDays: 365,
			Markets: []string{constants.MAIN_MARKET, constants.SZSE_MARKET},
		})

		responseBody, err := json.Marshal(stock)
		if err != nil {
			panic(err.Error())
		}

		w.Write(responseBody)
	})

	// 查询股票的逐笔交易
	http.HandleFunc("/ticks", func(w http.ResponseWriter, r *http.Request) {
		var (
			code     = r.URL.Query().Get("code")
			begin, _ = strconv.Atoi(r.URL.Query().Get("begin"))
			end, _   = strconv.Atoi(r.URL.Query().Get("end"))
		)

		if code == "" || begin == 0 || end == 0 {
			panic(errors.New("缺失必须的参数"))
		}

		store := &storage.Storage{}
		orders, err := store.GetFilesTicks(code, begin, end)
		if err != nil {
			panic(err.Error())
		}

		responseBody, err := json.Marshal(orders)
		if err != nil {
			panic(err.Error())
		}

		w.Write(responseBody)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
