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

	var products []ProductJSON

	if err := json.Unmarshal(bytes, &products); err != nil {
		logrus.WithError(err).Error("failed to unmarshal products from json")
		return Catalog{}, err
	}

	productsMap := make(map[string]Product)
	for _, product := range products {
		productsMap[product.SKU] = NewProduct(product.SKU, product.Name, product.Price)
	}
	return Catalog{productsMap}, nil
}

func GetDefaultCatalog() Catalog {
	return Catalog{
		products: map[string]Product{
			"ipd": NewProduct("ipd", "Super iPad", 54999),
			"mbp": NewProduct("mbp", "MacBook Pro", 139999),
			"atv": NewProduct("atv", "Apple TV", 10950),
			"vga": NewProduct("vga", "VGA adapter", 3000),
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
