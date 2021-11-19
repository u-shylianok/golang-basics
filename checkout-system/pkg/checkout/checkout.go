package checkout

import (
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/pricing/rules"
)

type SKUScanner interface {
	Scan(SKU string) error
	Total() (int64, error)
}

type Checkout struct {
	Catalog      model.Catalog
	PricingRules []rules.DiscountRuler
	Products     []model.Product
}

func NewCheckout(catalog model.Catalog, pricingRules []rules.DiscountRuler) *Checkout {
	return &Checkout{
		Catalog:      catalog,
		PricingRules: pricingRules,
		Products:     []model.Product{},
	}
}

func (c *Checkout) Scan(SKU string) error {

	product, err := c.Catalog.GetProduct(SKU)
	if err != nil {
		// TODO : Log scan error
		return err
	}

	c.Products = append(c.Products, product)
	return nil
}

func (c *Checkout) Total() (int64, error) {

	for _, rule := range c.PricingRules {
		c.Products = rule.GetDiscountProducts(c.Products)
	}

	var total int64
	for _, product := range c.Products {
		total += product.GetTotalPrice()
	}
	return total, nil
}
