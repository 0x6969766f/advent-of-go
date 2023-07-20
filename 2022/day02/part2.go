package day02

import (
	"fmt"
	"os"
	"strings"
)

const (
	Lose = "X"
	Draw = "Y"
	Win  = "Z"
)

type Round2 struct {
	OpponentMove string
	EndResult    string
}

func Part2() {
	f, err := os.ReadFile("./day02/input")
	if err != nil {
		panic(err)
	}
	rounds := createRounds2(string(f))
	score := playTournament2(rounds)
	fmt.Printf("- part#2: your total score is %d\n", score)
}

func createRounds2(input string) []Round2 {
	lines := strings.Split(input, "\n")
	rounds := make([]Round2, len(lines))
	for k, line := range lines {
		moves := strings.Split(line, " ")
		if len(moves) != 2 {
			panic(fmt.Sprintf("invalid input: %s", line))
		}
		rounds[k] = Round2{OpponentMove: moves[0], EndResult: moves[1]}
	}
	return rounds
}

// "Anyway, the second column says how the round needs to end: X means you need to lose,
// Y means you need to end the round in a draw, and Z means you need to win. Good luck!"
// The total score is still calculated in the same way, but now you need to figure out
// what shape to choose so the round ends as indicated.
// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
func playTournament2(input []Round2) int {
	totalScore := 0

	roundScore := 0
	for _, round := range input {
		switch round.OpponentMove {
		case OpponentRock:
			switch round.EndResult {
			case Lose:
				roundScore = ScoreScissors + ScoreLost
			case Draw:
				roundScore = ScoreRock + ScoreDraw
			case Win:
				roundScore = ScorePaper + ScoreWin
			}
		case OpponentPaper:
			switch round.EndResult {
			case Lose:
				roundScore = ScoreRock + ScoreLost
			case Draw:
				roundScore = ScorePaper + ScoreDraw
			case Win:
				roundScore = ScoreScissors + ScoreWin
			}
		case OpponentScissors:
			switch round.EndResult {
			case Lose:
				roundScore = ScorePaper + ScoreLost
			case Draw:
				roundScore = ScoreScissors + ScoreDraw
			case Win:
				roundScore = ScoreRock + ScoreWin
			}
		}

		totalScore += roundScore
		roundScore = 0
	}

	return totalScore
}
