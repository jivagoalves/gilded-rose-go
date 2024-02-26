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

func (s Simulation) Run(iterations int) Simulation {
	for range iterations {
		for _, item := range s.Stock {
			s.Output.WriteString(fmt.Sprint(item))
		}
		s.Output.WriteString("\n")
		s.Stock.Age()
	}
	return s
}
