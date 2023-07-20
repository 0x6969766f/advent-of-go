package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	rounds := createRounds2(mockInput)

	assert.Len(t, rounds, 3)
	assert.Equal(t, rounds[0], Round2{OpponentMove: OpponentRock, EndResult: Draw})
	assert.Equal(t, rounds[1], Round2{OpponentMove: OpponentPaper, EndResult: Lose})
	assert.Equal(t, rounds[2], Round2{OpponentMove: OpponentScissors, EndResult: Win})

	totalScore := playTournament2(rounds)

	assert.Equal(t, 12, totalScore)
}
