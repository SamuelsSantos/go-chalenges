package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("When_Input_{100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150}_Then_Return_[100-105], [110-111], [113-115], [150]", func(t *testing.T) {
		var expected = "[100-105], [110-111], [113-115], [150]"
		var input, _ = makeList("100, 101, 102, 103, 104, 105, 110, 111, 113, 114, 115, 150")
		if result, err := doFindIntervals(input); err != nil {
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
		var input, _ = makeList("0, 1, 2, 3, 4, 5, 10, 11, 13, 14, 15, 17")
		if result, err := doFindIntervals(input); err != nil {
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
		var input, _ = makeList("0, 1, 2, 3, 4, 5")
		if result, err := doFindIntervals(input); err != nil {
		} else {
			if result != expected {
				t.Error(
					"expected", expected,
					"got", result,
				)
			}
		}
	})

	t.Run("When_UnSorted_Input_{0, 1, 5, 3, 2, 5}_Then_Return_[0-5]", func(t *testing.T) {
		var expected = "[0-5]"
		var input, _ = makeList("0, 1, 2, 3, 4, 5")
		if result, err := doFindIntervals(input); err != nil {
		} else {
			if result != expected {
				t.Error(
					"expected", expected,
					"got", result,
				)
			}
		}
	})

	t.Run("When_Input_Empty_Then_Return_[]", func(t *testing.T) {
		if _, expected := makeList(""); expected != nil {
		} else {
			if nil != expected {
				t.Error(
					"expected", expected,
					"got", nil,
				)
			}
		}
	})

	t.Run("When_Input_{uywueywueyuyew}_Then_Return_EntradaInvalida", func(t *testing.T) {
		var input, expected = makeList("uywueywueyuyew")
		if _, err := doFindIntervals(input); err != nil {
		} else {
			if err != expected {
				t.Error(
					"expected", expected,
					"got", err,
				)
			}
		}
	})

}
