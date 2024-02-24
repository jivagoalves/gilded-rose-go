package stock

import (
	"gilded-rose/quality"
	"testing"
)

func TestStock(t *testing.T) {
	t.Run("empty stock should be empty after aging", func(t *testing.T) {
		var stock Stock

		Age(stock)

		assertEmpty(t, stock)
	})

	t.Run("aging the stock should decrease the quality of all items", func(t *testing.T) {
		stock := Stock{
			{
				name:    "Apple",
				quality: quality.Q(1),
			},
			{
				name:    "Orange",
				quality: quality.Q(2),
			},
		}

		Age(stock)

		assertEqual(t, quality.Zero, stock[0].quality)
		assertEqual(t, quality.Q(1), stock[1].quality)
	})
}

func TestItem(t *testing.T) {
	t.Run("aging an item should decrease the quality", func(t *testing.T) {
		item := Item{
			name:    "Apple",
			quality: quality.Q(2),
		}

		item.Age()
		item.Age()

		assertEqual(t, quality.Zero, item.quality)
	})

	t.Run("the quality of any item should never age below zero", func(t *testing.T) {
		item := Item{
			name:    "Apple",
			quality: quality.Zero,
		}

		item.Age()
		item.Age()

		assertEqual(t, quality.Zero, item.quality)
	})
}

func assertEmpty(t *testing.T, stock Stock) {
	t.Helper()
	if len(stock) != 0 {
		t.Errorf("expected %v to be empty", stock)
	}
}

func assertEqual[V comparable](t *testing.T, expected, actual V) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
