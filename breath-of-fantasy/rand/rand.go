package rand

import (
	"math/rand"
	"time"
)

// Roulette interface
type Roulette interface {
	Get() int
}

// LuckNumber struct
type LuckNumber struct{}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Get return a number between 0 e 100
func (LuckNumber) Get() int {
	return seededRand.Intn(101)
}

// GetPlayerStart return a number between 0 e 1
func GetPlayerStart() int {
	return seededRand.Intn(2)
}
