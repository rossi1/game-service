package utils

import (
	"math/rand"
	"time"
)

func GenerateComputerChoice() (randNum int) {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	randNum = rand.Intn(max-min) + min
	return
}
