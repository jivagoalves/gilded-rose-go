package main

import (
	"fmt"
	"gilded-rose/name"
	"gilded-rose/quality"
	"gilded-rose/stock"
	"strings"
)

func main() {
	s := Simulation{
		Stock: stock.Stock{
			{
				Name:    val(name.New("Apple")),
				Quality: val(quality.New(12)),
			},
			{
				Name:    val(name.New("Orange")),
				Quality: val(quality.New(10)),
			},
		},
		Output: &strings.Builder{},
	}.Run(10)

	fmt.Println(`
=========
  Stock
=========`)

	fmt.Println(s.Output.String())
}

func val[V any](v V, _ error) V {
	return v
}
