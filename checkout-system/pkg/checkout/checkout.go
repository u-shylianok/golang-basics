package checkout

type Checkout interface {
	Scan(SKU string) error
	Total() (int64, error)
}
