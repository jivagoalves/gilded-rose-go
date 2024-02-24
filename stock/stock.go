package stock

type Item struct {
	name    string
	quality int
}

func (i *Item) Age() {
	i.quality--
}

func Age(stock []Item) {
	for i := range stock {
		stock[i].Age()
	}
}
