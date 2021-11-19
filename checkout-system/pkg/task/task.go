package task

import (
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/pricing/rule"
)

func TaskWorkExample(SKUs []string) int64 {
	catalog := model.GetDefaultCatalog()

	rules := rule.GetDefaultRules()

	products := make([]model.Product, len(SKUs))

	for i, SKU := range SKUs {
		products[i] = catalog.Products[SKU]
		products[i].DiscountPrice = products[i].Price
	}

	for _, rule := range rules {
		products = rule.GetDiscountProducts(products)
	}

	var total int64
	for _, product := range products {
		total += product.DiscountPrice
	}

	return total
}
