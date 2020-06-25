package main

import (
	"testing"
)

type MockNumberError struct{}
type MockNumberNormal struct{}
type MockNumberLuck struct{}
type MockNumberCritical struct{}

func (MockNumberError) Get() int {
	return 1
}

func (MockNumberNormal) Get() int {
	return 16
}

func (MockNumberLuck) Get() int {
	return 71
}

func (MockNumberCritical) Get() int {
	return 97
}
func TestMain(t *testing.T) {

	t.Run("Luck Factor Normal", func(t *testing.T) {

		player1 := Player{Nome: "Player 1", Energy: 40, Power: 50}
		player2 := Player{Nome: "Player 2", Energy: 30, Power: 40}
		normal := MockNumberNormal{}
		var expected = player1

		result := round(&player1, &player2, normal)
		if result.Nome != expected.Nome {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}

		if player2.Energy > 0 {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}
	})

	t.Run("Luck Factor", func(t *testing.T) {

		player1 := Player{Nome: "Player 1", Energy: 20, Power: 50}
		player2 := Player{Nome: "Player 2", Energy: 10, Power: 40}
		luck := MockNumberLuck{}
		var expected = player1

		result := round(&player1, &player2, luck)
		if result.Nome != expected.Nome {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}

		if player2.Energy > 0 {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}
	})

	t.Run("Luck Factor Critical", func(t *testing.T) {

		player1 := Player{Nome: "Player 1", Energy: 20, Power: 40}
		player2 := Player{Nome: "Player 2", Energy: 20, Power: 60}
		critical := MockNumberCritical{}
		var expected = player2

		result := round(&player1, &player2, critical)
		if result.Nome != expected.Nome {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}

		if player1.Energy > 0 {
			t.Error(
				"expected", expected.Nome,
				"got", result.Nome,
			)
		}
	})

	t.Run("When_Check_IsEmpty_Then_True", func(t *testing.T) {

		var expected = true
		var result = isEmpty(" ")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Check_IsEmpty_Then_False", func(t *testing.T) {

		var expected = false
		var result = isEmpty(" A ")
		if result != expected {
			t.Error(
				"expected", expected,
				"got", result,
			)
		}
	})

	t.Run("When_Check_IsValidInputPersonagem_Then_Error", func(t *testing.T) {

		var expect error
		var err error
		_, err = isValidInput("A ")
		if err == nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}

		_, err = isValidInput("")
		if err == nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}
	})

	t.Run("When_Check_IsValidInput_Then_ArrayPropertiesPersonagem", func(t *testing.T) {

		var expect []string = []string{"A", "1", "1"}
		var err error
		var result []string

		result, err = isValidInput("A 1 1")
		if err != nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}

		if result == nil {
			t.Error(
				"expected", expect,
				"got", result,
			)
		}

	})

	t.Run("When_ReadStringPlayer_Then_Player", func(t *testing.T) {

		var expect error

		result, err := readPlayer("A 10 20")
		if err != nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}

		if result.Nome != "A" {
			t.Error(
				"expected", "A",
				"got", result,
			)
		}

		if result.Energy != 10 {
			t.Error(
				"expected", 10,
				"got", result,
			)
		}

		if result.Power != 20 {
			t.Error(
				"expected", 10,
				"got", result,
			)
		}
	})

	t.Run("When_ReadStringEmpty_Then_Error", func(t *testing.T) {

		var expect error
		var err error
		_, err = readPlayer(" ")
		if err == nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}
	})

	t.Run("When_ReadStringPlayerWithoutValidEnergy_Then_Error", func(t *testing.T) {

		var expect error
		var err error

		_, err = readPlayer("A A 1")
		if err == nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}
	})

	t.Run("When_ReadStringPlayerWithoutValidPower_Then_Error", func(t *testing.T) {

		var expect error
		var err error

		_, err = readPlayer("A 1 A")
		if err == nil {
			t.Error(
				"expected", expect,
				"got", err,
			)
		}
	})

}
