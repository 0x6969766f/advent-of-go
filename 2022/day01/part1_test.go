package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	elfs := countCalories(mockInput)

	assert.Len(t, elfs, 5)

	assert.Equal(t, 6000, elfs[0])
	assert.Equal(t, 4000, elfs[1])
	assert.Equal(t, 11000, elfs[2])
	assert.Equal(t, 24000, elfs[3])
	assert.Equal(t, 10000, elfs[4])

	elf, calories := findElf(elfs)

	assert.Equal(t, 3, elf)
	assert.Equal(t, 24000, calories)
}
