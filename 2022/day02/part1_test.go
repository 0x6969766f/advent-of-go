package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart01(t *testing.T) {
	t.Parallel()

	rounds := createRounds(mockInput)

	assert.Len(t, rounds, 3)
	assert.Equal(t, rounds[0], Round{OpponentMove: OpponentRock, PlayerMove: PlayerPaper})
	assert.Equal(t, rounds[1], Round{OpponentMove: OpponentPaper, PlayerMove: PlayerRock})
	assert.Equal(t, rounds[2], Round{OpponentMove: OpponentScissors, PlayerMove: PlayerScissors})

	totalScore := playTournament(rounds)

	assert.Equal(t, 15, totalScore)
}
