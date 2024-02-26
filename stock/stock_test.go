package stock

import (
	"gilded-rose/name"
	"gilded-rose/quality"
	"testing"
)

func TestStock(t *testing.T) {
	t.Run("empty stock should be empty after aging", func(t *testing.T) {
		var stock Stock

		stock.Age()

		assertEmpty(t, stock)
	})

	t.Run("aging the stock should decrease the quality of all items", func(t *testing.T) {
		stock := Stock{
			{
				Name:    val(name.New("Apple")),
				Quality: quality.One,
			},
			{
				Name:    val(name.New("Orange")),
				Quality: quality.Two,
			},
		}

		stock.Age()

		quality.AssertEqual(t, quality.Zero, stock[0].Quality)
		quality.AssertEqual(t, quality.One, stock[1].Quality)
	})
}

func assertEmpty(t *testing.T, stock Stock) {
	t.Helper()
	if len(stock) != 0 {
		t.Errorf("expected %v to be empty", stock)
	}
}

func val[V any](v V, _ error) V {
	return v
}
