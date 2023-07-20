package main

import (
	"0x6969766f/advent-of-go/day01"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Advent of Go (code) 2022 - DAY 1")

	elf, calories := day01.Part1()
	fmt.Printf("- part#1: elf %d has the most calories: %d\n", elf, calories)

	var e string
	elves, calories := day01.Part2()
	for i, elf := range elves {
		e += strconv.Itoa(elf)
		if i < len(elves)-1 {
			e += ", "
		}
	}
	fmt.Printf("- part#2: elves %s have the most calories: %d\n", e, calories)
}
