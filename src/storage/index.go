package storage

import (
	"encoding/json"
	"os"

	"github.com/arrebole/gambler/src/stock"
)

const indexFile = "data/index.json"

// ListAllStocks 查询所有的股票列表
func ListAllStocks() ([]stock.StockBase, error) {
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
