package main

import (
	"bufio"
	"errors"
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

func isValidValue(amount string) (float64, error) {
	return strconv.ParseFloat(strings.Replace(amount, ",", ".", 1), 64)
}

func readInfo(stdIn io.Reader, question string) string {
	fmt.Print(question)
	scanner := bufio.NewScanner(stdIn)
	scanner.Scan()
	return scanner.Text()
}

func doProcess(amountPayable, amountPaid string) (float64, *ChangesCoins, error) {

	var err error
	var payable float64
	var paid float64

	if payable, err = isValidValue(amountPayable); err != nil {
		fmt.Printf("Falha: %+v\n", err)
		return 0, nil, err
	}

	if paid, err = isValidValue(amountPaid); err != nil {
		fmt.Printf("Falha: %+v\n", err)
		return 0, nil, err
	}

	if paid < payable {
		return 0, nil, errors.New("O valor efetivamente pago é menor do que o valor a ser pago")
	}

	change := math.Abs(payable - paid)
	changeCoins := &ChangesCoins{}
	changeCoins.getChange(change, getCoins())
	return change, changeCoins, nil
}

func main() {

	amountPayable := readInfo(os.Stdin, "Valor a ser pago: R$ ")
	amountPaid := readInfo(os.Stdin, "Valor efetivamente pago: R$ ")

	if change, changeCoins, err := doProcess(amountPayable, amountPaid); err != nil {
		fmt.Printf("Falha: %+v.\n", err)
		os.Exit(1)
	} else {
		changeCoins.receipt(change)
	}
}
