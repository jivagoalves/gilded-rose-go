package item

import (
	"gilded-rose/name"
	"gilded-rose/quality"
)

type Item struct {
	Name    name.N
	Quality quality.Q
}

func (i *Item) Age() {
	i.Quality = i.Quality.Dec()
}
