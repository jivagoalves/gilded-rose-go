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
		sim := Simulation{
			Stock: stock.Stock{
				{
					Name:    val(name.New("Apple")),
					Quality: quality.Two,
				},
				{
					Name:    val(name.New("Orange")),
					Quality: val(quality.New(5)),
				},
			},
			Output: &strings.Builder{},
		}.Run(4)

		expected := `{Apple 2}{Orange 5}
{Apple 1}{Orange 4}
{Apple 0}{Orange 3}
{Apple 0}{Orange 2}
`
		assert.Equal(t, expected, sim.Output.String())
	})
}
