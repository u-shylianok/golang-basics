package rule

import "github.com/u-shylianok/golang-basics/checkout-system/pkg/model"

type Rule interface {
	GetDiscountProducts(products []model.Product) []model.Product
}

func GetDefaultRules() []Rule {
	return []Rule{
		&CounterToOneRule{
			FromSKU:    "atv",
			MaxCounter: 3,
			ToPrice:    0,
		},
		&CounterToAllRule{
			FromSKU:         "ipd",
			MoreThanCounter: 4,
			ToPrice:         49999,
		},
		&EveryFromToRule{
			FromSKU: "mbp",
			ToSKU:   "vga",
			ToPrice: 0,
		},
	}
}

type CounterToOneRule struct {
	FromSKU    string
	Counter    int
	MaxCounter int
	ToPrice    int64
}

func (r *CounterToOneRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU == r.FromSKU {
			r.Counter++
			if r.MaxCounter == r.Counter {
				r.Counter = 0
				products[i].DiscountPrice = r.ToPrice
			}
		}
	}
	return products
}

type CounterToAllRule struct {
	FromSKU         string
	Counter         int
	MoreThanCounter int
	ToPrice         int64
}

func (r *CounterToAllRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU == r.FromSKU {
			r.Counter++
		}
	}
	if r.Counter > r.MoreThanCounter {
		for i := range products {
			if products[i].SKU == r.FromSKU {
				products[i].DiscountPrice = r.ToPrice
			}
		}
	}
	return products
}

type EveryFromToRule struct {
	FromSKU string
	ToSKU   string
	Counter int
	ToPrice int64
}

func (r *EveryFromToRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU == r.FromSKU {
			r.Counter++
		}
	}
	if r.Counter == 0 {
		return products
	}

	for i := range products {
		if r.Counter == 0 {
			break
		}

		if products[i].SKU == r.ToSKU {
			products[i].DiscountPrice = r.ToPrice
			r.Counter--
		}
	}
	return products
}
