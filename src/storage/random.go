package storage

import (
	"math/rand"
	"time"

	"github.com/arrebole/gambler/src/stock"
)

type RandomOptions struct {
	Markets []string
	MinDays int
}

// Random 随机返回一支股票
func Random(option RandomOptions) *stock.StockInfo {
	store := &Storage{}

	stocks, err := store.List()
	if err != nil {
		return nil
	}

	var items []stock.StockBase
	for _, stock := range stocks {
		marketOk := false
		for _, market := range option.Markets {
			if stock.Market == market {
				marketOk = true
				break
			}
		}
		if !marketOk {
			continue
		}
		if len(store.ListFiles(safeCode(stock.Code))) < option.MinDays {
			continue
		}
		items = append(items, stock)
	}

	if len(items) == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	stockBase := items[rand.Intn(len(items))]
	minDate, maxDate := store.GetFilesRange(stockBase.Code)

	return &stock.StockInfo{
		StockBase: stockBase,
		MinDate:   minDate,
		MaxDate:   maxDate,
	}

}
