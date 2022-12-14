package main

import (
	"encoding/json"
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

	// 查询股票的k线
	// begin、latest 使用 unix 时间精确到秒
	http.HandleFunc("/klines", func(w http.ResponseWriter, r *http.Request) {
		var (
			code      = r.URL.Query().Get("code")
			level     = r.URL.Query().Get("level")
			begin, _  = strconv.Atoi(r.URL.Query().Get("begin"))
			latest, _ = strconv.Atoi(r.URL.Query().Get("latest"))
		)

		store := &storage.Storage{}
		klines, err := store.GetKlines(code, level, begin, latest)
		if err != nil {
			panic(err.Error())
		}

		responseBody, err := json.Marshal(klines)
		if err != nil {
			panic(err.Error())
		}

		w.Write(responseBody)
	})

	// 静态文件服务器
	http.Handle("/", http.FileServer(
		http.Dir("./webapp/dist"),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
