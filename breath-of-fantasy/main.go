package main

import (
	"JV-Samuel-Santos/breath-of-fantasy/rand"
	"bufio"
	"errors"
	"fmt"
	"io"
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

func round(heroi *Player, inimigo *Player, roulette rand.Roulette) *Player {

	fmt.Printf("%v atacou %v\n", heroi.Nome, inimigo.Nome)
	luckFactor := getLockFactor(roulette.Get())
	heroi.atack(luckFactor, inimigo)

	if inimigo.Energy <= 0 {
		return heroi
	}
	return round(inimigo, heroi, roulette)
}

func isValidInput(data string) ([]string, error) {

	if isEmpty(data) {
		return nil, errors.New("Entrada inválida")
	}

	personagem := strings.Split(data, " ")

	if len(personagem) < 3 {
		return nil, errors.New("Entrada inválida")
	}

	return personagem, nil
}

func readInfo(in io.Reader, question string) string {
	scanner := bufio.NewScanner(in)
	fmt.Println(question)
	scanner.Scan()
	return scanner.Text()
}

func readPlayer(data string) (*Player, error) {

	var err error
	var personagem []string
	var energy int
	var power int

	if personagem, err = isValidInput(data); err != nil {
		return nil, err
	}

	if energy, err = strconv.Atoi(personagem[1]); err != nil {
		return nil, errors.New("Entrada inválida")
	}

	if power, err = strconv.Atoi(personagem[2]); err != nil {
		return nil, errors.New("Entrada inválida")
	}

	player := Player{}
	player.Nome = personagem[0]
	player.Power = power
	player.Energy = energy
	return &player, nil
}

func doProcess(one, two string) (*Player, error) {

	var randonNumber = rand.LuckNumber{}

	var playerOne *Player
	var playerTwo *Player
	var err error

	if playerOne, err = readPlayer(one); err != nil {
		return nil, err
	}

	if playerTwo, err = readPlayer(two); err != nil {
		return nil, err
	}

	first := rand.GetPlayerStart()
	var winner *Player
	if first == 0 {
		fmt.Printf("Batalha entre %v e %v \n", playerOne.Nome, playerTwo.Nome)
		winner = round(playerOne, playerTwo, randonNumber)
	} else {
		fmt.Printf("Batalha entre %v e %v \n", playerTwo.Nome, playerOne.Nome)
		winner = round(playerTwo, playerOne, randonNumber)
	}
	return winner, nil
}

func main() {
	one := readInfo(os.Stdin, "Entre a primeira personagem: ")
	two := readInfo(os.Stdin, "Entre a segunda personagem: ")

	if winner, err := doProcess(one, two); err != nil {
		fmt.Printf("\nFalha: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Jogo acabou, o vencedor foi %v com HP restante de %v\n", winner.Nome, winner.Energy)
	}
}
