package aoc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Run(day, part uint) {
	fmt.Printf("running day %d, part %d\n", day, part)
	start := time.Now()
	defer func() { fmt.Printf("done in %v\n", time.Since(start)) }()

	switch day {
	case 1:
		runParts(day01Part1, day01Part2)(part)
	case 2:
		runParts(day02Part1, day02Part2)(part)
	case 3:
		runParts(day03Part1, day03Part2)(part)
	case 4:
		runParts(day04Part1, day04Part2)(part)
	case 5:
		runParts(day05Part1, day05Part2)(part)
	case 6:
		runParts(day06Part1, day06Part2)(part)
	case 7:
		runParts(day07Part1, day07Part2)(part)
	default:
		panic("unknown day")
	}
}

func runParts(part1, part2 func()) func(uint) {
	return func(part uint) {
		switch part {
		case 1:
			part1()
		case 2:
			part2()
		default:
			panic("unknown part")
		}
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func forEachLine(s string, f func(i int, line string)) {
	out := strings.Split(s, "\n")
	for i, line := range out[:len(out)-1] { // Strip newline at EOF
		f(i, line)
	}
}

func charToNumber(c rune) (int64, error) {
	return strconv.ParseInt(string(c), 10, 64)
}

func stringToNumber(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

var numbers = map[string]int64{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func wordToNumber(s string) (int64, error) {
	for i := 3; i <= min(5, len(s)); i++ {
		if v, ok := numbers[s[:i]]; ok {
			return v, nil
		}
	}
	return 0, errors.New("not a number")
}

var digits = map[rune]any{
	'0': nil,
	'1': nil,
	'2': nil,
	'3': nil,
	'4': nil,
	'5': nil,
	'6': nil,
	'7': nil,
	'8': nil,
	'9': nil,
}

func isDigit(c rune) bool {
	_, yes := digits[c]
	return yes
}
