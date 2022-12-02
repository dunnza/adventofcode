package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dunnza/adventofcode/utils"
)

func Day01() {
	fmt.Println("Day One:")
	input := utils.ReadFile("./advent2022/day01/day01.txt")
	partOne(input)
	partTwo(input)
	fmt.Println()
}

func partOne(input string) {
	// chunk data into lines, summing up the number until we get to an empty line
	current, most := 0, 0
	for _, line := range strings.Split(input, "\n") {
		// if we hit an empty line, that means we've counted all the calories for
		// the current elf, so we can now do our comparisons
		if line == "" {
			if current > most {
				most = current
			}

			current = 0
			continue
		}

		// we assume there's no error for simplicity
		calories, err := strconv.Atoi(line)
		utils.Check(err)
		current += calories
	}

	// print results
	fmt.Printf("Part One, most calories: %d\n", most)
}

func partTwo(input string) {
	// chunk data into lines, summing up the number until we get to an empty line
	current := 0
	first, second, third := 0, 0, 0
	for _, line := range strings.Split(input, "\n") {
		// if we hit an empty line, that means we've counted all the calories for
		// the current elf, so we can now do our comparisons
		if line == "" {
			temp := current
			if temp > first {
				first, temp = temp, first
				if temp > second {
					second, temp = temp, second
					if temp > third {
						third = temp
					}
				} else if temp > third {
					third = temp
				}
			} else if temp > second {
				second, temp = temp, second
				if temp > third {
					third = temp
				}
			} else if temp > third {
				third = temp
			}

			current = 0
			continue
		}

		// we assume there's no error for simplicity
		calories, err := strconv.Atoi(line)
		utils.Check(err)
		current += calories
	}

	// print results
	fmt.Printf("Part Two, top three calories: %d + %d + %d = %d\n", first, second, third, first+second+third)
}
