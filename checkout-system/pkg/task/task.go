package task

import "github.com/u-shylianok/golang-basics/checkout-system/pkg/model"

func GetDefaultCatalog() model.Catalog {
	return model.Catalog{
		Products: map[string]model.Product{
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

func TaskWorkExample(SKUs []string) int64 {
	catalog := GetDefaultCatalog()

	var total int64

	for _, SKU := range SKUs {
		total += catalog.Products[SKU].Price
	}

	return total
}
