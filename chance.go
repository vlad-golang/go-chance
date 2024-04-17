package chance

import (
	"math/rand"
)

func FlipCoin() bool {
	return rand.Intn(2) == 0
}

func RussianRoulette() bool {
	return rand.Intn(6) == 0
}

func Percentage(percentage float64) bool {
	return rand.Float64() < percentage/100
}
