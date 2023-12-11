package aoc

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
)

//go:embed day11.data
var day11data string

func parseGalaxies(age int64) [][2]int64 {
	galaxies := make([][2]int64, 0)

	// We keep track of the offset in the y-direction as we go along, and fix the offset
	// in the x-direction afterwards.
	yOffset := int64(0)
	forEachLine(day11data, func(y int, line string) {
		empty := true
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, [2]int64{int64(x), int64(y) + yOffset})
				empty = false
			}
		}
		if empty {
			yOffset += age
		}
	})

	// Sort by x-coordinate
	slices.SortFunc(galaxies, func(a, b [2]int64) int {
		return cmp.Compare(a[0], b[0])
	})

	for xOffset, i, col := int64(0), 0, int64(0); i < len(galaxies); {
		if col == galaxies[i][0] {
			// This is a galaxy in the current column, add current x offset and advance to
			// next galaxy.
			galaxies[i][0] += xOffset
			i++
			continue
		}
		if col+1 == galaxies[i][0] {
			// This is a galaxy in the next column, which means the next column isn't empty.
			// Advance to that column.
			col++
			continue
		}

		// In this case the next column is empty. Adjust x offset and advance to that
		// column.
		xOffset += age
		col++
	}

	return galaxies
}

func galaxyDistance(galaxies [][2]int64) int64 {
	sum := int64(0)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			dist := int64abs(galaxies[i][0]-galaxies[j][0]) +
				int64abs(galaxies[i][1]-galaxies[j][1])
			sum += dist
		}
	}
	return sum
}

func day11Part1() {
	galaxies := parseGalaxies(1)
	sum := galaxyDistance(galaxies)
	fmt.Println("result: ", sum)
}

func day11Part2() {
	galaxies := parseGalaxies(999999)
	sum := galaxyDistance(galaxies)
	fmt.Println("result: ", sum)
}
