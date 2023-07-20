package day01

import (
	"fmt"
	"math"
	"os"
)

func Part1() {
	f, err := os.ReadFile("./day01/input")
	if err != nil {
		panic(err)
	}
	calorieGroups := countCalories(string(f))
	elf, calories := findElf(calorieGroups)
	fmt.Printf("- part#1: elf #%d has the most calories: %d\n", elf, calories)
}

func findElf(groups map[int]int) (int, int) {
	elf := 0
	total := math.MinInt64
	for index, sum := range groups {
		if groups[index] > total {
			total = sum
			elf = index
		}
	}
	return elf, total
}
