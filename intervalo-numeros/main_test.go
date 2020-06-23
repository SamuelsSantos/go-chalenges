package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("When_Input_{100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150}_Then_Return_[100-105], [110-111], [113-115], [150]", func(t *testing.T) {
		var expected = "[100-105], [110-111], [113-115], [150]"
		if result, err := agrupar("100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150"); err != nil {
		} else {
			if result != expected {
				t.Error(
					"expected", expected,
					"got", result,
				)
			}
		}

	})

	t.Run("When_Input_{0, 1, 2, 3, 4, 5, 10, 11, 13, 14, 15, 17}_Then_Return_[0-5], [10-11], [13-15], [17]", func(t *testing.T) {
		var expected = "[0-5], [10-11], [13-15], [17]"
		if result, err := agrupar("0, 1, 2, 3, 4, 5, 10, 11, 13, 14, 15, 17"); err != nil {
		} else {
			if result != expected {
				t.Error(
					"expected", expected,
					"got", result,
				)
			}
		}

	})

	t.Run("When_Input_{0, 1, 2, 3, 4, 5}_Then_Return_[0-5]", func(t *testing.T) {
		var expected = "[0-5]"
		if result, err := agrupar("0, 1, 2, 3, 4, 5"); err != nil {
		} else {
			if result != expected {
				t.Error(
					"expected", expected,
					"got", result,
				)
			}
		}

	})

}
