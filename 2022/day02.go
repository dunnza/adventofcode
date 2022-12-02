package main

import (
	"fmt"
	"strings"
)

const (
	opponentRock     = "A"
	opponentPaper    = "B"
	opponentScissors = "C"
	myRock           = "X"
	myPaper          = "Y"
	myScissors       = "Z"
	lose             = "X"
	draw             = "Y"
	win              = "Z"
)

func day02_a(input string) {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opp, me := getParts(line)
		score += getPointsForChoice(me)
		score += getPointsForOutcome(opp, me)
	}

	fmt.Printf("Part One, my score: %d\n", score)
}

func day02_b(input string) {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		opp, outcome := getParts(line)
		me := getChoiceFromOutcome(opp, outcome)
		score += getPointsForChoice(me)
		score += getPointsForOutcome(opp, me)
	}

	fmt.Printf("Part Two, my score based on outcome: %d\n", score)
}

func Day02() {
	input := readFile("./input/day02.txt")
	day02_a(input)
	day02_b(input)
}

func getParts(line string) (string, string) {
	temp := strings.Split(line, " ")
	return temp[0], temp[1]
}

func getPointsForChoice(choice string) int {
	switch choice {
	case myRock:
		return 1
	case myPaper:
		return 2
	case myScissors:
		return 3
	default:
		panic("got a choice we didn't expect: " + choice)
	}
}

func getPointsForOutcome(opp, me string) int {
	if (opp == opponentRock && me == myPaper) || (opp == opponentPaper && me == myScissors) || (opp == opponentScissors && me == myRock) {
		return 6
	} else if (opp == opponentRock && me == myRock) || (opp == opponentPaper && me == myPaper) || (opp == opponentScissors && me == myScissors) {
		return 3
	} else if (opp == opponentRock && me == myScissors) || (opp == opponentPaper && me == myRock) || (opp == opponentScissors && me == myPaper) {
		return 0
	} else {
		panic(fmt.Sprintf("got a combination we didn't expect: %s (opp) vs. %s (me)", opp, me))
	}
}

func getChoiceFromOutcome(opp, outcome string) string {
	if (opp == opponentRock && outcome == lose) || (opp == opponentPaper && outcome == win) || (opp == opponentScissors && outcome == draw) {
		return myScissors
	} else if (opp == opponentRock && outcome == draw) || (opp == opponentPaper && outcome == lose) || (opp == opponentScissors && outcome == win) {
		return myRock
	} else if (opp == opponentRock && outcome == win) || (opp == opponentPaper && outcome == draw) || (opp == opponentScissors && outcome == lose) {
		return myPaper
	} else {
		panic(fmt.Sprintf("got a combination we didn't expect: %s (opp) with outcome %s", opp, outcome))
	}
}
