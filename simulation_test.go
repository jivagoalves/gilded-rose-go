package main

import (
	"gilded-rose/assert"
	"gilded-rose/name"
	"gilded-rose/quality"
	"gilded-rose/stock"
	"strings"
	"testing"
)

func TestSimulation(t *testing.T) {
	t.Run("simulation should output the aging progress of the underlying stock", func(t *testing.T) {
		appleName, _ := name.New("Apple")

		sim := Simulation{
			Stock: stock.Stock{
				{
					Name:    appleName,
					Quality: quality.Two,
				},
			},
			Output: &strings.Builder{},
		}

		sim.Run(3)

		assert.Equal(t, "{Apple 1}{Apple 0}{Apple 0}", sim.Output.String())
	})
}
