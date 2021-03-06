package rules

import (
	"fmt"

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
	r.Counter.Set(0)
	for i := range products {
		if products[i].SKU() == r.FromSKU {
			r.Counter.Inc()
		}
	}
	if r.Counter.Get() > r.MoreThanCount {
		for i := range products {
			if products[i].SKU() == r.ToSKU {
				products[i].SetDiscount(r.ToPrice)
			}
		}
	}
	return products
}

func (r DiscountMoreThanCountRule) String() string {
	return fmt.Sprintf("if [%s] products count more than %d, set %s to all [%s] products", r.FromSKU, r.MoreThanCount, model.DollarsString(r.ToPrice), r.ToSKU)
}
