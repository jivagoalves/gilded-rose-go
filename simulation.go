package main

import (
	"fmt"
	"gilded-rose/stock"
	"strings"
)

type Simulation struct {
	Stock  stock.Stock
	Output *strings.Builder
}

func (s Simulation) Run(iterations int) {
	for range iterations {
		s.Stock.Age()
		for _, item := range s.Stock {
			s.Output.WriteString(fmt.Sprint(item))
		}
	}
}
