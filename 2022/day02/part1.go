package day02

import (
	"fmt"
	"os"
	"strings"
)

const (
	OpponentRock     = "A"
	OpponentPaper    = "B"
	OpponentScissors = "C"

	PlayerRock     = "X"
	PlayerPaper    = "Y"
	PlayerScissors = "Z"

	ScoreRock     = 1
	ScorePaper    = 2
	ScoreScissors = 3

	ScoreLost = 0
	ScoreDraw = 3
	ScoreWin  = 6
)

type Round struct {
	PlayerMove   string
	OpponentMove string
}

func Part1() {
	f, err := os.ReadFile("./day02/input")
	if err != nil {
		panic(err)
	}
	rounds := createRounds(string(f))
	score := playTournament(rounds)

	fmt.Printf("- part#1: your total score is %d\n", score)
}

func createRounds(input string) []Round {
	lines := strings.Split(input, "\n")
	rounds := make([]Round, len(lines))
	for k, line := range lines {
		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			panic(fmt.Sprintf("invalid input: %s", line))
		}
		rounds[k] = Round{OpponentMove: moves[0], PlayerMove: moves[1]}
	}
	return rounds
}

// winner of the tournament is the player with highest score
// your total score is sum of your scores in each round
// single round score is based on the shape you select:
// - rock: 1 point
// - paper: 2 points
// - scissors: 3 points
// plus the score for the outcome of the round:
// - lost: 0 points
// - draw: 3 points
// - won: 6 points
// rock defeats scissors, scissors defeats paper, paper defeats rock
// if both players select the same shape, the round is a draw

func playTournament(rounds []Round) int {
	totalScore := 0
	roundScore := 0
	for _, round := range rounds {
		switch round.OpponentMove {
		case OpponentRock:
			switch round.PlayerMove {
			case PlayerRock:
				roundScore = ScoreRock + ScoreDraw
			case PlayerPaper:
				roundScore = ScorePaper + ScoreWin
			case PlayerScissors:
				roundScore = ScoreScissors + ScoreLost
			}
		case OpponentPaper:
			switch round.PlayerMove {
			case PlayerRock:
				roundScore = ScoreRock + ScoreLost
			case PlayerPaper:
				roundScore = ScorePaper + ScoreDraw
			case PlayerScissors:
				roundScore = ScoreScissors + ScoreWin
			}
		case OpponentScissors:
			switch round.PlayerMove {
			case PlayerRock:
				roundScore = ScoreRock + ScoreWin
			case PlayerPaper:
				roundScore = ScorePaper + ScoreLost
			case PlayerScissors:
				roundScore = ScoreScissors + ScoreDraw
			}
		}
		totalScore += roundScore
		roundScore = 0
	}
	return totalScore
}
