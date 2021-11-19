package checkout

import (
	"github.com/sirupsen/logrus"
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

	product, err := c.Catalog.Get(SKU)
	if err != nil {
		logrus.WithError(err).Error("failed to get product from catalog")
		return err
	}
	logrus.WithField("SKU", product.SKU).Info("product scanned")

	c.Products = append(c.Products, product)
	return nil
}

func (c *Checkout) Total() (int64, error) {
	logrus.Info("TOTAL:")

	for _, rule := range c.PricingRules {
		c.Products = rule.GetDiscountProducts(c.Products)
	}

	var total int64
	for _, product := range c.Products {
		total += product.GetTotalPrice()
		logrus.Info(product)
	}
	return total, nil
}
