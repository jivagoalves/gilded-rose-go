package stock

import (
	"gilded-rose/item"
)

type Stock []item.Item

func Age(stock Stock) {
	for i := range stock {
		stock[i].Age()
	}
}
