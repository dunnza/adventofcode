package day04

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dunnza/adventofcode/utils"
)

func Day04() {
	fmt.Println("Day Four:")
	input := utils.ReadFile("./advent2022/day04/day04.txt")
	partOne(input)
	partTwo(input)
	fmt.Println()
}

func partOne(input string) {
	overlaps := 0
	for _, line := range strings.Split(input, "\n") {
		rx := regexp.MustCompile("\\d+")
		matches := rx.FindAllString(line, -1)
		nums := sliceMap[int](matches, toInt)
		if (nums[0] <= nums[2] && nums[1] >= nums[3]) || (nums[2] <= nums[0] && nums[3] >= nums[1]) {
			overlaps += 1
		}
	}

	fmt.Printf("Part One, consuming overlaps: %d\n", overlaps)
}

func partTwo(input string) {
	overlaps := 0
	for _, line := range strings.Split(input, "\n") {
		rx := regexp.MustCompile("\\d+")
		matches := rx.FindAllString(line, -1)
		nums := sliceMap[int](matches, toInt)
		if nums[0] <= nums[3] && nums[1] >= nums[2] {
			overlaps += 1
		}
	}

	fmt.Printf("Part Two, consuming + partial overlaps: %d\n", overlaps)
}

func sliceMap[T any](items []string, f func(string) T) []T {
	mapped := make([]T, len(items))
	for i, v := range items {
		mapped[i] = f(v)
	}

	return mapped
}

func toInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
