package day01

import (
	"os"
	"sort"
)

func Part2() ([]int, int) {
	f, err := os.ReadFile("./day01/input")
	if err != nil {
		panic(err)
	}
	calorieGroups := countCalories(string(f))
	elves, calories := findElves(calorieGroups)
	return elves, calories
}

func findElves(groups map[int]int) ([]int, int) {
	keys := make([]int, 0, len(groups))
	for key := range groups {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(a, b int) bool {
		return groups[keys[a]] > groups[keys[b]]
	})

	elves := keys[:3]
	calories := 0
	for _, k := range elves {
		calories += groups[k]
	}

	return elves, calories
}
