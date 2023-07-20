package day01

import (
	"strconv"
	"strings"
)

func countCalories(input string) map[int]int {
	calories := map[int]int{}

	elf := 0
	sum := 0

	lines := strings.Split(string(input), "\n")
	lastLine := lines[len(lines)-1]

	for _, line := range lines {
		if len(line) == 0 {
			calories[elf] = sum
			elf++
			sum = 0
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += num

		if line == lastLine {
			calories[elf] = sum
		}
	}

	return calories
}
