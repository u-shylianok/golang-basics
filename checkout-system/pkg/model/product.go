package model

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Product struct {
	sku      string
	name     string
	price    int64
	discount int64
	total    int64
}

type ProductJSON struct {
	SKU   string `json:"SKU"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func NewProduct(sku, name string, price int64) Product {
	return Product{
		sku:   sku,
		name:  name,
		price: price,
		total: price,
	}
}

func (p *Product) SKU() string {
	return p.sku
}

func (p *Product) Name() string {
	return p.name
}

func (p *Product) Price() int64 {
	return p.price
}

func (p *Product) Discount() int64 {
	return p.discount
}

func (p *Product) SetDiscount(newPrice int64) {
	discount := p.price - newPrice
	if discount <= 0 || discount > p.price {
		logrus.Info("trying to set an unreal discount - skip")
		return
	}
	if discount < p.discount {
		logrus.Info("trying to set a lower discount - skip")
		return
	}
	logrus.WithFields(logrus.Fields{
		"total":    DollarsString(newPrice),
		"discount": DollarsString(-discount),
	}).Info("discount set successfully")
	p.discount = discount
	p.total = newPrice
}

func (p *Product) Total() int64 {
	return p.total
}

func (p Product) String() string {
	price := DollarsString(p.price)
	total := DollarsString(p.Total())
	if p.discount == 0 {
		return fmt.Sprintf("\nProduct: %s\nSKU: [%s]\nPrice: %s\nTotal: %s", p.name, p.sku, price, total)
	}

	discount := DollarsString(p.discount)
	return fmt.Sprintf("\nProduct: %s\nSKU: %s\nPrice: %s\nDiscount: -%s\nTotal: %s", p.name, p.sku, price, discount, total)
}

func DollarsString(price int64) string {
	return fmt.Sprintf("%.2f$", float64(price)/100)
}
