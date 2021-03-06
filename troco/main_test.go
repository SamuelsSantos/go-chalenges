package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {

	t.Run("When_GetCoins_The_Result_List_Of_Coins", func(t *testing.T) {
		var expected = 11
		var result = len(getCoins())
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_GetChange_{5.78}_The_Result_[{0.25}:1 {0.5}:1 {5}:1]", func(t *testing.T) {
		changeCoins := &ChangesCoins{}
		changeCoins.getChange(5.78, getCoins())
		var expected = ChangesCoins{}
		expected[Coin{value: 5.0}] = 1
		expected[Coin{value: 0.25}] = 1
		expected[Coin{value: 0.5}] = 1

		if len(*changeCoins) != len(expected) {
			t.Error(
				"expected", len(expected),
				"got", len(*changeCoins),
			)
		}

		for coin, amount := range *changeCoins {
			if expected[coin] != amount {
				t.Error(
					"expected", expected[coin],
					"got", amount,
				)
			}
		}
	})

	t.Run("When_GetChange_{15.0}_The_Result_map[{5}:1 {10}:1]", func(t *testing.T) {
		changeCoins := &ChangesCoins{}
		changeCoins.getChange(15.0, getCoins())
		var expected = ChangesCoins{}
		expected[Coin{value: 10.0}] = 1
		expected[Coin{value: 5.0}] = 1

		if len(*changeCoins) != len(expected) {
			t.Error(
				"expected", len(expected),
				"got", len(*changeCoins),
			)
		}

		for coin, amount := range *changeCoins {
			if expected[coin] != amount {
				t.Error(
					"expected", expected[coin],
					"got", amount,
				)
			}
		}
	})

	t.Run("When_GetChange_{14.20}_The_Result_map[{0.1}:2 {2}:2 {10}:1]", func(t *testing.T) {
		changeCoins := &ChangesCoins{}
		changeCoins.getChange(14.20, getCoins())
		var expected = ChangesCoins{}
		expected[Coin{value: 10.0}] = 1
		expected[Coin{value: 2.0}] = 2
		expected[Coin{value: 0.10}] = 2

		if len(*changeCoins) != len(expected) {
			t.Error(
				"expected", len(expected),
				"got", len(*changeCoins),
			)
		}

		for coin, amount := range *changeCoins {
			if expected[coin] != amount {
				t.Error(
					"expected", expected[coin],
					"got", amount,
				)
			}
		}
	})

	t.Run("When_GetChange_{0.00}_The_Result_map[]", func(t *testing.T) {
		changeCoins := &ChangesCoins{}
		changeCoins.getChange(14.20, getCoins())
		var expected = ChangesCoins{}
		expected[Coin{value: 10.0}] = 1
		expected[Coin{value: 2.0}] = 2
		expected[Coin{value: 0.10}] = 2

		if len(*changeCoins) != len(expected) {
			t.Error(
				"expected", len(expected),
				"got", len(*changeCoins),
			)
		}
	})

	t.Run("When_FormatFloat_{4.19999999999}_WithPrecision{2}_The_Result_{4.20}", func(t *testing.T) {
		var expected = 4.20
		var result, _ = formatFloat(4.19999999999, 2)
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_DoProcess{Value:35.0-Paid:50.0}_The_Result_[{5}:1 {10}:1]", func(t *testing.T) {

		var expected = 15.0
		result, _, err := doProcess("35.0", "50.0")

		if err != nil {
			t.Error(
				"expected", nil,
				"got", err,
			)
		}

		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_DoProcess{Empty,Value}_The_Result_Error", func(t *testing.T) {

		var expected error
		_, _, err := doProcess("", "60")

		if err == nil {
			t.Error(
				"expected", expected,
				"got", err,
			)
		}
	})

	t.Run("When_DoProcess{Empty,Paid}_The_Result_Error", func(t *testing.T) {

		var expected error
		_, _, err := doProcess("40.0", "")

		if err == nil {
			t.Error(
				"expected", expected,
				"got", err,
			)
		}
	})

	t.Run("When_DoProcess{Empty}_The_Result_Error", func(t *testing.T) {

		var expected error
		_, _, err := doProcess("40.0", "35.00")

		if err == nil {
			t.Error(
				"expected", expected,
				"got", err,
			)
		}
	})

	t.Run("When_DoProcess{Empty}_The_Result_Error", func(t *testing.T) {

		var stdin bytes.Buffer
		stdin.Write([]byte("30.00"))
		var expected = "30.00"
		var result string
		if result = readInfo(&stdin, "Value"); result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

}
