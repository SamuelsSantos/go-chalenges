package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Coin float64 value
type Coin struct {
	value float64
}

// ChangesCoins -> [100]0.0 value coin and amount, value of 100, 50, 20, 10, 5, 2, 1.0, 0.5, 0.25, 0.10, 0.05
type ChangesCoins map[Coin]float64

func getCoins() []Coin {
	return []Coin{
		Coin{value: 100.00},
		Coin{value: 50.00},
		Coin{value: 20.00},
		Coin{value: 10.00},
		Coin{value: 5.00},
		Coin{value: 2.00},
		Coin{value: 1.0},
		Coin{value: 0.5},
		Coin{value: 0.25},
		Coin{value: 0.10},
		Coin{value: 0.05},
	}
}

const min = 0.05

func printValue(amount float64, value float64) {
	svalue := strconv.FormatFloat(value, 'f', 2, 64)
	fmt.Printf(" - %+v x R$ %+v\n", amount, strings.Replace(svalue, ".", ",", 1))
}

func formatFloat(value float64, precision int) (float64, error) {
	svalue := strconv.FormatFloat(value, 'f', precision, 64)
	return strconv.ParseFloat(svalue, 64)
}

func (b *ChangesCoins) receipt(total float64) {

	list := (*b)
	stotal := strconv.FormatFloat(total, 'f', 2, 64)
	fmt.Printf("Troco: R$ %+v\n", strings.Replace(stotal, ".", ",", 1))
	for coin, amount := range list {
		printValue(amount, coin.value)
	}
}

func (b *ChangesCoins) getChange(value float64, coins []Coin) *ChangesCoins {

	if value <= 0 {
		return b
	}

	index := 0
	if value < coins[index].value {
		index++
		return (*b).getChange(value, coins[index:])
	}

	var mod float64
	var amount float64
	mod, _ = formatFloat(math.Mod(value, coins[index].value), 2)
	amount = math.Floor(value / coins[index].value)

	if mod == 0 || mod < min {
		(*b)[coins[index]] = amount
		return b
	}

	(*b)[coins[index]] = amount
	index++
	return (*b).getChange(mod, coins[index:])
}

func readFloat(stdIn io.Reader) (float64, error) {
	scanner := bufio.NewScanner(stdIn)
	scanner.Scan()
	in := strings.Replace(scanner.Text(), ",", ".", 1)
	return strconv.ParseFloat(in, 64)
}

func doProcess(stdIn io.Reader) (float64, *ChangesCoins, error) {

	var amountPaid float64
	var amountPayable float64
	var err error

	fmt.Print("Valor a ser pago: R$ ")
	if amountPayable, err = readFloat(stdIn); err != nil {
		fmt.Printf("Falha: %+v\n", err)
		return 0, nil, err
	}

	fmt.Print("Valor efetivamente pago: R$  ")
	if amountPaid, err = readFloat(stdIn); err != nil {
		fmt.Printf("Falha: %+v\n", err)
		return 0, nil, err
	}

	if amountPaid < amountPayable {
		fmt.Printf("Falha: %+v\n", "O valor efetivamente pago é menor do que o valor a ser pago.")
		return 0, nil, err
	}

	change := math.Abs(amountPayable - amountPaid)
	changeCoins := &ChangesCoins{}
	changeCoins.getChange(change, getCoins())
	return change, changeCoins, nil
}

func main() {
	if change, changeCoins, err := doProcess(os.Stdin); err != nil {
		os.Exit(1)
	} else {
		changeCoins.receipt(change)
	}
}
