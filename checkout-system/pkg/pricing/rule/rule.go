package rule

import "github.com/u-shylianok/golang-basics/checkout-system/pkg/model"

type Rule interface {
	GetDiscountProducts(products []model.Product) []model.DiscountProduct
}
