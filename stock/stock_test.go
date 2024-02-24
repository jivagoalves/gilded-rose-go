package stock

import (
	"testing"
)

func TestStock(t *testing.T) {
	t.Run("empty stock should be empty after aging", func(t *testing.T) {
		var stock []Item

		Age(stock)

		assertEmpty(t, stock)
	})

	t.Run("aging the stock should decrease the quality of all items", func(t *testing.T) {
		stock := []Item{
			{
				name:    "Apple",
				quality: 1,
			},
			{
				name:    "Orange",
				quality: 2,
			},
		}

		Age(stock)

		assertEqual(t, 0, stock[0].quality)
		assertEqual(t, 1, stock[1].quality)
	})
}

func TestItem(t *testing.T) {
	t.Run("aging an item should decrease the quality", func(t *testing.T) {
		item := Item{
			name:    "Apple",
			quality: 2,
		}

		item.Age()
		item.Age()

		assertEqual(t, 0, item.quality)
	})
}

func assertEmpty(t *testing.T, stock []Item) {
	t.Helper()
	if len(stock) != 0 {
		t.Errorf("expected %v to be empty", stock)
	}
}

func assertEqual(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %d, actual %d", expected, actual)
	}
}
