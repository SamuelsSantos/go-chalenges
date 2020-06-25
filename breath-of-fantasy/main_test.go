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

		player1 := Player{Nome: "Player 1", Energy: 20, Power: 50}
		player2 := Player{Nome: "Player 2", Energy: 20, Power: 40}
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
		player2 := Player{Nome: "Player 2", Energy: 20, Power: 40}
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

		player1 := Player{Nome: "Player 1", Energy: 20, Power: 50}
		player2 := Player{Nome: "Player 2", Energy: 20, Power: 40}
		critical := MockNumberCritical{}
		var expected = player1

		result := round(&player1, &player2, critical)
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
}
