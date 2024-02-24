package stock

import (
	"gilded-rose/assert"
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
				quality: quality.One,
			},
			{
				name:    "Orange",
				quality: quality.Two,
			},
		}

		Age(stock)

		assert.Equal(t, quality.Zero.Value(), stock[0].quality.Value())
		assert.Equal(t, quality.One.Value(), stock[1].quality.Value())
	})
}

func TestItem(t *testing.T) {
	t.Run("aging an item should decrease the quality", func(t *testing.T) {
		item := Item{
			name:    "Apple",
			quality: quality.Two,
		}

		item.Age()
		item.Age()

		assert.Equal(t, quality.Zero.Value(), item.quality.Value())
	})

	t.Run("the quality of any item should never age below zero", func(t *testing.T) {
		item := Item{
			name:    "Apple",
			quality: quality.Zero,
		}

		item.Age()
		item.Age()

		assert.Equal(t, quality.Zero.Value(), item.quality.Value())
	})
}

func assertEmpty(t *testing.T, stock Stock) {
	t.Helper()
	if len(stock) != 0 {
		t.Errorf("expected %v to be empty", stock)
	}
}
