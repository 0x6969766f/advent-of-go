package day03

import (
	"fmt"
	"os"
	"strings"
)

type Rucksack struct {
	Left  string
	Right string
}

type Sacks map[int]Rucksack

type Items map[int]string

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Part1() {
	f, err := os.ReadFile("./day03/input")
	if err != nil {
		panic(err)
	}
	rucksack := createRuckshack(string(f))
	items := findItems(rucksack)
	priority := calculatePriority(items)
	fmt.Printf("- part#3: sum of priorities is %d\n", priority)
}

func calculatePriority(items Items) int {
	sum := 0
	for _, item := range items {
		sum += strings.Index(characters, item) + 1
	}
	return sum
}

func createRuckshack(input string) Sacks {
	lines := strings.Split(input, "\n")
	sacks := make(Sacks, len(lines))
	for i, line := range lines {
		pocket := strings.Split(line, " ")
		if len(pocket) != 1 {
			panic(fmt.Sprintf("invalid input: %s", line))
		}
		size := len([]rune(pocket[0])) / 2
		sacks[i] = Rucksack{Left: pocket[0][:size], Right: pocket[0][size:]}
	}
	return sacks
}

func findItems(shacks Sacks) Items {
	items := make(Items)
	for i, shack := range shacks {
		left := strings.Split(shack.Left, ",")
		if len(left) != 1 {
			panic(fmt.Sprintf("invalid input: %s", shack.Left))
		}
	inner:
		for _, char := range left[0] {
			c := string(char)
			if strings.Contains(shack.Right, c) {
				items[i] = c
				break inner
			}
		}
	}
	return items
}
