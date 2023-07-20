package day01

import (
	"math"
	"os"
)

func Part1() (int, int) {
	f, err := os.ReadFile("./day01/input")
	if err != nil {
		panic(err)
	}
	calorieGroups := countCalories(string(f))
	elf, calories := findElf(calorieGroups)
	return elf, calories
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
