package model

import "fmt"

type Catalog struct {
	products map[string]Product
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

func (c *Catalog) GetProduct(SKU string) (Product, error) {
	product, ok := c.products[SKU]
	if !ok {
		return product, fmt.Errorf("product is not in the catalog")
	}
	return product, nil
}
