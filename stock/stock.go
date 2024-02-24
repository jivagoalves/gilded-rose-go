package stock

import "gilded-rose/quality"

type Item struct {
	name    string
	quality quality.Q
}

type Stock []Item

func (i *Item) Age() {
	if i.quality == quality.Zero {
		return
	}
	i.quality = i.quality.Dec()
}

func Age(stock Stock) {
	for i := range stock {
		stock[i].Age()
	}
}
