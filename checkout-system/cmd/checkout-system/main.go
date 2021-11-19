package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/checkout"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/pricing/rules"
)

func main() {
	logrus.Info("start program")

	productsPath := os.Getenv("PRODUCTS_PATH")
	if productsPath == "" {
		logrus.Error("PRODUCTS_PATH environment variable should not be empty")
		// return or GetDefaultCatalog?
		return
	}
	logrus.WithField("path", productsPath).Info("path derived from environment variables")

	catalog, err := model.LoadCatalog(productsPath)
	if err != nil {
		logrus.WithError(err).Error("failed to load catalog")
		// return or GetDefaultCatalog?
		return
	}

	pricingRules := rules.GetDefaultRules()

	{
		logrus.Info(" ---------- start test 1 ---------- ")
		co := checkout.NewCheckout(catalog, pricingRules)
		SKUs := []string{"atv", "atv", "atv", "vga"}
		for _, sku := range SKUs {
			co.Scan(sku)
		}
		co.Total()
		logrus.Info(" ---------- end test 1 ---------- ")
	}
	{
		logrus.Info(" ---------- start test 2 ---------- ")
		co := checkout.NewCheckout(catalog, pricingRules)
		SKUs := []string{"atv", "ipd", "ipd", "atv", "ipd", "ipd", "ipd"}
		for _, sku := range SKUs {
			co.Scan(sku)
		}
		co.Total()
		logrus.Info(" ---------- end test 2 ---------- ")
	}
	{
		logrus.Info(" ---------- start test 3 ---------- ")
		co := checkout.NewCheckout(catalog, pricingRules)
		SKUs := []string{"mbp", "vga", "ipd"}
		for _, sku := range SKUs {
			co.Scan(sku)
		}
		co.Total()
		logrus.Info(" ---------- end test 3 ---------- ")
	}
	logrus.Info("end program")
}
