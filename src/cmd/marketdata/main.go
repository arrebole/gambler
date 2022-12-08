package main

import (
	"github.com/arrebole/gambler/src/constants"
	"github.com/arrebole/gambler/src/provider"
	"github.com/arrebole/gambler/src/storage"
)

func main() {

	// 查询所有的股票列表
	store := &storage.Storage{}

	stocks, err := store.List()
	if err != nil {
		panic(err.Error())
	}

	for _, stock := range stocks {

		// 排除创业板和京版的股票
		if stock.Market == constants.MAIN_MARKET || stock.Market == constants.SZSE_MARKET {
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
}
