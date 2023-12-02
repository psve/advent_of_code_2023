package main

import (
	"aoc/aoc"
	"flag"
)

func main() {
	day := flag.Uint("day", 0, "day to run")
	part := flag.Uint("part", 0, "part to run")
	flag.Parse()
	aoc.Run(*day, *part)
}
