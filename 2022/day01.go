package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day01() {
	data, err := os.ReadFile("./input/day01.txt")
	check(err)
	input := string(data)

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
		check(err)
		current += calories
	}

	// print results
	fmt.Printf("Part One, most calories: %d\n", first)
	fmt.Printf("Part Two, top three calories: %d + %d + %d = %d\n", first, second, third, first+second+third)
}
