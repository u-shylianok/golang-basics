package model

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/u-shylianok/golang-basics/checkout-system/pkg/file"
)

type Catalog struct {
	products map[string]Product
}

func LoadCatalog(path string) (Catalog, error) {
	bytes, err := file.ReadBytes(path)
	if err != nil {
		logrus.Error(err)
		return Catalog{}, err
	}

	var products []Product

	if err := json.Unmarshal(bytes, &products); err != nil {
		logrus.WithError(err).Error("failed to unmarshal products from json")
		return Catalog{}, err
	}

	productsMap := make(map[string]Product)
	for _, product := range products {
		productsMap[product.SKU] = product
	}
	return Catalog{productsMap}, nil
}

func GetDefaultCatalog() Catalog {
	return Catalog{
		products: map[string]Product{
			"ipd": {
				SKU:   "ipd",
				Name:  "Super iPad",
				Price: 54999,
			},
			"mbp": {
				SKU:   "mbp",
				Name:  "MacBook Pro",
				Price: 139999,
			},
			"atv": {
				SKU:   "atv",
				Name:  "Apple TV",
				Price: 10950,
			},
			"vga": {
				SKU:   "vga",
				Name:  "VGA adapter",
				Price: 3000,
			},
		},
	}
}

func (c *Catalog) Get(productSKU string) (Product, error) {
	product, ok := c.products[productSKU]
	if !ok {
		return product, fmt.Errorf("product is not in the catalog")
	}
	return product, nil
}
