package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Day 1:")
	day01()
	fmt.Println()

	fmt.Println("Day 2:")
	Day02()
	fmt.Println()
}

// check in a real app you would probably not panic, but in little
// challenges like this it just makes feedback quicker
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	check(err)
	return string(data)
}
