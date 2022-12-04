package day03

import (
	"fmt"
	"strings"

	"github.com/dunnza/adventofcode/utils"
)

func Day03() {
	fmt.Println("Day Three:")
	input := utils.ReadFile("./advent2022/day03/day03.txt")
	partOne(input)
	partTwo(input)
	fmt.Println()
}

func partOne(input string) {
	total := 0
	for _, line := range strings.Split(input, "\n") {
		length := len(line)
		cOne := line[:length/2]
		cTwo := line[length/2:]
		for _, item := range cOne {
			if strings.ContainsRune(cTwo, item) {
				total += strings.IndexRune(priority, item) + 1
				break
			}
		}
	}

	fmt.Printf("Part One, total priority: %d\n", total)
}

func partTwo(input string) {
	total := 0
	lines := strings.Split(input, "\n")

	// we do len(lines) + 1 so the modulo arithmetic works, we never try to index the slice past its length
	for i := 0; i < len(lines)+1; i++ {
		if i%3 == 0 && i != 0 {
			group := lines[i-3 : i]
			for _, item := range group[0] {
				if strings.ContainsRune(group[1], item) && strings.ContainsRune(group[2], item) {
					total += strings.IndexRune(priority, item) + 1
					break
				}
			}
		}
	}

	fmt.Printf("Part Two, priority for group badges: %d\n", total)
}

const priority = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
