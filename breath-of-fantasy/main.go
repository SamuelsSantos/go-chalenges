package main

import (
	"JV-Samuel-Santos/troco/rand"
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Player type
type Player struct {
	Nome   string // Player Name
	Power  int    // Player Power
	Energy int    // Player Energy
}

// LuckFactor type
type LuckFactor struct {
	Min     int     // Min Factor
	Max     int     // Max Factor
	Damage  float64 // Damage Factor
	Message string  // Message Round
}

var wrong LuckFactor = LuckFactor{Min: 0, Max: 15, Damage: 0, Message: "Errou - %d HP \n"}
var normal LuckFactor = LuckFactor{Min: 16, Max: 70, Damage: 0.33, Message: "Normal - %d HP \n"}
var good LuckFactor = LuckFactor{Min: 71, Max: 96, Damage: 0.396, Message: "Sorte!!! - %d HP \n"}
var critical LuckFactor = LuckFactor{Min: 97, Max: 100, Damage: 0.495, Message: "Crítico! - %d HP \n"}

func isEmpty(text string) bool {
	return strings.TrimSpace(text) == ""
}

func getLockFactor(luckNumber int) LuckFactor {
	switch {
	case luckNumber <= 15:
		return wrong
	case luckNumber >= 16 && luckNumber <= 70:
		return normal
	case luckNumber >= 71 && luckNumber <= 96:
		return good
	default:
		return critical
	}
}

func (heroi *Player) atack(factor LuckFactor, inimigo *Player) *Player {
	damage := int(math.Floor(float64((*heroi).Power) * factor.Damage))
	inimigo.Energy -= damage
	fmt.Printf(factor.Message, damage)
	return inimigo
}

func round(playerOne *Player, playerTwo *Player, roulette rand.Roulette) *Player {
	if playerOne.Energy <= 0 {
		return playerTwo
	}

	if playerTwo.Energy <= 0 {
		return playerOne
	}

	fmt.Printf("%v atacou %v\n", playerOne.Nome, playerTwo.Nome)
	luckFactor := getLockFactor(roulette.Get())
	playerOne.atack(luckFactor, playerTwo)
	return round(playerTwo, playerOne, roulette)
}

func fight(playerOne *Player, playerTwo *Player) {

	var randonNumber = rand.LuckNumber{}

	fmt.Printf("Batalha entre %v e %v \n", playerOne.Nome, playerTwo.Nome)
	winner := round(playerOne, playerTwo, randonNumber)
	fmt.Printf("Jogo acabou, o vencedor foi %v com HP restante de %v\n", winner.Nome, winner.Energy)
}

func isValid(data string) ([]string, error) {

	if isEmpty(data) {
		return nil, errors.New("Entrada inválida")
	}

	personagem := strings.Split(data, " ")

	if len(personagem) < 3 {
		return nil, errors.New("Entrada inválida")
	}

	for _, element := range personagem {
		if isEmpty(element) {
			return nil, errors.New("Entrada inválida")
		}
	}

	return personagem, nil
}

func readPlayer(data string) (*Player, error) {

	var err error
	var personagem []string
	if personagem, err = isValid(data); err != nil {
		return nil, err
	}

	player := Player{}
	player.Nome = personagem[0]
	if player.Energy, err = strconv.Atoi(personagem[1]); err != nil {
		return nil, errors.New("Entrada inválida")
	}

	if player.Power, err = strconv.Atoi(personagem[2]); err != nil {
		return nil, errors.New("Entrada inválida")
	}

	return &player, nil

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Entre a primeira personagem: ")
	scanner.Scan()
	var playerOne *Player
	var playerTwo *Player
	var err error
	if playerOne, err = readPlayer(scanner.Text()); err != nil {
		fmt.Printf("\nFalha: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Entre a segunda personagem: ")
	scanner.Scan()
	if playerTwo, err = readPlayer(scanner.Text()); err != nil {
		fmt.Printf("\nFalha: %v\n", err)
		os.Exit(1)
	}

	first := rand.GetPlayerStart()
	if first == 0 {
		fight(playerOne, playerTwo)
	} else {
		fight(playerTwo, playerOne)
	}

}
