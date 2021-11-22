package rules

import (
	"fmt"

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
	r.Counter.Set(0)
	for i := range products {
		if products[i].SKU() == r.FromSKU {
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

		if products[i].SKU() == r.ToSKU {
			products[i].SetDiscount(r.ToPrice)
			r.Counter.Dec()
		}
	}
	return products
}

func (r DiscountOneToOneRule) String() string {
	return fmt.Sprintf("for every one [%s] product set price %s to one [%s] product", r.FromSKU, model.DollarsString(r.ToPrice), r.ToSKU)
}
