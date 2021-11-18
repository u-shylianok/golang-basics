package model

type Catalog struct {
	Products map[string]Product
}

func GetDefaultCatalog() Catalog {
	return Catalog{
		Products: map[string]Product{
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
