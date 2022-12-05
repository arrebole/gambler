package main

import (
	"github.com/arrebole/gambler/src/provider"
	"github.com/arrebole/gambler/src/storage"
)

func main() {

	// 查询所有的股票列表
	stocks, err := storage.ListAllStocks()
	if err != nil {
		panic(err.Error())
	}

	for _, stock := range stocks {

		// 排除创业板和京版的股票
		if stock.Market == "北交所" || stock.Market == "创业板" {
			continue
		}

		// 查询和保存股票时间段内的所有数据
		err := provider.FetchAndSaveMarketData(
			stock.Code,
			"20160101",
			"20221130",
		)
		if err != nil {
			panic(err.Error())
		}
	}
}
