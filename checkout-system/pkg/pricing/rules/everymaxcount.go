package rules

import (
	"fmt"

	"github.com/u-shylianok/golang-basics/checkout-system/pkg/counter"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
)

// условие скидки - каждый (i == MaxCount) товар FromSKU будет пременена скидка ToPrice
type DiscountEveryMaxCountRule struct {
	FromSKU  string
	Counter  counter.Counter
	MaxCount int
	ToPrice  int64
}

func (r *DiscountEveryMaxCountRule) GetDiscountProducts(products []model.Product) []model.Product {

	for i := range products {
		if products[i].SKU() == r.FromSKU {
			r.Counter.Inc()
			if r.MaxCount == r.Counter.Get() {
				r.Counter.Set(0)
				products[i].SetDiscount(r.ToPrice)
			}
		}
	}
	return products
}

func (r DiscountEveryMaxCountRule) String() string {
	return fmt.Sprintf("every (i == %d) product [%s] set new price = %s", r.MaxCount, r.FromSKU, model.ToDollars(r.ToPrice))
}
