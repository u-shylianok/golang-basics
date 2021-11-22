package rules

import (
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/counter"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
)

type DiscountRuler interface {
	GetDiscountProducts(products []model.Product) []model.Product
}

func GetDefaultRules() []DiscountRuler {
	return []DiscountRuler{
		&DiscountEveryMaxCountRule{
			Counter:  counter.NewSafeCounter(),
			FromSKU:  "atv",
			MaxCount: 3,
			ToPrice:  0,
		},
		&DiscountMoreThanCountRule{
			Counter:       counter.NewSafeCounter(),
			FromSKU:       "ipd",
			ToSKU:         "ipd",
			MoreThanCount: 4,
			ToPrice:       49999,
		},
		&DiscountOneToOneRule{
			Counter: counter.NewSafeCounter(),
			FromSKU: "mbp",
			ToSKU:   "vga",
			ToPrice: 0,
		},
	}
}
