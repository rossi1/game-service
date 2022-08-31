package utils

import (
	"math/rand"
	"time"
)

type Choice struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Choices int64

const (
	Rock     Choices = 1
	Paper            = 2
	Scissors         = 3
	Lizard           = 4
	Spock            = 5
)

func GetChoices() []Choice {
	var choices = []Choice{
		{Id: 1, Name: "rock"},
		{Id: 2, Name: "paper"},
		{Id: 3, Name: "scissors"},
		{Id: 4, Name: "lizard"},
		{Id: 5, Name: "spock"},
	}
	return choices
}

func GetRandomChoice() Choice {
	var (
		choice  Choice
		randNum int
	)
	min := 1
	max := 6
	rand.Seed(time.Now().UnixNano())
	randNum = rand.Intn(max-min) + min

	switch randNum {
	case int(Rock):
		choice = Choice{Id: int(Rock), Name: "rock"}

	case int(Paper):
		choice = Choice{Id: int(Paper), Name: "paper"}

	case int(Scissors):
		choice = Choice{Id: int(Scissors), Name: "scissors"}

	case int(Lizard):
		choice = Choice{Id: int(Lizard), Name: "lizard"}

	case int(Spock):
		choice = Choice{Id: int(Spock), Name: "spock"}
	}
	return choice
}

func DetermineWinner(userAction, computerAction int) (result, winner string, user, computer int) {
	victories := map[int][]Choices{
		3: {Lizard, Paper},
		2: {Spock, Rock},
		1: {Lizard, Scissors},
		4: {Spock, Paper},
		5: {Scissors, Rock},
	}
	defeats := victories[userAction]
	if userAction == computerAction {
		winner = "Tie"
		result = "tie"
		user = userAction
		computer = computerAction
		return
	}

	for _, defeat := range defeats {
		if computerAction == int(defeat) {
			result = "win"
			winner = "User"
			user = userAction
			computer = computerAction
			return
		}

	}

	result = "lose"
	winner = "Computer"
	user = userAction
	computer = computerAction
	return

}
