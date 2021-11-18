package task

import "github.com/u-shylianok/golang-basics/checkout-system/pkg/model"

func TaskWorkExample(SKUs []string) int64 {
	catalog := model.GetDefaultCatalog()

	var total int64

	for _, SKU := range SKUs {
		total += catalog.Products[SKU].Price
	}

	return total
}
