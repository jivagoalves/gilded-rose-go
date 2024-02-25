package item

import (
	"gilded-rose/name"
	"gilded-rose/quality"
	"testing"
)

func TestItem(t *testing.T) {
	appleName, _ := name.New("Apple")

	t.Run("aging an item should decrease the quality", func(t *testing.T) {
		apple := Item{
			Name:    appleName,
			Quality: quality.Two,
		}

		apple.Age()
		apple.Age()

		quality.AssertEqual(t, quality.Zero, apple.Quality)
	})

	t.Run("the quality of any item should never age below zero", func(t *testing.T) {
		apple := Item{
			Name:    appleName,
			Quality: quality.Zero,
		}

		apple.Age()
		apple.Age()

		quality.AssertEqual(t, quality.Zero, apple.Quality)
	})
}
