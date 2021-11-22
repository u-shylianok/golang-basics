package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/checkout"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/model"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/pricing/rules"
)

func main() {
	fmt.Println("application started")

	verbosePrt := flag.Bool("v", false, "verbose mode - shows all logs information")
	flag.Parse()
	if verbosePrt != nil && *verbosePrt {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}

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
		SKUs := []string{"atv", "atv", "atv", "vga"}
		CheckoutTest(2, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"atv", "ipd", "ipd", "atv", "ipd", "ipd", "ipd"}
		CheckoutTest(2, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"mbp", "vga", "ipd"}
		CheckoutTest(3, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"vga", "vga", "vga", "mbp", "mbp", "ipd"}
		CheckoutTest(4, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"ipd", "vga", "mbp", "vga", "vga", "mbp"}
		CheckoutTest(5, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"ipd", "atv", "vga", "ipd", "mbp", "ipd", "vga", "ipd", "vga", "ipd", "mbp", "atv", "atv", "atv", "ipd"}
		CheckoutTest(6, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv", "atv"}
		CheckoutTest(7, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"ipd", "ipd", "ipd", "ipd"}
		CheckoutTest(8, SKUs, catalog, pricingRules)
	}
	{
		SKUs := []string{"ipd", "ipd", "ipd", "ipd", "ipd"}
		CheckoutTest(9, SKUs, catalog, pricingRules)
	}

	fmt.Println("application closed")
}

func CheckoutTest(testNum int, SKUs []string, catalog model.Catalog, pricingRules []rules.DiscountRuler) {

	fmt.Printf("\n ---------- start test %d ---------- \n", testNum)
	co := checkout.NewCheckout(catalog, pricingRules)
	for _, sku := range SKUs {
		co.Scan(sku)
		fmt.Printf("product [%s] scanned\n", sku)
	}
	co.Total()
	fmt.Printf(" ---------- end test %d ---------- \n", testNum)
}
