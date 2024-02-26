package stock

import (
	"gilded-rose/item"
)

type Stock []item.Item

func (s Stock) Age() {
	for i := range s {
		s[i].Age()
	}
}
