package rules

import (
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/counter"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
)

// условие скидки - на каждый один товар FromSKU - применяется цена ToPrice к одному ToSKU
type DiscountOneToOneRule struct {
	FromSKU string
	ToSKU   string
	Counter counter.Counter
	ToPrice int64
}

func (r *DiscountOneToOneRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU == r.FromSKU {
			r.Counter.Inc()
		}
	}
	if r.Counter.Get() == 0 {
		return products
	}

	for i := range products {
		if r.Counter.Get() == 0 {
			break
		}

		if products[i].SKU == r.ToSKU {
			products[i].SetDiscount(r.ToPrice)
			r.Counter.Dec()
		}
	}
	return products
}
