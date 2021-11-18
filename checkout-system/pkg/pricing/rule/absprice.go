package rule

import "github.com/u-shylianok/golang-basics/checkout-system/pkg/model"

type AbsolutePrice struct {
	FromSKU   string
	ToSKU     string
	ToPrice   int64
	Condition func() bool
}

func (p *AbsolutePrice) GetDiscountProducts(products []model.Product) []model.DiscountProduct {
	var discountProducts []model.DiscountProduct

	for _, product := range products {
		if product.Name == p.FromSKU {
			if p.Condition() {
				discountProducts = append(discountProducts, product.ToDiscountProduct(p.ToPrice))
			}
		} else {
			discountProducts = append(discountProducts, product.ToDiscountProduct(product.Price))
		}
	}
	return discountProducts
}
