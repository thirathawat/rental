package movie

import "github.com/thirathawat/rental/price"

type M struct {
	title string
	price price.Pricer
}

func New(title string, price price.Pricer) M {
	return M{
		title: title,
		price: price,
	}
}

func (m M) Title() string {
	return m.title
}

func (m M) PriceCode() price.Code {
	return m.price.Code()
}

func (m M) Change(daysRented int) float64 {
	return m.price.Change(daysRented)
}
