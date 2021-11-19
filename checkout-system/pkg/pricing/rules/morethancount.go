package rules

import (
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/counter"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
)

// условие скидки - если количество FromSKU больше MoreThanCount - применяется цена ToPrice ко всем ToSKU
type DiscountMoreThanCountRule struct {
	Counter counter.Counter

	FromSKU       string
	ToSKU         string
	MoreThanCount int
	ToPrice       int64
}

func (r *DiscountMoreThanCountRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU == r.FromSKU {
			r.Counter.Inc()
		}
	}
	if r.Counter.Get() > r.MoreThanCount {
		for i := range products {
			if products[i].SKU == r.ToSKU {
				products[i].SetDiscount(r.ToPrice)
			}
		}
	}
	return products
}
