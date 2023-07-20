package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	rucksack := createRuckshack(mockInput)

	assert.Equal(t, 6, len(rucksack))

	assert.Equal(t, Rucksack{
		Left:  string("vJrwpWtwJgWr"),
		Right: string("hcsFMMfFFhFp"),
	}, rucksack[0])

	assert.Equal(t, Rucksack{
		Left:  string("jqHRNqRjqzjGDLGL"),
		Right: string("rsFMfFZSrLrFZsSL"),
	}, rucksack[1])

	assert.Equal(t, Rucksack{
		Left:  string("PmmdzqPrV"),
		Right: string("vPwwTWBwg"),
	}, rucksack[2])

	assert.Equal(t, Rucksack{
		Left:  string("wMqvLMZHhHMvwLH"),
		Right: string("jbvcjnnSBnvTQFn"),
	}, rucksack[3])

	assert.Equal(t, Rucksack{
		Left:  string("ttgJtRGJ"),
		Right: string("QctTZtZT"),
	}, rucksack[4])

	assert.Equal(t, Rucksack{
		Left:  string("CrZsJsPPZsGz"),
		Right: string("wwsLwLmpwMDw"),
	}, rucksack[5])

	items := findItems(rucksack)

	assert.Equal(t, string("p"), items[0])
	assert.Equal(t, string("L"), items[1])
	assert.Equal(t, string("P"), items[2])
	assert.Equal(t, string("v"), items[3])
	assert.Equal(t, string("t"), items[4])
	assert.Equal(t, string("s"), items[5])

	priority := calculatePriority(items)

	assert.Equal(t, 157, priority)
}
