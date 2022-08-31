package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomChoice(t *testing.T) {
	choice := GetRandomChoice()

	assert.NotEmpty(t, choice)
}

func TestDetermineWinner(t *testing.T) {
	user := 3
	computer := 2
	result, _, _, _ := DetermineWinner(user, computer)
	assert.Equal(t, result, "win")
}
